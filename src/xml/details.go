package xml

import (
	"encoding/xml"
)

type exonaration struct {
	XMLName           xml.Name `xml:"Exoneracion"`
	TipoDocumento     string   `xml:"TipoDocumento"`
	NumeroDocumento   string   `xml:"NumeroDocumento"`
	NombreInstitucion string   `xml:"NombreInstitucion"`
	FechaEmision      string   `xml:"FechaEmision"`
	MontoImpuesto     string   `xml:"MontoImpuesto"`
	PorcentajeCompra  int      `xml:"PorcentajeCompra"`
}

type tax struct {
	XMLName     xml.Name    `xml:"Impuesto"`
	Codigo      string      `xml:"Codigo"`
	Tarifa      int         `xml:"Tarifa"`
	Monto       int         `xml:"Monto"`
	Exoneracion exonaration `xml:"Exoneracion"`
}

type code struct {
	XMLName xml.Name `xml:"Codigo"`
	Tipo    string   `xml:"Tipo"`
	Codigo  string   `xml:"Codigo"`
}

type lineOfDetail struct {
	XMLName               xml.Name `xml:"LineaDetalle"`
	NumeroLinea           int      `xml:"NumeroLinea"`
	Codigo                code     `xml:"Codigo"`
	Cantidad              int      `xml:"Cantidad"`
	UnidadMedida          string   `xml:"UnidadMedida"`
	UnidadMedidaComercial string   `xml:"UnidadMedidaComercial"`
	Detalle               string   `xml:"Detalle"`
	PrecioUnitario        int      `xml:"PrecioUnitario"`
	MontoTotal            int      `xml:"MontoTotal"`
	MontoDescuento        int      `xml:"MontoDescuento"`
	NaturalezaDescuento   int      `xml:"NaturalezaDescuento"`
	SubTotal              int      `xml:"SubTotal"`
	Impuesto              tax      `xml:"Impuesto"`
	MontoTotalLinea       int      `xml:"MontoTotalLinea"`
}

type detail struct {
	XMLName      xml.Name       `xml:"DetalleServicio"`
	LineaDetalle []lineOfDetail `xml:"LineaDetalle"`
}

func (d *detail) AddLine() {

	line := lineOfDetail{} // ADD PARAMETERS

	d.LineaDetalle = append(d.LineaDetalle, line)
}
