package main

import "github.com/sysu-lemon/EasyApp_Server/routers"

func main() {
	r := routers.Init()
	r.Run()
}
