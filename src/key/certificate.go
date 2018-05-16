package key

import (
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/ctreminiom/api-facturacion/src/conf"
)

// Template ....
type Template struct {
	Headers     []map[string]string
	Certificate string
}

// Certificates ....
func Certificates() Template {

	certificate, err := ioutil.ReadFile(conf.GetVariable("certificate") + "data.crt")

	if err != nil {
		log.Fatal(err)
	}

	var blocks [][]byte
	var header []map[string]string

	for {
		var blockStructure *pem.Block
		blockStructure, certificate = pem.Decode(certificate)

		if blockStructure == nil {
			break
		}

		if blockStructure.Type == "CERTIFICATE" {
			blocks = append(blocks, blockStructure.Bytes)
			header = append(header, blockStructure.Headers)
		}
	}

	return Template{header, base64.StdEncoding.EncodeToString(blocks[0])}

}
