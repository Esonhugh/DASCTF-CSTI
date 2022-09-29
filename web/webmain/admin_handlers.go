package webmain

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AdminAuth(c *gin.Context) {
	admin_cookie, err := c.Cookie("admin")
	if err != nil {
		c.JSON(401, gin.H{
			"error": "no admin Auth",
		})
		c.Abort()
		return
	}
	if admin_cookie != adminCookie {
		c.JSON(401, gin.H{
			"error": "no admin Auth! Bad Cookie",
		})
		c.Abort()
		return
	}
	c.Next()
}

func AdminView(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"name": CurrentAdminPayload,
	})
}

func AdminGetFlag(c *gin.Context) {
	log.Println("Got Flag")
	c.JSON(200, gin.H{
		"flag": FLAG,
	})
}
