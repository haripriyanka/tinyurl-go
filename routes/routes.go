package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {

	router := gin.Default()

	router.GET("/:hash_id", GetOriginalUrlByHash)

	router.PUT("/create/:originalurl", CreateNewUrl)

	router.Run()
}
