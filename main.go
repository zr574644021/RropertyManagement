package main

import (
	_ "PowerManage/routers"
	_ "PowerManage/initData"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

