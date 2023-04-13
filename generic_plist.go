package main

import (
	"github.com/gin-gonic/gin"
	"howett.net/plist"
)

// GenericConfig represents the base necessary keys for a configuration:
// the configuration URL, and a key/value array of available services and their URLs.
type GenericConfig struct {
	ConfigurationUrl string            `plist:"configurationURL"`
	Services         map[string]string `plist:"services"`
}

// ResponseStatus describes valid response types.
type ResponseStatus string

const (
	ResponseStatusSuccess = "success"
	ResponseStatusError   = "authorizationFailed"
)

// GenericResponse is a basic response for many requests.
type GenericResponse struct {
	StatusCode ResponseStatus `plist:"statusCode"`
}

// Write encodes the given struct as an old-style plist.
func Write(from interface{}, to *gin.Context) {
	e := plist.NewEncoderForFormat(to.Writer, plist.OpenStepFormat)
	err := e.Encode(from)
	if err != nil {
		// TODO: not panic
		panic(err)
	}
}

func WriteResponse(from interface{}, to *gin.Context) {
	e := plist.NewEncoderForFormat(to.Writer, plist.OpenStepFormat)
	err := e.Encode(from)
	if err != nil {
		// TODO: not panic
		panic(err)
	}
}
