package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora/v3"
	"io"
	"log"
	"net/http"
)

// RequestLoggerMiddleware is taken from
// https://github.com/gin-gonic/gin/issues/961#issuecomment-557931409
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)

		for name, value := range c.Request.Header {
			log.Print(aurora.Magenta(name), " ", "=>", " ", aurora.Yellow(value))
		}
		log.Println(aurora.Cyan(string(body)))
		c.Next()
	}
}

const (
	UsernameKey = "username"
)

func RequireAuthorization() gin.HandlerFunc {
	// TODO: Have proper authentication
	// Currently assumes someuser:testing1234.
	authorization := "Basic c29tZXVzZXI6dGVzdGluZzEyMzQ="

	return func(c *gin.Context) {
		given := c.GetHeader("Authorization")
		if given != authorization {
			c.Header("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO: properly update username
		c.Set(UsernameKey, "someuser")
	}
}

// check checks if, er, err errs.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// RoutingTable is a way to perform rudimentary host-based routing.
type RoutingTable struct {
	routes map[string]http.Handler
}

// NewRoutingTable creates a new routing table.
func NewRoutingTable() *RoutingTable {
	return &RoutingTable{
		routes: make(map[string]http.Handler),
	}
}

// HandleHost creates a new router for the provided domain.
func (t *RoutingTable) HandleHost(host string) *gin.Engine {
	handler := gin.Default()
	handler.Use(RequestLoggerMiddleware())
	t.routes[host] = handler
	return handler
}

// ServeHTTP performs host-based routing.
func (t *RoutingTable) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	handler, ok := t.routes[host]

	// If we don't know this host, throw a 404.
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Otherwise, fall back to Gin.
	handler.ServeHTTP(w, r)
}

func main() {
	LoadGlobalConfig()

	// We need three domains:
	// Our first is for "configuration.apple.com", which we leverage for several things.
	// The second is for our actual domain with API endpoints.
	// The third is for "certinfo.<base domain>", leveraged for certificate-related requests.
	routingTable := NewRoutingTable()
	config := routingTable.HandleHost("configuration.apple.com")
	configGroup := config.Group("/configurations")
	{
		configGroup.GET("/internetservices/dotmacpreferencespane/1/clientConfiguration.plist", dotMacPrefPaneConfig)
		configGroup.GET("/macosx/ichat/1/clientConfiguration.plist", ichatConfig)
		configGroup.GET("/internetservices/issupport/2_27a4cv2b6061/clientConfig.plist", issupportConfig)
	}

	// API endpoints and Web Objects
	endpoints := routingTable.HandleHost(globalConfig.BaseDomain)
	webGroup := endpoints.Group("/WebObjects")
	{
		infoGroup := webGroup.Group("/Info.woa/wa")
		{
			infoGroup.POST("/DynamicUI/dotMacPreferencesPaneMessage", paneMessage)
			infoGroup.POST("/Query/accountInfo", accountInfo)
			infoGroup.POST("/XMLRPC/accountInfo", accountInfoRPC)
		}
	}

	// Certificates
	certificates := routingTable.HandleHost("certinfo." + globalConfig.BaseDomain)
	certificates.StaticFile("/dotMacCA.pem", "./certs/dotMacCA.pem")
	certificates.GET("/locate", locateHandler)
	certificates.POST("/archive", RequireAuthorization(), archiveHandler)

	// Here we go!
	http.ListenAndServe(globalConfig.ListenAddress, routingTable)
}
