package main

import (
	"fmt"
	"golang-CICD/config"
	"golang-CICD/lib"
	"golang-CICD/router"
	"os"
	"runtime"
)

const (
	ConfigPath = "./config/config.yml"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: run with env [pro|dev]")
		return
	}
	env := os.Args[1]
	if env == "pro" {
		_ = os.Setenv("PROGRAM_ENV", "pro")
	}

	config.ReadEnv(ConfigPath)
	lib.InitLogger()
	lib.InitRedis()
	lib.InitMysql()
	runtime.GOMAXPROCS(runtime.NumCPU())
	engine := router.SetupServer()
	serverUrl := fmt.Sprintf("%s:%s", config.Conf.Server.Host, config.Conf.Server.Port)
	fmt.Println(fmt.Sprintf("server Listen: http://%s", serverUrl))
	_ = engine.Run(serverUrl)
}
