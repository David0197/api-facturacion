package logic

import (
	"encoding/xml"
	"fmt"
	"time"
)

const xmlns = "https://tribunet.hacienda.go.cr/docs/esquemas/2017/v4.2/facturaElectronica"
const xmlnsDS = "http://www.w3.org/2000/09/xmldsig#"
const xmlnsXDS = "http://www.w3.org/2001/XMLSchema"
const xmlnsXSI = "http://www.w3.org/2001/XMLSchema-instance"
const xsi = "https://tribunet.hacienda.go.cr/docs/esquemas/2017/v4.2/facturaElectronica FacturaElectronica_V.4.2.xsd"

/*Factura ....*/
type factura struct {
	XMLName           xml.Name `xml:"FacturaElectronica"`
	Xmlns             string   `xml:"xmlns,attr"`
	XmlnsDS           string   `xml:"xmlns:ds,attr"`
	XmlnsXSD          string   `xml:"xmlns:xsd,attr"`
	XmlnsXSI          string   `xml:"xmlns:xsi,attr"`
	Xsi               string   `xml:"xsi:schemaLocation,attr"`
	Clave             string   `xml:"Clave"`
	NumeroConsecutivo string   `xml:"NumeroConsecutivo"`
	FechaEmision      string   `xml:"FechaEmision"`
}

/*DoMagic ....*/
func (f factura) doMagic() {
	fmt.Println("FUNKA")
}

/*Create ....*/
func (f factura) create(clave string, numero string) factura {

	new := factura{
		Xmlns: xmlns, XmlnsDS: xmlnsDS, XmlnsXSD: xmlnsXDS, XmlnsXSI: xmlnsXSI, Xsi: xsi,
		Clave:             clave,
		NumeroConsecutivo: numero,
		FechaEmision:      time.Now().Format(time.RFC3339)}

	return new
}

/*Create ....*/
func Create() {

	fmt.Println("asdsa")

}
