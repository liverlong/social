package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	userGroup := engine.Group("/user")

	{
		// todo implements interface
		userGroup.POST("/send_sms_code")
		userGroup.POST("/login")
	}

	engine.Run(":8080")
}
