package key

import (
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/ctreminiom/api-facturacion/src/conf"
	"golang.org/x/crypto/pkcs12"
)

// SplitCertificate ....
func SplitCertificate(patch string, password string) {
	p12Certificate, err := ioutil.ReadFile(patch)

	if err != nil {
		log.Fatal(err)
	}

	blocks, err := pkcs12.ToPEM(p12Certificate, password)

	if err != nil {
		log.Fatal(err)
	}

	var certificates, privateKey []byte

	for _, block := range blocks {

		if block.Type == "PRIVATE KEY" {
			privateKey = append(privateKey, pem.EncodeToMemory(block)...)
		}

		if block.Type == "CERTIFICATE" {
			certificates = append(certificates, pem.EncodeToMemory(block)...)
		}

	}

	ioutil.WriteFile(conf.GetVariable("certificate")+"data.crt", certificates, 0644)

	ioutil.WriteFile(conf.GetVariable("private")+"private.key", privateKey, 0644)

}
