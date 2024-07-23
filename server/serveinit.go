package main

import (
	"github.com/gin-gonic/gin"
	"gowebtest/dbconnect"
	"gowebtest/route"
	"gowebtest/server/index"
	user2 "gowebtest/server/user"
	"net/http"
)

func ServeInit() {
	//测试数据库连接
	db := dbconnect.Database{}
	db.DatabaseConnect()
	//注册路由
	router := gin.Default()
	routes := []route.Routegroup{
		{
			GroupName: "Index",
			Routes: []route.Route{
				{http.MethodGet, "/", index.Index},
			},
		},
		{
			GroupName: "User",
			Routes: []route.Route{
				{http.MethodGet, "/user/login", user2.UserLoginInit},
				{http.MethodPost, "/user/login", user2.UserLogin},
				{http.MethodPost, "/user/register", user2.UserRegister},
			},
		},
	}
	route.InitRoutes(router, routes)
	router.Run()
}
