package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Result string

const (
	ResultSuccess             = "Success"
	ResultFailedConsistency   = "FailedConsistencyCheck"
	ResultNotAvailable        = "NotAvailable"
	ResultNotAuthorized       = "NotAuthorized"
	ResultNotImplemented      = "NotImplemented"
	ResultServiceError        = "FailedServiceError"
	ResultCSRFailedVerify     = "FailedCSRDidNotVerify"
	ResultNotSupportedAccount = "FailedNotSupportedForAccount"
	ResultNoExistingCSR       = "FailedNoExistingCSR"
	ResultPendingCSR          = "FailedPendingCSR"
	ResultNotAllowed          = "FailedNotAllowed"
	ResultParameterError      = "FailedParameterError"
	ResultCertAlreadyExists   = "FailedCertAlreadyExists"
	ResultAlreadyExists       = "FailedAlreadyExists"
	ResultFailed              = "Failed"
	ResultRedirected          = "SuccessRedirected"
	ResultQueued              = "SuccessQueued"
)

func archiveHandler(c *gin.Context) {
	username := c.GetString(UsernameKey)

	wrapper := NewXMLRPCWrapper(c)
	switch wrapper.MethodName() {
	case "archive.fetch":
		// TODO: Read from database or etc, not flat-file storage
		test, _ := os.ReadFile(fmt.Sprintf("./certs/%s/SharedServices_PKCS12.pfx", username))
		response := base64.StdEncoding.EncodeToString(test)

		members := []Member{
			{
				"resultCode",
				ResultSuccess,
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
	// Requests are made in the form of "/locate?someuser&type=dmSharedServices"
	// This should typically be a key-value pair. However, our username ("someuser")
	// has no value associated with it.
	// We loop through all keys and assume the first key without a value is our username.
	var username string
	for key, value := range c.Request.URL.Query() {
		if len(value) == 1 && value[0] == "" {
			username = key
		}
	}

	if username == "" {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	c.File(fmt.Sprintf("./certs/%s/userCertificate.pem", username))
}
