package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haripriyanka/tinyurl-go/models"
)

func GetOriginalUrlByHash(c *gin.Context) {

	hash := c.Params.ByName("hash_id")

	if url, err := models.GetOriginalUrlByHash(hash); err == nil {
		c.JSON(http.StatusOK, url)

	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func CreateNewUrl(c *gin.Context) {

	orgurl := c.Params.ByName("originalurl")

	if url, err := models.CreateNewUrl(orgurl); err == nil {
		c.JSON(http.StatusOK, url)

	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
