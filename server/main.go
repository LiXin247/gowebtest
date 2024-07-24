package main

import (
	"gowebtest/dbconnect"
	"gowebtest/initialize"
)

func main() {
	dbconnect.DbInit()     //数据库初始化
	initialize.RouteInit() //路由初始化

}
