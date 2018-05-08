package xml

/*

const xmlns = "https://tribunet.hacienda.go.cr/docs/esquemas/2017/v4.2/facturaElectronica"
const xmlnsDS = "http://www.w3.org/2000/09/xmldsig#"
const xmlnsXDS = "http://www.w3.org/2001/XMLSchema"
const xmlnsXSI = "http://www.w3.org/2001/XMLSchema-instance"
const xsi = "https://tribunet.hacienda.go.cr/docs/esquemas/2017/v4.2/facturaElectronica FacturaElectronica_V.4.2.xsd"

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
	Emisor            emisor   `xml:"Emisor"`
	Receptor          receptor `xml:"Receptor"`
	CondicionVenta    string   `xml:"CondicionVenta"`
	PlazoCredito      string   `xml:"PlazoCredito"`
	MedioPago         string   `xml:"MedioPago"`
	DetalleServicio   detalle  `xml:"DetalleServicio"`
}

func Create() {

	emisor := add()

	receptor := addReceptor()

	deta := addDetails()

	new := factura{
		Xmlns:             xmlns,
		XmlnsDS:           xmlnsDS,
		XmlnsXSD:          xmlnsXDS,
		XmlnsXSI:          xmlnsXSI,
		Xsi:               xsi,
		Clave:             "11111",
		NumeroConsecutivo: "22222",
		FechaEmision:      time.Now().Format(time.RFC3339),
		Emisor:            emisor,
		Receptor:          receptor,
		CondicionVenta:    "01",
		PlazoCredito:      "02",
		MedioPago:         "01",
		DetalleServicio:   deta}

	output, err := xml.MarshalIndent(new, "  ", "    ")

	if err != nil {

		fmt.Printf("error: %v\n", err)

	}

	os.Stdout.Write(output)
}

*/
