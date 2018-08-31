package main

import (
	_ "CloudManagerSystem/routers"
	_ "CloudManagerSystem/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
