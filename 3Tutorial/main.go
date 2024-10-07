//What is middleware
//How to use middleware in go
//Different ways to write middleware functions in Go and why.
//Apply Middleware to routes, routes group and whole application at once.

/*

Middleware:-
In a request pipeline, a middleware has access to request and response object and the next function in the application in the request response cycle.

A application flow with middleware is like :

Request ===> Port ===> Connection Handler ==> App code (Main entry point) ==> Request Middleware ===> Code logic ===> Response Middleware ===> Port ===> Response to client.


*/

package main

import (
	"3Tutorial/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Authenticate)
	//This is using middleware at app level. Applied to all the endpoints.
	router.GET("/getData", getData)
	router.GET("/getData1", getData1)
	router.GET("/getData2", getData2)
	router.Run(":5000")
}

// Apply the middleware only to this endpoint
// router.GET("/getData2", middleware.Authenticate, getData2)

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am getData framework",
	})
}
func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am getData1 framework",
	})
}
func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am getData2 framework",
	})
}
