package main

import (
	"fmt"

	"github.com/ctreminiom/api-facturacion/src/key"
)

func main() {

	//xml.CreateBill("Carlos")
	//security.Sign()

	//key.SplitCertificate(conf.GetVariable("p12"), conf.GetVariable("p12_password"))

	//fmt.Println(key.PrivateKey())

	//fmt.Println(key.PrivateKey().Header)

	fmt.Println(key.Certificates())

}
