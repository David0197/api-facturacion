package main

import (
	"github.com/ctreminiom/api-facturacion/src/security"
	"github.com/ctreminiom/api-facturacion/src/xml"
)

func main() {

	xml.CreateBill("Carlos")

	security.Sign()

}
