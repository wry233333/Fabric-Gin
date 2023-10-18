package server

import (
	"io/ioutil"
	"net/http"

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
	value, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
	a := []string{"CreateAsset", string(value)}
	_, err = App.Set(a)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"data": value,
	})
}