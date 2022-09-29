package webmain

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

var GIN *gin.Engine

var adminCookie string
var FLAG string

func Init() {
	GIN = gin.Default()
	GIN.LoadHTMLGlob("www/*")
	adminCookie = GetAdminCookie()
	FLAG = GetFlag()
	RegisterRoutes(GIN)
}

func GetAdminCookie() string {
	content, err := ioutil.ReadFile("./keys/admin_cookie.txt")
	if err != nil {
		log.Panicln("GetAdminCookie Error: ", err)
	}
	return strings.TrimSpace(string(content))
}

func GetFlag() string {
	content, err := ioutil.ReadFile("./keys/flag.txt")
	if err != nil {
		log.Panicln("GetFlag Error:", err)
	}
	return strings.TrimSpace(string(content))
}

func RegisterRoutes(g *gin.Engine) {
	adminRouter := g.Group("/admin", AdminAuth)
	adminRouter.GET("/view", AdminView)
	adminRouter.GET("/flag", AdminGetFlag)

	g.GET("/", UserSend)
	g.Any("/report", ReportToAdmin)
}
