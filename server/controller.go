package server

import (
	"github.com/gin-gonic/gin"
)

func AddAssets (c *gin.Context) {
	cid := c.Query("cid")
	a := []string{"set", "ID3", cid}
	_, err := App.Set(a)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"message": "Success",
	})
}

func QueryAssets (c *gin.Context) {
	cid := c.Query("cid")
	a := []string{"get", cid}
	ret, err := App.Get(a)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"message": ret,
	})
}