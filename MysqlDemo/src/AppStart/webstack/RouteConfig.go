package AppStart

import (
	"github.com/gin-gonic/gin"
	"Services"
	"net/http"
)

func RouteConfig() {
	router := gin.Default()
	router.GET("/user/:name", Services.RouterPropertyHandler)
	router.GET("/user/:name/*action", Services.RouterPropertySecHandler)
	router.GET("/welcome", Services.RouterQueryStringHandler)
	router.POST("/form_post", Services.RouterPostBodyHandler)
	router.PUT("/post", Services.RouterPutQueryStringBodyHandler)
	router.POST("/login", Services.RouterPropertyBindHandler)
	router.POST("/login2", Services.RouterPropertyBindSecHandler)
	router.GET("/render", Services.RouterPropertyBindXMLHandler)
	router.GET("/redict", Services.RouterRedictURLHandler)
	router.GET("/redict/TargetHandler", Services.RouterRedictTargetURLHandler)
	v1 := router.Group("/v1")
	v1.GET("/grouphandler", Services.RouterGroupV1Handler)
	v2 := router.Group("/v2")
	v2.GET("/grouphandler", Services.RouterGroupV2Handler)

	//单个路由中间件(要放在全局中间件之前，或者在里面标识那个可以调用此中间件)
	router.GET("/singlemiddleware", SignalFilter, Services.RouterSingleMiddleWareHandler)

	//群组中间件
	gv1 := router.Group("/v1")
	gv1.GET("/groupFilterV1handler", GroupFilter, Services.RouterGroupFilterV1handler)

	//全局中间件
	//只要注册中间件的过程之前设置的路由，
	//将不会受注册的中间件所影响。只有注册了中间件一下代码的路由函数规则，才会被中间件装饰。
	router.Use(CAServerGlobalFilterHandler())
	{
	}
	router.Use(AuthorityFilterHandler())
	{
	}
	router.GET("/middleware", Services.RouterGlobalMiddleWareHandler)

	router.GET("/sync", Services.RouterSynchandler)
	router.GET("/async", Services.RouterAsyncchandler)
	//群组中间件
	http.ListenAndServe(":8000", router)

}
