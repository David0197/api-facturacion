package security

import "encoding/xml"

type keyInfo struct {
	XMLName xml.Name `xml:"ds:KeyInfo"`
	X509    string   `xml:"ds:X509Data>ds:X509Certificate"`
}

func buildKeyInfo() keyInfo {
	return keyInfo{X509: "PARAMETRO DE LA LLAVE"}
}
