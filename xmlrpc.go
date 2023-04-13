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
	doc *xmlquery.Node
}

// NewXMLRPCWrapper returns a new wrapper for the given context.
func NewXMLRPCWrapper(c *gin.Context) (XMLRPCWrapper, error) {
	// Parse the body initially.
	doc, err := xmlquery.Parse(c.Request.Body)
	if err != nil {
		return XMLRPCWrapper{}, err
	}

	return XMLRPCWrapper{ctx: c, doc: doc}, nil
}

// MethodName returns the parsed methodName. If not possible, it returns an empty string.
// TODO(spotlightishere): do we need to read passed value for parameters?
func (w *XMLRPCWrapper) MethodName() string {
	element := w.doc.SelectElement("//methodCall/methodName")
	if element == nil {
		// We may not have a methodName in this document.
		return ""
	}

	// Return the data of the methodName's first sibling
	// (i.e., its normal value).
	return element.FirstChild.Data
}

// ParseStringParams handles all parameters as strings.
// TODO(spotlightishere): This should likely be restructured, as not all parameters are strings.
func (w *XMLRPCWrapper) ParseStringParams() []string {
	elements := w.doc.SelectElements("//methodCall/params/param/value/string")

	var params []string
	for _, element := range elements {
		// Each "param" element within should have no children - just data.
		params = append(params, element.FirstChild.Data)
	}
	return params
}

// Member describes a usable key/value pair within XML-RPC responses.
type Member struct {
	Name  string      `xml:"name"`
	Value interface{} `xml:"value>string"`
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
