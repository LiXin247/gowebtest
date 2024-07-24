package main

import "gowebtest/initialize"

func main() {
	initialize.DbInit()    //数据库初始化
	initialize.RouteInit() //路由初始化

}
