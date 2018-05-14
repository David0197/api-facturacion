package security

import (
	"encoding/xml"
	"log"
	"os"
)

// Sign function returns a XML signed
func Sign() {

	newSignature := signature{
		SignedInfo:     buildSignedInfo(),
		SignatureValue: buildSignatureValue(),
		KeyInfo:        buildKeyInfo(),
		Object:         buildObject(),
	}

	output, err := xml.MarshalIndent(newSignature, "  ", "    ")

	if err != nil {
		log.Panic(err)
	}

	os.Stdout.Write(output)
}

type signature struct {
	XMLName        xml.Name       `xml:"ds:Signature"`
	ID             string         `xml:"Id,attr"`
	XMLNS          string         `xml:"xmlns:ds,attr"`
	SignedInfo     signedInfo     `xml:"ds:SignedInfo"`
	SignatureValue signatureValue `xml:"ds:SignatureValue"`
	KeyInfo        keyInfo        `xml:"ds:KeyInfo"`
	Object         object         `xml:"ds:Object"`
}

func (s *signature) Merge(text string) {

}
