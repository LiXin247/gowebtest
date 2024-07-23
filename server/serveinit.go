package main

import (
	"github.com/gin-gonic/gin"
	"gowebtest/data"
	"gowebtest/index"
	"gowebtest/route"
	"gowebtest/user"
	"net/http"
)

func ServeInit() {
	//测试数据库连接
	db := data.Database{}
	db.DatabaseConnect()
	//注册路由
	router := gin.Default()
	routes := []route.Routegroup{
		{
			GroupName: "Index",
			Routes: []route.Route{
				{Method: http.MethodGet, Path: "/", HandlerFunc: index.Index},
			},
		},
		{
			GroupName: "User",
			Routes: []route.Route{
				{Method: http.MethodGet, Path: "/user/login", HandlerFunc: user.UserLogin},
				{Method: http.MethodGet, Path: "/user/register", HandlerFunc: user.UserRegister},
			},
		},
	}
	route.InitRoutes(router, routes)
}
