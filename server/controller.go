package server

import (
	"github.com/gin-gonic/gin"
)

type Asset struct {
	AppraisedValue int    `json:"AppraisedValue"`
	Color          string `json:"Color"`
	ID             string `json:"ID"`
	Owner          string `json:"Owner"`
	Size           int    `json:"Size"`
}

func QueryAssets (c *gin.Context) {
	ID := c.Param("ID")
	a := []string{"QueryAssets", ID}
	ret, err := App.Get(a)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data": ret,
	})
}

func AddAssets (c *gin.Context) {
	var asset Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.AbortWithError(500, err)
	}
	a := []string{"CreateAsset", asset.ID, asset.Color, string(asset.Size), asset.Owner, string(asset.AppraisedValue)}
	_, err := App.Set(a)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data": nil,
	})
}