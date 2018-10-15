package main

import (
	Dao "GoDemo/dao"
	Post "GoDemo/post"
	Router "GoDemo/router"
	Util "GoDemo/util"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

//main o
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	Dao.Init()
	authorized := r.Group("/post")
	authorized.Use(Util.TokenAuth())
	{
		authorized.GET("/test", Post.Test)
	}

	Router.Home(r)
	// r.Run(":8081")
	endless.ListenAndServe(":8081", r)
	// r.RunTLS(":8081")
}
