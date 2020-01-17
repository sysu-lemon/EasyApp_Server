package main

import "com/EasyApp_Server/routers"

func main() {
	r := routers.Init()
	r.Run(":8000")
}

