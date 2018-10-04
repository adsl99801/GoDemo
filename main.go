package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": "u1", "value": "v1"})
	})
	r.GET("/data", func(c *gin.Context) {
		var engine *xorm.Engine
		var err error
		engine, err = xorm.NewEngine("postgres", "postgresql://localhost/lfodb?user=lfo&password=lfo")
		if err != nil {
			log.Fatal(err)
			return
		}
		engine.ShowSQL(true)
		c.String(http.StatusOK, "data")
	})
	r.Run(":8080")
}
