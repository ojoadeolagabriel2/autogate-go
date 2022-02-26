package main

import (
	"autogate/api"
	"autogate/app"
)

const (
	EnvPgUsernameKey = "ENV_PG_USERNAME"
	EnvPgPwdKey      = "ENV_PG_USERNAME"
)

func main() {
	application := app.App{}

	application.InitializeRouter()
	//application.InitializeDb(
	//	util.GetConfig(EnvPgUsernameKey, "postgres"),
	//	util.GetConfig(EnvPgPwdKey, "postgres"),
	//)

	api.RegisterRoutes(application.Router)
}
