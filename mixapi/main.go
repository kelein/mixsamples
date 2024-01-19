package main

import (
	"mixapi/commands"
	_ "mixapi/config/viper"
	_ "mixapi/config/dotenv"
	_ "mixapi/di"
	"github.com/mix-go/xcli"
	"github.com/mix-go/xutil/xenv"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(xenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
