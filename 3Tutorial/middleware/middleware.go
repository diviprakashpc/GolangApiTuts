package middleware

import "github.com/gin-gonic/gin"

//This is a request middleware code. We will write authentication middleware that will authenticate the value of auth token in the header.

//There is a specific syntax to follow while writing a middleware to tell gin framework that it is a middlware.

func Authenticate(c *gin.Context){
	if !(c.Request.Header.Get("Token")=="auth"){
		c.AbortWithStatusJSON(500,gin.H{
			"Message": " Token not present",
		})
		return; // Else it will execute the next middleware.
	}
	c.Next(); // If it has a token named auth.
}