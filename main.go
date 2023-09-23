package main

import (
	"fmt"
	"log"

	v1 "github.com/mytlogos/netbox_application_controller/api/v1"
	"github.com/mytlogos/netbox_application_controller/controllers"
	"github.com/mytlogos/netbox_application_controller/services"
	"github.com/mytlogos/netbox_application_controller/services/backend"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
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
	{
		// Group based on /api/counter path prefix.
		counterAPI := api.Party("/counter")
		// Optionally, a <trick> to keep the `m` local variable
		// unaccessible outside of this block's scope. That
		// way you can register many mvc apps for different Parties
		// with a "m" variable.
		// Alternatively you can use the mvc.Configure function :)

		// Register a new MVC Application to the counterAPI Party.
		m := mvc.New(counterAPI)
		m.Register(
			// Register a static dependency (static because it doesn't accept an iris.Context,
			// only one instance of that it's used). Helps us to keep a global counter across
			// clients requests.
			services.NewGlobalCounter(),
			// Register a dynamic dependency (GetFields accepts an iris.Context,
			// it binds a new instance on every request). Helps us to
			// set custom fields based on the request handler.
			accesslog.GetFields,
		)
		// Register our controller.
		m.Handle(new(controllers.Counter))
	}

	api.UseRouter(versioning.Aliases(versioning.AliasMap{
		versioning.Empty: "1.0.0",
	}))

	v1Api := versioning.NewGroup(api, ">=1.0.0 <2.0.0")
	v1.AddApi(v1Api.API, storage)

	app.Listen(fmt.Sprintf(":%d", config.Server.Port))
}
