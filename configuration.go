package main

import (
	"github.com/gin-gonic/gin"
)

// dotMacPrefPaneConfig provides the configuration utilized for the MobileMe preference pane.
// https://configuration.apple.com/internetservices/dotmacpreferencespane/1/clientConfiguration.plist
func dotMacPrefPaneConfig(c *gin.Context) {
	config := GenericConfig{
		ConfigurationUrl: c.Request.URL.String(),
		Services: map[string]string{
			"dotMacPreferencesPaneMessage":         baseDomain("/WebObjects/Info.woa/wa/DynamicUI/dotMacPreferencesPaneMessage"),
			"dotMacPreferencesPaneMessageVersion2": baseDomain("/WebObjects/Info.woa/wa/DynamicUI/dotMacPreferencesPaneMessage"),
		},
	}
	WriteOldStyle(c, config)
}

// LanguageConfig provides specific values for localizations.
type LanguageConfig struct {
	Default string `plist:"default"`
	EN      string `plist:"en"`
	JA      string `plist:"ja"`
	FR      string `plist:"fr"`
	DE      string `plist:"de"`
}

// ichatConfig provides the configuration used with iChat and other services.
// https://configuration.apple.com/macosx/ichat/1/clientConfiguration.plist
func ichatConfig(c *gin.Context) {
	type ChatConfig struct {
		GenericConfig
		LocalizedURLs map[string]LanguageConfig `plist:"localizedURLs"`
	}

	config := ChatConfig{
		GenericConfig: GenericConfig{
			ConfigurationUrl: c.Request.URL.String(),
			Services: map[string]string{
				"accountInfo": baseDomain("/WebObjects/Info.woa/wa/Query/accountInfo"),
			},
		},
		LocalizedURLs: map[string]LanguageConfig{
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
	WriteOldStyle(c, config)
}

// issupportConfig provides the configuration necessary for ISSupport.framework to function.
// https://configuration.apple.com/configurations/internetservices/issupport/2_27a4cv2b6061/clientConfig.plist
func issupportConfig(c *gin.Context) {
	config := map[string]interface{}{
		"realmSupportEnabled": 1,
		"applicationOverrides": map[string]map[string]string{
			"fs2c": {
				"iDiskURL":  subdomain("fileservices"),
				"sIDiskURL": idiskDomain(),
				"deltaURL":  subdomain("delta"),
			},
		},
		"accountInfoURL":         baseDomain("/WebObjects/Info.woa/wa/XMLRPC/accountInfo"),
		"accountInfoURL2":        baseDomain("/WebObjects/Info.woa/wa/Query/accountInfo"),
		"referralLookupURL":      "http://homepage.mac.com/dotmackitsupport/Referrals",
		"iDiskURL":               idiskDomain(),
		"sIDiskURL":              idiskDomain(),
		"mobilePublishConfigURL": baseDomain("/WebObjects/MobileServices.woa/xmlrpc"),
		"commentsURL":            baseDomain("/WebObjects/WSComments.woa/xmlrpc"),
		"indexingURL":            subdomain("webservices"),
		"indexingBatchSize":      50,
		"commentsBatchSize":      50,
		"signUpURL": LanguageConfig{
			// We don't appear to need the default key.
			// However, it's ignored - that allows us to reuse this struct :)
			Default: "http://www.apple.com/mobileme/share-your-world/index.html",
			EN:      "http://www.apple.com/mobileme/share-your-world/index.html",
			DE:      "http://www.apple.com/de/mobileme/share-your-world/index.html",
			FR:      "http://www.apple.com/fr/mobileme/share-your-world/index.html",
			JA:      "http://www.apple.com/jp/mobileme/share-your-world/index.html",
		},
		"otherParameters": map[string]string{},
	}
	WriteOldStyle(c, config)
}
