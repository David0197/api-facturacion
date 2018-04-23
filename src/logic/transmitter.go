package logic

import "encoding/xml"

type fax struct {
	XMLName     xml.Name `xml:"Fax"`
	CodigoPais  string   `xml:"CodigoPais"`
	NumTelefono string   `xml:"NumTelefono"`
}

type telefono struct {
	XMLName     xml.Name `xml:"Telefono"`
	CodigoPais  string   `xml:"CodigoPais"`
	NumTelefono string   `xml:"NumTelefono"`
}

type ubicacion struct {
	XMLName    xml.Name `xml:"Ubicacion"`
	Provincia  string   `xml:"Provincia"`
	Canton     string   `xml:"Canton"`
	Distrito   string   `xml:"Distrito"`
	Barrio     string   `xml:"Barrio"`
	OtrasSenas string   `xml:"OtrasSenas"`
}

type identification struct {
	XMLName xml.Name `xml:"Identification"`
	Tipo    string   `xml:"Tipo"`
	Numero  string   `xml:"Numero"`
}

type emisor struct {
	XMLName           xml.Name         `xml:"Emisor"`
	Nombre            string           `xml:"Nombre"`
	Identification    []identification `xml:"Identification"`
	NombreComercial   string           `xml:"NombreComercial"`
	Ubicacion         []ubicacion      `xml:"Ubicacion"`
	Telefono          []telefono       `xml:"Telefono"`
	Fax               []fax            `xml:"Fax"`
	CorreoElectronico string           `xml:"CorreoElectronico"`
}

func (e emisor) add() emisor {

	e = emisor{Nombre: "Carlos", NombreComercial: "asdasd", CorreoElectronico: "asdasd"}

	return e
}
