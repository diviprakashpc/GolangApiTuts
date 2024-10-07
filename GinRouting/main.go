package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//gin.New() will be same but default middleware wonnt get attached.

	//If we use like this

	// router.Use(gin.Logger()) // We can attach the default middlewares even after using gin.New.

	//This simply means that we are going to create a gin router with default middleware for logger and recovery. Logger will log every request by default for us and recovery in case of server error will automatic set errorcode to 500 and send it to client . This is one of use case.

	//Gin context contains every bit of info about req and res.

	// router.GET("/getData",func(c *gin.Context){
	// 	c.JSON(200,gin.H{
	// 		"data":"Hii I am GIN Framework",
	// 	})
	// })

	router.GET("/getData", getData)
	router.POST("/postData", postData)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)
	router.Run(":5000") // Default port is 8080. Pass custom port.
	// This run will calling listen and serve method of golang packaage by default. And clear the port automatic for us.
	//http.ListenAndServe(":5000",router)

	//: The colon here is important. Else It wont work.
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am gin framework",
	})
	// gin.H is a shorthand type for creating a map that represents a JSON object. Specifically, it is defined as gin.H = map[string]interface{}. This allows you to easily create JSON responses without having to define a struct for simple cases.
	// c.JSON(statusCode, data): This method sends a JSON response. The first parameter is the HTTP status code, and the second parameter is the data you want to include in the response, which can be a struct, a map (like gin.H), or any other type that can be marshaled to JSON.

}

func postData(c *gin.Context) {
	body := c.Request.Body
	//This body needs to be in string format but is in io.ReadCloser Format . So for that
	fmt.Println("Body received from ginContext \n", body)
	value, _ := io.ReadAll(body)
	fmt.Println("This is value after readAll \n", value)
	bodyData := string(value)
	fmt.Println("After parsing value from readAll to String \n", bodyData)
	c.JSON(200, gin.H{
		"data":     "This is post request",
		"bodyData": string(value),
	})

	// Body received from ginContext
	//  		&{0xc0000ae000 <nil> <nil> true true {0 0} false false false 0xac6f80}
	// This is value after readAll
	// 		[123 10 32 32 34 97 34 32 58 32 34 98 34 44 10 32 32 34 72 69 76 76 79 34 32 58 32 34 72 69 76 76 79 79 79 79 79 79 34 10 32 32 10 125]
	// After parsing value from readAll to String
	// 		{
	// 		"a" : "b",
	// 		"HELLO" : "HELLOOOOOO"

	// 		}
}

//http://localhost:5000/getQueryString?name=Mark&age=30

func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	//This would have been difficult in plain golang.
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
		"name": name,
		"age":  age,
	})
}

//http://localhost:5000/getUrlData/Mark/30

func getUrlData(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	//This would have been difficult in plain golang.
	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
		"name": name,
		"age":  age,
	})
}
