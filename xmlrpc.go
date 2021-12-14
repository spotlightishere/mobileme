package main

import (
	"encoding/xml"
	"github.com/antchfx/xmlquery"
	"github.com/gin-gonic/gin"
	"net/http"
)

// XMLRPCWrapper assists in wrapping a *gin.Context to read and write XML-RPC requests.
type XMLRPCWrapper struct {
	ctx *gin.Context
}

// NewXMLRPCWrapper returns a new wrapper for the given context.
func NewXMLRPCWrapper(c *gin.Context) XMLRPCWrapper {
	return XMLRPCWrapper{ctx: c}
}

// MethodName returns the parsed methodName. If not possible, it returns an empty string.
// TODO(spotlightishere): do we need to read passed value for parameters?
func (w *XMLRPCWrapper) MethodName() string {
	doc, err := xmlquery.Parse(w.ctx.Request.Body)
	if err != nil {
		// Hopefully an invalid request.
		return ""
	}

	element := doc.SelectElement("//methodCall/methodName")
	if element == nil {
		// We may not have a methodName in this document.
		return ""
	}

	// Return the data of the methodName's first sibling
	// (i.e., its normal value).
	return element.FirstChild.Data
}

// Member describes a usable key/value pair within XML-RPC responses.
type Member struct {
	Name  string `xml:"name"`
	Value string `xml:"value>string"`
}

// MethodResponse hacks together a lot of logic to provide a valid XML-RPC response.
type MethodResponse struct {
	XMLName xml.Name `xml:"methodResponse"`
	Members []Member `xml:"params>param>value>struct>member"`
}

// Response encodes a valid XML-RPC response for the passed members.
func (w *XMLRPCWrapper) Response(passed []Member) {
	result := MethodResponse{
		Members: passed,
	}

	w.ctx.Writer.WriteString(xml.Header)
	w.ctx.XML(http.StatusOK, result)
}
