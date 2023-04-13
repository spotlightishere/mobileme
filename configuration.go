package main

import (
	"github.com/gin-gonic/gin"
)

func dotMacPrefPaneConfig(c *gin.Context) {
	paneMessage := globalConfig.baseDomain("/WebObjects/Info.woa/wa/DynamicUI/dotMacPreferencesPaneMessage")

	config := GenericConfig{
		ConfigurationUrl: c.Request.URL.String(),
		Services: map[string]string{
			"dotMacPreferencesPaneMessage":         paneMessage,
			"dotMacPreferencesPaneMessageVersion2": paneMessage,
		},
	}
	Write(config, c)
}

func ichatConfig(c *gin.Context) {
	type Language struct {
		Default string `plist:"default"`
		EN      string `plist:"en"`
		JA      string `plist:"ja"`
		FR      string `plist:"fr"`
		DE      string `plist:"de"`
	}

	type ChatConfig struct {
		GenericConfig
		LocalizedURLs map[string]Language `plist:"localizedURLs"`
	}

	config := ChatConfig{
		GenericConfig: GenericConfig{
			ConfigurationUrl: c.Request.URL.String(),
			Services: map[string]string{
				"accountInfo": "https://www.mac.com/WebObjects/Info.woa/wa/Query/accountInfo",
			},
		},
		LocalizedURLs: map[string]Language{
			"createTrialURL": {
				Default: "http://support.apple.com/kb/TS4058",
				EN:      "http://support.apple.com/kb/TS4058",
				JA:      "http://support.apple.com/kb/TS4058?viewlocale=ja_JP",
				FR:      "http://support.apple.com/kb/TS4058?viewlocale=fr_FR",
				DE:      "http://support.apple.com/kb/TS4058?viewlocale=de_DE",
			},
			"createSubscriberURL": {
				Default: "https://secure.me.com/wo/WebObjects/Signup.woa/wa/subscribe",
				EN:      "https://secure.me.com/wo/WebObjects/Signup.woa/wa/subscribe?lang=en&cty=US",
				JA:      "https://secure.me.com/wo/WebObjects/Signup.woa/wa/subscribe?lang=ja&cty=JP",
				FR:      "https://secure.me.com/wo/WebObjects/Signup.woa/wa/subscribe?lang=fr&cty=FR",
				DE:      "https://secure.me.com/wo/WebObjects/Signup.woa/wa/subscribe?lang=de&cty=DE",
			},
			"revokeCertificateURL": {
				Default: "https://secure.me.com/account/",
				EN:      "https://secure.me.com/account/en",
				JA:      "https://secure.me.com/account/ja",
				FR:      "https://secure.me.com/account/fr",
				DE:      "https://secure.me.com/account/de",
			},
			"iChatEncryptionExpiredURL": {
				Default: "http://www.mac.com/WebObjects/Welcome",
				EN:      "http://www.mac.com/WebObjects/Welcome.woa/wa/default?lang=en&cty=US",
				JA:      "http://www.mac.com/WebObjects/Welcome.woa/wa/default?lang=ja&cty=JP",
				FR:      "http://www.mac.com/WebObjects/Welcome.woa/wa/default?lang=fr&cty=FR",
				DE:      "http://www.mac.com/WebObjects/Welcome.woa/wa/default?lang=de&cty=DE",
			},
			"iChatEncryptionLearnMoreURL": {
				Default: "http://www.apple.com/mobileme/features/mac.html",
				EN:      "http://www.apple.com/mobileme/features/mac.html",
				JA:      "http://www.apple.com/mobileme/features/mac.html",
				FR:      "http://www.apple.com/mobileme/features/mac.html",
				DE:      "http://www.apple.com/mobileme/features/mac.html",
			},
			"forgottenPasswordURL": {
				Default: "https://iforgot.apple.com/cgi-bin/WebObjects/DSiForgot.woa/wa/iforgot?app_type=ext&app_id=114",
				EN:      "https://iforgot.apple.com/cgi-bin/WebObjects/DSiForgot.woa/wa/iforgot?app_type=ext&app_id=114&language=en",
				JA:      "https://iforgot.apple.com/cgi-bin/WebObjects/DSiForgot.woa/wa/iforgot?app_type=ext&app_id=114&language=ja",
				FR:      "https://iforgot.apple.com/cgi-bin/WebObjects/DSiForgot.woa/wa/iforgot?app_type=ext&app_id=114&language=fr",
				DE:      "https://iforgot.apple.com/cgi-bin/WebObjects/DSiForgot.woa/wa/iforgot?app_type=ext&app_id=114&language=de",
			},
		},
	}
	Write(config, c)
}
