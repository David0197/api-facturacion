package key

import (
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func extractCertificate() string {

	certificates, err := ioutil.ReadFile("data.crt")

	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(certificates)

	return (base64.StdEncoding.EncodeToString(block.Bytes))

}
