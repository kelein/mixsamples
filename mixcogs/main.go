package main

import (
	"mixcogs/commands"
	_ "mixcogs/config/viper"
	_ "mixcogs/config/dotenv"
	_ "mixcogs/di"
	"github.com/mix-go/xcli"
	"github.com/mix-go/xutil/xenv"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(xenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
