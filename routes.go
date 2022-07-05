package main

import (
	"Oenone/handler/admin"
	"Oenone/handler/middleware"
	"Oenone/handler/public"
	"Oenone/handler/user"
	"github.com/gin-gonic/gin"
)

func BindRoutes(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware)

	r.GET("/auth/info", middleware.AuthMiddleware, user.GetInfoHandler)

	r.POST("/register", user.RegisterHandler)
	r.POST("/registerCaptcha", user.RegisterCaptchaHandler)
	r.POST("/login", user.LoginHandler)
	r.GET("/profile", middleware.AuthMiddleware, user.GetProfileHandler)

	// 搜索栏 提示
	//r.GET("/neighborhood", public.SearchNeighborhoodHandler)
	//r.GET("/house", public.SearchHouseHandler)
	r.GET("/search", public.SearchHandler)

	r.GET("/house", public.HouseDetailHandler)
	r.POST("/house", middleware.AuthMiddleware, user.PostHouseHandler)
	r.PUT("/house", middleware.AuthMiddleware, user.UpdateHouseHandler)
	r.DELETE("/house", middleware.AuthMiddleware, user.OfflineHouseHandler)

	r.GET("/counter", public.HouseCounterHandler)

	r.GET("/houseList", public.HouseListHandler)

	r.POST("/star", middleware.AuthMiddleware, user.StarHandler)
	r.DELETE("/star", middleware.AuthMiddleware, user.UnstarHandler)

	r.POST("/appointment", middleware.AuthMiddleware, user.MakeAppointmentHandler)
	r.DELETE("/appointment", middleware.AuthMiddleware, user.DelAppointmentHandler)

	r.PUT("/noticeAllRead", middleware.AuthMiddleware, user.NoticeAllReadHandler)
	r.GET("/noticeUnread", middleware.AuthMiddleware, user.GetUnreadNoticeAmountHandler)

	r.GET("/reviewHistories", middleware.AuthMiddleware, user.GetReviewHistoriesHandler)

	adminGroups := r.Group("/admin")
	{
		adminGroups.POST("login", admin.LoginHandler)
		adminGroups.GET("auth/info", admin.GetAdminInfoHandler)

		adminGroups.GET("manage", admin.ManageHandler)

		adminGroups.POST("review", admin.ReviewHandler)
		adminGroups.POST("admin", admin.AddAdminHandler)
		adminGroups.DELETE("admin", admin.DeleteAdminHandler)
		adminGroups.DELETE("user", admin.DeleteUserHandler)
	}
}
