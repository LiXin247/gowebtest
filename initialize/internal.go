package initialize

import (
	"github.com/gin-gonic/gin"
	"gowebtest/api/v1/admin"
	"gowebtest/api/v1/card"
	"gowebtest/api/v1/user"
	"gowebtest/middleware"
	"gowebtest/route"
	"net/http"
)

func RouteInit() {
	//注册路由
	router := gin.Default()
	routes := []route.Routegroup{

		{
			GroupName: "user",
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
	router.Use(middleware.Cors())
	route.InitRoutes(router, routes)
	router.Run(":8080")
}
