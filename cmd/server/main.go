package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"log"
)

func homePage(context *gin.Context) {
	docs.SwaggerInfo.Title = "Swaggo Example API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "valoru-software.com"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	log.Printf("Query params: %v\n", context.Request.URL.Query())
	log.Printf("Request params: %v\n", context.Params)
}

func startServer() {
	router := gin.Default()
	router.GET("/test/:id", homePage)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run()
	if err != nil {
		log.Fatalf("What the heck: %s", err.Error())
	}
}

func main() {
	startServer()
}
