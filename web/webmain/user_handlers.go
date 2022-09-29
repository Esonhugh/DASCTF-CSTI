package webmain

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

var CurrentAdminPayload = "guest"

func UserSend(c *gin.Context) {
	if c.Request.Method == "POST" {
		c.AbortWithError(400, errors.New("bad Request"))
	}
	id := c.Query("name")
	if id == "" {
		id = "guest"
	}
	escaped := strings.Replace(id, "<", "[", -1)
	escaped = strings.Replace(escaped, ">", "]", -1)
	c.HTML(200, "index.html", gin.H{
		"name": escaped,
	})
}

func ReportToAdmin(c *gin.Context) {
	userpayload := c.PostForm("name")
	if strings.TrimSpace(userpayload) == "" {
		c.AbortWithError(404, errors.New("Empty url"))
	}
	CurrentAdminPayload = userpayload
	log.Println("Reported to admin: ", userpayload)
	c.Redirect(302, "/?name="+userpayload)
}
