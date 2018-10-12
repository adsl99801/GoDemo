package main

import (
	Router "GoDemo/router"

	Dao "GoDemo/dao"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

//main o
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	dbInit()
	Router.Home(r)
	Router.Post(r)
	// r.Run(":8081")
	endless.ListenAndServe(":8081", r)
	// r.RunTLS(":8081")
}
func dbInit() {
	dao := new(Dao.MembersDao)
	dao.Init()
}
