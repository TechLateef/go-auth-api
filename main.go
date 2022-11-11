package main

import (
	"github.com/gin-gonic/gin"
	initializer "github.com/techlateef/go-auth/Initializer"
	"github.com/techlateef/go-auth/controller"
)

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDb()
	// initializer.SyncDatabase()

}

func main() {
	r := gin.Default()
	r.POST("/Signup", controller.Signup)

	r.Run()
}
