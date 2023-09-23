package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mytlogos/netbox_application_controller/controllers"
	"github.com/mytlogos/netbox_application_controller/services"
)

func agentApi(agentApi router.Party, storage *services.Storage) {
	{
		cpuAPI := agentApi.Party("/cpu")
		m := mvc.New(cpuAPI)
		m.Register(storage)
		m.Handle(new(controllers.CPU))
	}
	{
		devicePI := agentApi.Party("/device")
		m := mvc.New(devicePI)
		m.Register(storage)
		m.Handle(new(controllers.Device))
	}
	{
		diskApi := agentApi.Party("/disk")
		m := mvc.New(diskApi)
		m.Register(storage)
		m.Handle(new(controllers.Disk))
	}
	{
		networkApi := agentApi.Party("/network")
		m := mvc.New(networkApi)
		m.Register(storage)
		m.Handle(new(controllers.Network))
	}
	{
		ramAPI := agentApi.Party("/ram")
		m := mvc.New(ramAPI)
		m.Register(storage)
		m.Handle(new(controllers.Ram))
	}
	{
		appAPI := agentApi.Party("/application")
		m := mvc.New(appAPI)
		m.Register(storage)
		m.Handle(new(controllers.App))
	}
}

func addUserApi(api router.Party, storage *services.Storage) {
	opts := basicauth.Options{
		// TODO: replace this hardcoded with config
		Allow:        basicauth.AllowUsersFile("users.yml"),
		Realm:        basicauth.DefaultRealm,
		ErrorHandler: basicauth.DefaultErrorHandler,
		// [...more options]
	}
	auth := basicauth.New(opts)
	api.Use(auth)

	{
		userAPI := api.Party("/")
		m := mvc.New(userAPI)
		m.Register(storage)
		m.Handle(new(controllers.User))
	}
}

func AddApi(api router.Party, storage *services.Storage) {
	agentApi(api.Party("/agent"), storage)
	addUserApi(api.Party("/user"), storage)
}
