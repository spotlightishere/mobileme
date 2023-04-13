package main

import "github.com/gin-gonic/gin"

type AccountInfoResponse struct {
	GenericResponse
	ServicesAvailable []string `plist:"payload>servicesAvailable"`
}

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
	WriteResponse(response, c)
}
