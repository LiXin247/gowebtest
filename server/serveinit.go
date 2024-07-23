package main

import (
	"github.com/gin-gonic/gin"
	"gowebtest/dbconnect"
	"gowebtest/route"
	"gowebtest/server/admin"
	"gowebtest/server/card"
	"gowebtest/server/index"
	"gowebtest/server/user"
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
				{http.MethodGet, "login", user.UserLoginInit},
				{http.MethodPost, "login", user.UserLogin},
				{http.MethodPost, "register", user.UserRegister},
				{http.MethodGet, "home/:name", user.UserHome},
			},
		},
		{
			GroupName: "admin",
			Routes: []route.Route{
				{http.MethodGet, "home/:name", admin.AdminHome},
			},
		},
		{
			GroupName: "card",
			Routes: []route.Route{
				{http.MethodGet, "cards", card.Cards},
				{http.MethodGet, "gacha", card.GachaCard},
				{http.MethodPost, "cardsadd", card.CardAdd},
			},
		},
	}
	route.InitRoutes(router, routes)
	router.Run()
}
