package xml

import (
	"encoding/xml"
	"fmt"
)

type receiver struct {
	XMLName                  xml.Name       `xml:"Receptor"`
	Nombre                   string         `xml:"Nombre"`
	Identification           identification `xml:"Identification"`
	IdentificacionExtranjero string         `xml:"IdentificacionExtranjero"`
	NombreComercial          string         `xml:"NombreComercial"`
	Ubicacion                ubicacion      `xml:"Ubicacion"`
	Telefono                 telefono       `xml:"Telefono"`
	Fax                      fax            `xml:"Fax"`
	CorreoElectronico        string         `xml:"CorreoElectronico"`
}

func (r *receiver) build(prueba string) {

	fmt.Println("sadasd")
}
