package main

import (
	"github.com/gin-gonic/gin"
	"gowebtest/data"
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
			GroupName: "customer",
			Routes: []route.Route{
				{Method: http.MethodGet, Path: "/get", HandlerFunc: user.UserLogin},
			},
		},
	}
	route.InitRoutes(router, routes)
}
