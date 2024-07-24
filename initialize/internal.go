package initialize

import (
	"github.com/gin-gonic/gin"
	"gowebtest/api/v1/admin"
	card2 "gowebtest/api/v1/card"
	user2 "gowebtest/api/v1/user"
	"gowebtest/data"
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
				{http.MethodGet, "login", user2.UserLoginInit},
				{http.MethodPost, "login", user2.UserLogin},
				{http.MethodPost, "register", user2.UserRegister},
				{http.MethodGet, "home/:name", user2.UserHome},
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
				{http.MethodGet, "cards", card2.Cards},
				{http.MethodGet, "gacha", card2.GachaCard},
				{http.MethodPost, "cardsadd", card2.CardAdd},
			},
		},
	}
	router.Use(middleware.Cors())
	route.InitRoutes(router, routes)
	router.Run(":8080")
}
func DbInit() {
	//测试数据库连接
	db := data.Database{}
	db.DatabaseConnect()

}
