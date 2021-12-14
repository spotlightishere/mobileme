package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ResultCode string

const (
	ResultCodeSuccess = "Success"
)

func archiveHandler(c *gin.Context) {
	wrapper := NewXMLRPCWrapper(c)
	switch wrapper.MethodName() {
	case "archive.fetch":
		// TODO: handle per user
		test, _ := ioutil.ReadFile("./certs/SharedServices_PKCS12.pfx")
		response := base64.StdEncoding.EncodeToString(test)

		members := []Member{
			{
				"resultCode",
				ResultCodeSuccess,
			},
			{
				"resultBody",
				response,
			},
		}
		wrapper.Response(members)
	default:
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
}

func locateHandler(c *gin.Context) {
	// TODO(spotlightishere): Properly parse username/etc to return accordingly
	test, _ := ioutil.ReadFile("./certs/userCertificate.pem")
	c.Writer.Write(test)
}
