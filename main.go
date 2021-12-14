package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora/v3"
	"io"
	"io/ioutil"
	"log"
)

// RequestLoggerMiddleware is taken from
// https://github.com/gin-gonic/gin/issues/961#issuecomment-557931409
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		for name, value := range c.Request.Header {
			log.Print(aurora.Magenta(name), " ", "=>", " ", aurora.Yellow(value))
		}
		log.Println(aurora.Cyan(string(body)))
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(RequestLoggerMiddleware())

	// Certificates
	r.StaticFile("/dotMacCA.pem", "./certs/dotMacCA.pem")

	config := r.Group("/configurations")
	{
		config.GET("/internetservices/dotmacpreferencespane/1/clientConfiguration.plist", dotMacPrefPaneConfig)
		config.GET("/macosx/ichat/1/clientConfiguration.plist", ichatConfig)
	}

	web := r.Group("/WebObjects")
	{
		info := web.Group("/Info.woa/wa/DynamicUI")
		{
			info.POST("/dotMacPreferencesPaneMessage", paneMessage)
		}
	}

	r.GET("/locate", locateHandler)
	r.POST("/archive", archiveHandler)

	r.Run(":8080")
}
