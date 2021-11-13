package main

import (
	"fmt"
	"net/http"

	//sapi "GinServer/Todos/api"
	api "Golang/api"
	"Golang/api/helpers"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Gin-Gonic Server")
	helpers.InitDB()
	// append(todoList, {"Id":1,"Name":"Rohit"})
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", checkStatus())
	api.Init(router)
	s := &http.Server{
		Addr:    ":4700",
		Handler: router,
	}
	s.ListenAndServe()
}

func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
