package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountInfoResponse struct {
	GenericResponse
	ServicesAvailable []string `plist:"payload>servicesAvailable"`
}

// accountInfo handles the WebObjects accountInfo endpoint.
func accountInfo(c *gin.Context) {
	// TODO: Authentication
	// (lol)

	// We'll enable all services by default.
	response := AccountInfoResponse{
		GenericResponse: GenericResponse{
			StatusCode: ResponseStatusSuccess,
		},
		ServicesAvailable: []string{
			"iDisk",
			"iSync",
			"Backup",
			"iChatEncryption",
			"SharingCertificate",
			"BTMM",
			"Email",
			"DotMacMail",
			"WebHosting",
		},
	}
	WriteXML(c, response)
}

// accountInfoRPC handles the XML-RPC accountInfo endpoint.
func accountInfoRPC(c *gin.Context) {
	wrapper, err := NewXMLRPCWrapper(c)
	if err != nil {
		// TODO: Proper error format
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	// This is a little odd. We should have three parameters here:
	// 1. Username
	// 2. Password
	// 3. Query type
	params := wrapper.ParseStringParams()

	// TODO: Authentication
	//username := params[0]
	//password := params[1]
	queryType := params[2]

	if queryType != "daysLeftUntilExpiration" {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	wrapper.Response([]Member{
		{
			Name:  "daysLeftUntilExpiration",
			Value: 25,
		},
	})
}
