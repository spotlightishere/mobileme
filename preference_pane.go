package main

import (
	"github.com/gin-gonic/gin"
)

const (
	messageHTML = `Created on %@, mail storage %@, disk storage %@`
)

// PaneMessage describes all necessary fields within an authentication request.
type PaneMessage struct {
	Version    int        `plist:"version"`
	StatusCode string     `plist:"statusCode"`
	Service    PaneStatus `plist:"service"`

	CanBuyMore        bool     `plist:"canBuyMore"`
	CreationDate      string   `plist:"createDateString"`
	DiskStorage       int      `plist:"iDiskStorageInMB"`
	MailStorage       int      `plist:"mailStorageInMB"`
	PublicFolder      string   `plist:"publicFolder"`
	ServicesAvailable []string `plist:"servicesAvailable"`
	UpgradeURL        string   `plist:"upgradeURL"`

	HTML              string   `plist:"messageHTML"`
	SubstitutionOrder []string `plist:"substitutionOrder"`
}

// PaneStatus describes valid authorization response types.
type PaneStatus string

const (
	PaneStatusSuccess = "success"
	PaneStatusError   = "authorizationFailed"
)

func paneMessage(c *gin.Context) {
	plist := PaneMessage{
		Version:    2,
		StatusCode: PaneStatusSuccess,
		Service:    "dotMacPreferencesPaneMessageVersion2",

		CanBuyMore:   true,
		CreationDate: "2021-12-11",
		DiskStorage:  1000.0,
		MailStorage:  1000.0,
		PublicFolder: "https://idisk.mac.com/spotlight",
		// We'll enable all services by default.
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
		UpgradeURL: "https://www.mac.com",

		HTML: messageHTML,
		SubstitutionOrder: []string{
			"createDateString",
			"mailStorageInMB",
			"iDiskStorageInMB",
		},
	}
	Write(plist, c)
}
