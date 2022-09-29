package main

import (
	"mygame/webmain"
)

func main() {
	webmain.Init()
	webmain.GIN.Run(":8088")
}
