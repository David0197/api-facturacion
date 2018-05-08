package xml

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
