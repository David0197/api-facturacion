package key

import (
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/ctreminiom/api-facturacion/src/conf"
)

//Skeleton ...
type Skeleton struct {
	Header map[string]string
	Body   string
}

// PrivateKey ....
func PrivateKey() Skeleton {

	pemData, err := ioutil.ReadFile(conf.GetVariable("private") + "private.key")

	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(pemData)

	return Skeleton{block.Headers, base64.StdEncoding.EncodeToString(block.Bytes)}

}
