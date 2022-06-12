package user

import "github.com/google/wire"

func initApp() app {
	wire.Build(NewApp, newDBConfig, newDB)
	return app{}
}
