package xml

import "encoding/xml"

type receptor struct {
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

func addReceptor() receptor {

	newIdentification := identification{Tipo: "01", Numero: "111111"}
	newUbication := ubicacion{Provincia: "01", Canton: "02", Distrito: "03", Barrio: "asdasd", OtrasSenas: "asdsad"}
	newTelefono := telefono{CodigoPais: "506", NumTelefono: "444444"}
	newFax := fax{CodigoPais: "506", NumTelefono: "asdsa"}

	new := receptor{
		Nombre:            "sadsadsa",
		Identification:    newIdentification,
		NombreComercial:   "asdadsasa",
		Ubicacion:         newUbication,
		Telefono:          newTelefono,
		Fax:               newFax,
		CorreoElectronico: "ctreminiom079@gmail.com"}

	return new

}
