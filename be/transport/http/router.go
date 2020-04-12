package http

// 注册路由
/**
如需要登陆态验证，需要放到/api/空间下
如使用token验证，需要放到/openapi/空间下
*/
func (r *runner) registerRouter() {
	api := r.e.Group("/se/api")
	api.POST("/register", r.bs.RegisterAccount)
	api.POST("/login", r.bs.Login)

	openApi := r.e.Group("/openapi")
	openApi.Use(JWTAuth())
	openApi.GET("/account/info", r.bs.QuerySingleAccountInfo)
	openApi.POST("/account", r.bs.ModifyAccountInfo)
}
