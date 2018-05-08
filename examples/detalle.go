/*package xml

import (
	"encoding/xml"
)

type exoneracion struct {
	XMLName           xml.Name `xml:"Exoneracion"`
	TipoDocumento     string   `xml:"TipoDocumento"`
	NumeroDocumento   string   `xml:"NumeroDocumento"`
	NombreInstitucion string   `xml:"NombreInstitucion"`
	FechaEmision      string   `xml:"FechaEmision"`
	MontoImpuesto     string   `xml:"MontoImpuesto"`
	PorcentajeCompra  int      `xml:"PorcentajeCompra"`
}

type impuesto struct {
	XMLName     xml.Name    `xml:"Impuesto"`
	Codigo      string      `xml:"Codigo"`
	Tarifa      int         `xml:"Tarifa"`
	Monto       int         `xml:"Monto"`
	Exoneracion exoneracion `xml:"Exoneracion"`
}

type codigo struct {
	XMLName xml.Name `xml:"Codigo"`
	Tipo    string   `xml:"Tipo"`
	Codigo  string   `xml:"Codigo"`
}

type lineaDetalle struct {
	XMLName               xml.Name `xml:"LineaDetalle"`
	NumeroLinea           int      `xml:"NumeroLinea"`
	Codigo                codigo   `xml:"Codigo"`
	Cantidad              int      `xml:"Cantidad"`
	UnidadMedida          string   `xml:"UnidadMedida"`
	UnidadMedidaComercial string   `xml:"UnidadMedidaComercial"`
	Detalle               string   `xml:"Detalle"`
	PrecioUnitario        int      `xml:"PrecioUnitario"`
	MontoTotal            int      `xml:"MontoTotal"`
	MontoDescuento        int      `xml:"MontoDescuento"`
	NaturalezaDescuento   int      `xml:"NaturalezaDescuento"`
	SubTotal              int      `xml:"SubTotal"`
	Impuesto              impuesto `xml:"Impuesto"`
	MontoTotalLinea       int      `xml:"MontoTotalLinea"`
}

type detalle struct {
	XMLName      xml.Name       `xml:"DetalleServicio"`
	LineaDetalle []lineaDetalle `xml:"LineaDetalle"`
}

func agregar() lineaDetalle {

	newCodigo := codigo{Tipo: "01", Codigo: "02"}

	newExo := exoneracion{
		TipoDocumento:     "asdasd",
		NumeroDocumento:   "asdasd",
		NombreInstitucion: "sdasds",
		FechaEmision:      "asdasd",
		MontoImpuesto:     "asdasd",
		PorcentajeCompra:  4324342}

	type exoneracion struct {
		XMLName           xml.Name `xml:"Exoneracion"`
		TipoDocumento     string   `xml:"TipoDocumento"`
		NumeroDocumento   string   `xml:"NumeroDocumento"`
		NombreInstitucion string   `xml:"NombreInstitucion"`
		FechaEmision      string   `xml:"FechaEmision"`
		MontoImpuesto     string   `xml:"MontoImpuesto"`
		PorcentajeCompra  int      `xml:"PorcentajeCompra"`
	}

	newImpuesto := impuesto{
		Codigo:      "01",
		Tarifa:      3242,
		Monto:       32423,
		Exoneracion: newExo}

	new := lineaDetalle{
		NumeroLinea:           1,
		Codigo:                newCodigo,
		Cantidad:              1,
		UnidadMedida:          "asdas",
		UnidadMedidaComercial: "asdasd",
		Detalle:               "sadasdas",
		PrecioUnitario:        2331,
		MontoTotal:            3324,
		NaturalezaDescuento:   3242,
		SubTotal:              32424,
		Impuesto:              newImpuesto}

	return new

}

func addDetails() detalle {

	data := []lineaDetalle{agregar()}

	data = append(data, agregar())

	new := detalle{LineaDetalle: data}

	return new


}

*/
