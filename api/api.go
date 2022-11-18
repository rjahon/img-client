package api

import (
	"github.com/rjahon/img-client/api/docs"
	"github.com/rjahon/img-client/api/handlers"

	"github.com/rjahon/img-client/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger API
// @description Image Client
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	docs.SwaggerInfo.Title = cfg.ServiceName

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, platform-id, X-Forwarded-Host, Referer")
	config.AllowHeaders = append(config.AllowHeaders, "*")

	r.Use(gin.Logger(), gin.Recovery(), cors.New(config))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/img", h.CreateImg)
	r.GET("/img/:id", h.GetImgByID)
	r.GET("/img", h.GetImgs)

	return
}
