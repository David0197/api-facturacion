package xml

import (
	"encoding/xml"
	"fmt"
)

type transmitter struct {
	XMLName           xml.Name       `xml:"Emisor"`
	Nombre            string         `xml:"Nombre"`
	Identification    identification `xml:"Identification"`
	NombreComercial   string         `xml:"NombreComercial"`
	Ubicacion         ubicacion      `xml:"Ubicacion"`
	Telefono          telefono       `xml:"Telefono"`
	Fax               fax            `xml:"Fax"`
	CorreoElectronico string         `xml:"CorreoElectronico"`
}

func (r *transmitter) build(prueba string) {

	fmt.Println("sadasd")
}
