package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	v1 "github.com/mytlogos/netbox_application_controller/api/v1"
	"github.com/mytlogos/netbox_application_controller/services"
	"github.com/mytlogos/netbox_application_controller/services/backend"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/versioning"

	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	Server struct {
		Port int
	}
	Backend struct {
		Server string
		Apikey string
	}
	VersionSearcher *struct {
		Server string
	}
}

func runVersionSearcher(config services.VersionSearcherConfig) {
	versionSearch := services.NewSearcher(config)

	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		err := versionSearch.Search()

		if err != nil {
			slog.Error("version searcher had an error", "error", err)
		}
	}
}

func main() {
	config.WithOptions(config.ParseEnv)
	config.WithOptions(config.Readonly)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yml")

	if err != nil {
		panic(err)
	}

	var myConfig Config
	config.Decode(&myConfig)

	bc := backend.BackendConfig{
		Server: myConfig.Backend.Server,
		Apikey: myConfig.Backend.Apikey,
	}

	be := backend.NewBackend(bc)
	err = be.InitBackend()

	if err != nil {
		log.Fatalln(err)
	}

	storage := services.NewStorage(be)
	storage.Init()

	if myConfig.VersionSearcher != nil {
		if myConfig.VersionSearcher.Server == "" {
			log.Fatalln("config has VersionSearcher specified but not its server")
		}
		slog.Info("version searcher configured - starting")
		go runVersionSearcher(services.VersionSearcherConfig{
			Backend: be,
			Argus: &services.ArgusConfig{
				Server: myConfig.VersionSearcher.Server,
			},
		})
	}

	run(myConfig, storage)
}

func run(config Config, storage *services.Storage) {
	app := iris.New()
	// Serve our front-end and its assets.
	app.HandleDir("/", iris.Dir("./app/dist"))

	// Note, it's buffered, so make sure it's closed so it can flush any buffered contents.
	ac := accesslog.File("./access.log")
	defer ac.Close()

	app.Logger().SetLevel("debug")
	app.UseRouter(ac.Handler)
	app.UseRouter(recover.New())
	// Group routes and mvc apps based on /api path prefix.
	api := app.Party("/api")

	api.UseRouter(versioning.Aliases(versioning.AliasMap{
		versioning.Empty: "1.0.0",
	}))

	v1Api := versioning.NewGroup(api, ">=1.0.0 <2.0.0")
	v1.AddApi(v1Api.API, storage)

	app.Listen(fmt.Sprintf(":%d", config.Server.Port))
}
