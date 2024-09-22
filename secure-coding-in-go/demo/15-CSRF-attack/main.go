package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	csrf "github.com/utrack/gin-csrf"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.GET("/protected", func(c *gin.Context) {
		c.String(200, csrf.GetToken(c))
	})

	r.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

	r.Run(":8080")
}


/*
In Postman, 

First, you need to make a GET request to the /protected endpoint to get a CSRF token.

Make a POST request with the requested CSRF token embedded the header X-CSRF-TOKEN.

Note, on the client-side, handling the CSRF token, including saving it and embedding it in the request header (such as X-CSRF-TOKEN) is typically handled by the client-side JavaScript. 

Note the client-side and server-side codes are working together to access the token. Other processes from other webpages can't get access the token. Because they don't have the access to the response. 
*/