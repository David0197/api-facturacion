package security

import "encoding/xml"

// SUB-TAG
type canonicalizationMethod struct {
	XMLName   xml.Name `xml:"ds:CanonicalizationMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

// SUB-TAG
type signatureMethod struct {
	XMLName   xml.Name `xml:"ds:SignatureMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

// SUB-TAG
type transform struct {
	XMLName   xml.Name `xml:"ds:Transform"`
	Algorithm string   `xml:"Algorithm,attr"`
	XPath     string   `xml:"ds:XPath,omitempty"`
}

// SUB-TAG
type transforms struct {
	XMLName   xml.Name `xml:"ds:Transforms"`
	Transform []transform
}

// SUB-TAG
type digestMethod struct {
	XMLName   xml.Name `xml:"ds:DigestMethod,omitempty"`
	Algorithm string   `xml:"Algorithm,attr"`
}

// SUB-TAG
type references struct {
	XMLName      xml.Name `xml:"ds:Reference"`
	ID           string   `xml:"Id,attr"`
	Type         string   `xml:"Type,attr"`
	URI          string   `xml:"URI,attr"`
	Transforms   transforms
	DigestMethod digestMethod
	DigestValue  string `xml:"ds:DigestValue,omitempty"`
}

// TAG
type signedInfo struct {
	XMLName                xml.Name `xml:"ds:SignedInfo"`
	CanonicalizationMethod canonicalizationMethod
	SignatureMethod        signatureMethod
	Reference              []references
}

func buildSignedInfo() signedInfo {

	var xPath []transform
	xPath = append(xPath, transform{Algorithm: "http://www.w3.org/TR/1999/REC-xpath-19991116", XPath: "not(ancestor-or-self::ds:Signature)"})
	xPath = append(xPath, transform{Algorithm: "http://www.w3.org/TR/1999/REC-xpath-19991116"})

	var c14n []transform
	c14n = append(c14n, transform{Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#"})

	var tags []references
	tags = append(tags, references{
		ID:           "r-id-1",
		Transforms:   transforms{Transform: xPath},
		DigestMethod: digestMethod{Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256"},
		DigestValue:  "PARAMETRO DE LA LLAVE",
	})

	tags = append(tags, references{
		Type:         "http://uri.etsi.org/01903#SignedProperties",
		URI:          "#xades-id-e34ffbff277e8d1432e864436aa11882",
		Transforms:   transforms{Transform: c14n},
		DigestMethod: digestMethod{Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256"},
		DigestValue:  "PARAMETRO DE LA LLAVE",
	})

	data := signedInfo{
		CanonicalizationMethod: canonicalizationMethod{Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#"},
		SignatureMethod:        signatureMethod{Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256"},
		Reference:              tags,
	}

	return data
}
