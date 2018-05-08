package xml

import (
	"encoding/xml"
	"fmt"
)

const xmlns = "https://tribunet.hacienda.go.cr/docs/esquemas/2017/v4.2/facturaElectronica"
const xmlnsDS = "http://www.w3.org/2000/09/xmldsig#"
const xmlnsXDS = "http://www.w3.org/2001/XMLSchema"
const xmlnsXSI = "http://www.w3.org/2001/XMLSchema-instance"
const xsi = "https://tribunet.hacienda.go.cr/docs/esquemas/2017/v4.2/facturaElectronica FacturaElectronica_V.4.2.xsd"

type bill struct {
	XMLName           xml.Name    `xml:"FacturaElectronica"`
	Xmlns             string      `xml:"xmlns,attr"`
	XmlnsDS           string      `xml:"xmlns:ds,attr"`
	XmlnsXSD          string      `xml:"xmlns:xsd,attr"`
	XmlnsXSI          string      `xml:"xmlns:xsi,attr"`
	Xsi               string      `xml:"xsi:schemaLocation,attr"`
	Clave             string      `xml:"Clave"`
	NumeroConsecutivo string      `xml:"NumeroConsecutivo"`
	FechaEmision      string      `xml:"FechaEmision"`
	Emisor            transmitter `xml:"Emisor"`
	Receptor          receiver    `xml:"Receptor"`
	CondicionVenta    string      `xml:"CondicionVenta"`
	PlazoCredito      string      `xml:"PlazoCredito"`
	MedioPago         string      `xml:"MedioPago"`
	DetalleServicio   detail      `xml:"DetalleServicio"`
}

func (b *bill) build() {

	fmt.Println("Asdsadsa")

	newReceiver := receiver{}
	newReceiver.build("Prueba")

	newTransmitter := transmitter{}
	newTransmitter.build("asdasd")

}

// CreateBill ...
func CreateBill(text string) {

	newBill := bill{}
	newBill.build()

}
