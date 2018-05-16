package security

import "encoding/xml"

type signatureValue struct {
	XMLName xml.Name `xml:"ds:SignatureValue"`
	ID      string   `xml:"Id,attr"`
	Value   string   `xml:",chardata"`
}

func buildSignatureValue() signatureValue {
	return signatureValue{ID: "value-id-e34ffbff277e8d1432e864436aa11882", Value: "PARAMETRO"}
}
