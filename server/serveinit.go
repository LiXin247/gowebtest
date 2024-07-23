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
				{http.MethodGet, "/", index.Index},
			},
		},
		{
			GroupName: "User",
			Routes: []route.Route{
				{http.MethodGet, "/user/login", user.UserLoginInit},
				{http.MethodPost, "/user/login", user.UserLogin},
				{http.MethodPost, "/user/register", user.UserRegister},
			},
		},
	}
	route.InitRoutes(router, routes)
	router.Run()
}
