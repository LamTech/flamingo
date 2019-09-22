package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccess(c *gin.Context,content interface{}){
	c.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"message": "success",
		"content": content,
	})
}

func JsonError(c *gin.Context,errCode uint,message interface{}){
	c.JSON(http.StatusOK, gin.H{
		"errCode": errCode,
		"message": message,
	})
}
