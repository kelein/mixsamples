package main

import (
	"mixweb/commands"
	_ "mixweb/config/viper"
	_ "mixweb/config/dotenv"
	_ "mixweb/di"
	"github.com/mix-go/xcli"
	"github.com/mix-go/xutil/xenv"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(xenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
