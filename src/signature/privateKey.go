package signature

import (
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func extractPrivateKey() string {

	pemData, err := ioutil.ReadFile("key.pem")

	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(pemData)

	return base64.StdEncoding.EncodeToString(block.Bytes)

}
