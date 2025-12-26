package constants

import "net/textproto"

var (
	XServiceName  = textproto.CanonicalMIMEHeaderKey("x-service-name") // no usages
	XApiKey       = textproto.CanonicalMIMEHeaderKey("x-api-key")      // no usages
	XRequestAt    = textproto.CanonicalMIMEHeaderKey("x-request-at")   // no usages
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")  // no usages
)
