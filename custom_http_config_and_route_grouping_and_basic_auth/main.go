package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	auth := gin.BasicAuth(gin.Accounts{
		"user" : "pass",
		"user2" : "pass2",
		"user3" : "pass3",
	//username : password that should match.
	})

	//Basic auth means the sended username and password should be in the authorized list that we have provided above.

	// Gin to use basic authentication mechanism,
	//And needs to have the username and password.

	router.GET("/getUrlData/:name/:age",getUrlData)

	admin := router.Group("/admin",auth)
	{
		admin.GET("/getData",getData)
		admin.POST("/postData",postData)
	}
	//This is route grouping.
	client := router.Group("/client")
	{
		client.GET("/getQueryString",getQueryString)
	}

	//Configuration for our server

	server := &http.Server{
		Addr: ":5000",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout : 10 * time.Second,
	}
 // With this we can pass custom http configuration for our server.
   server.ListenAndServe();
}

func getData(c *gin.Context){
	c.JSON(200,gin.H{
		"data":"Hi i am gin framework",
	})
}

func postData(c *gin.Context){
	body := c.Request.Body
	//This body needs to be in string format but is in io.ReadCloser Format . So for that
	fmt.Println(body)
	value, _ := io.ReadAll(body);
	fmt.Println(value)


	c.JSON(200,gin.H{
		"data" : "This is post request",
		"bodyData" : string(value),
	})
}

//http://localhost:5000/getQueryString?name=Mark&age=30

func getQueryString(c *gin.Context){
	name := c.Query("name")
	age := c.Query("age")
	//This would have been difficult in plain golang.
	c.JSON(200,gin.H{
		"data":"Hi I am GIN Framework",
		"name" : name,
		"age" : age,
	})
}

//http://localhost:5000/getUrlData/Mark/30

func getUrlData(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")
	//This would have been difficult in plain golang.
	c.JSON(200,gin.H{
		"data":"Hi I am GIN Framework",
		"name" : name,
		"age" : age,
	})
}
