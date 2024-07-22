package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Route 定义每个路由的结构体
type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

// Routegroup 定义一组路由的结构体
type Routegroup struct {
	GroupName string
	Routes    []Route
}

func InitRoutes(router *gin.Engine, routeGroups []Routegroup) {
	for _, group := range routeGroups {
		rg := router.Group(group.GroupName)
		for _, route := range group.Routes {
			switch route.Method {
			case http.MethodGet:
				rg.GET(route.Path, route.HandlerFunc)
			case http.MethodPost:
				rg.POST(route.Path, route.HandlerFunc)
			case http.MethodPut:
				rg.PUT(route.Path, route.HandlerFunc)
			case http.MethodDelete:
				rg.DELETE(route.Path, route.HandlerFunc)
			// 添加其他HTTP方法
			default:
				// 处理未识别的HTTP方法
				rg.Any(route.Path, route.HandlerFunc)
			}
		}
	}
}
