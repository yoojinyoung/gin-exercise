package main

import "yoojinyoung/gin/app"

func main() {
	appModule := app.NewAppModule()
	appModule.Init()
}
