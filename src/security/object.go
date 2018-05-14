package security

import "encoding/xml"

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties --> signingCertificate --> Cert --> CertDigest
type certDigest struct {
	XMLName      xml.Name     `xml:"xades:CertDigest"`
	DigestMethod digestMethod `xml:"ds:DigestMethod"`
	DigestValue  string       `xml:"ds:DigestValue"`
}

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties --> signingCertificate --> Cert --> IssuerSerial
type issuerSerial struct {
	XMLName          xml.Name `xml:"xades:IssuerSerial"`
	X509IssuerName   string   `xml:"ds:X509IssuerName"`
	X509SerialNumber string   `xml:"ds:X509SerialNumber"`
}

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties --> signingCertificate --> Cert
type cert struct {
	XMLName      xml.Name `xml:"xades:Cert"`
	CertDigest   certDigest
	IssuerSerial issuerSerial
}

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties --> signingCertificate
type signingCertificate struct {
	XMLName xml.Name `xml:"xades:SigningCertificate"`
	Cert    cert
}

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties --> signaturePolicyIdentifier --> sigPolicyHash
type sigPolicyHash struct {
	XMLName      xml.Name     `xml:"xades:SigPolicyHash"`
	DigestMethod digestMethod `xml:"ds:DigestMethod"`
	DigestValue  string       `xml:"ds:DigestValue"`
}

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties --> signaturePolicyIdentifier --> sigPolicyID
type sigPolicyID struct {
	XMLName    xml.Name `xml:"xades:SigPolicyId"`
	Identifier string   `xml:"xades:Identifier"`
}

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties --> signaturePolicyIdentifier
type signaturePolicyIdentifier struct {
	XMLName       xml.Name `xml:"xades:SignaturePolicyIdentifier>xades:SignaturePolicyId"`
	SigPolicyID   sigPolicyID
	SigPolicyHash sigPolicyHash
}

// object --> qualifyingProperties --> signedProperties --> signedSignatureProperties
type signedSignatureProperties struct {
	XMLName                   xml.Name `xml:"xades:SignedSignatureProperties"`
	SigningTime               string   `xml:"xades:SigningTime"`
	SigningCertificate        signingCertificate
	SignaturePolicyIdentifier signaturePolicyIdentifier
}

// object --> qualifyingProperties --> signedProperties --> signedDataObjectProperties --> dataObjectFormat
type dataObjectFormat struct {
	XMLName         xml.Name `xml:"xades:DataObjectFormat"`
	ObjectReference string   `xml:"ObjectReference,attr"`
	MimeType        string   `xml:"xades:MimeType"`
}

// object --> qualifyingProperties --> signedProperties --> signedDataObjectProperties
type signedDataObjectProperties struct {
	XMLName          xml.Name `xml:"xades:SignedDataObjectProperties"`
	DataObjectFormat dataObjectFormat
}

// object --> qualifyingProperties --> signedProperties
type signedProperties struct {
	XMLName                    xml.Name `xml:"xades:SignedProperties"`
	ID                         string   `xml:"xmlns:Id,attr"`
	SignedSignatureProperties  signedSignatureProperties
	SignedDataObjectProperties signedDataObjectProperties
}

// object --> qualifyingProperties
type qualifyingProperties struct {
	XMLName          xml.Name `xml:"xades:QualifyingProperties"`
	XADES            string   `xml:"xmlns:xades,attr"`
	Target           string   `xml:"Target,attr"`
	SignedProperties signedProperties
}

// ROOT TAG
type object struct {
	XMLName              xml.Name `xml:"ds:Object"`
	QualifyingProperties qualifyingProperties
}

func buildObject() object {

	signedDOP := signedDataObjectProperties{
		DataObjectFormat: dataObjectFormat{
			ObjectReference: "#r-id-1",
			MimeType:        "application/octet-stream",
		},
	}

	signCerti := signingCertificate{
		Cert: cert{
			CertDigest:   certDigest{DigestMethod: digestMethod{Algorithm: "http://www.w3.org/2000/09/xmldsig#sha1"}, DigestValue: "PARAMETRO DE LA LLAVE"},
			IssuerSerial: issuerSerial{X509IssuerName: "PARAMETRO DE LA LLAVE", X509SerialNumber: "PARAMETRO DE LA LLAVE"},
		},
	}
	signaturePI := signaturePolicyIdentifier{
		SigPolicyID: sigPolicyID{Identifier: "https://tribunet.hacienda.go.cr/docs/esquemas/2016/v4.1/Resolucion_Comprobantes_Electronicos_DGT-R-48-2016.pdf"},

		SigPolicyHash: sigPolicyHash{DigestMethod: digestMethod{Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256"}, DigestValue: "PARAMETRO DE LA LLAVE"},
	}

	signedSP := signedSignatureProperties{
		SigningTime:               "2016-11-25T16:35:06Z",
		SigningCertificate:        signCerti,
		SignaturePolicyIdentifier: signaturePI,
	}

	signPro := signedProperties{
		ID: "xades-id-e34ffbff277e8d1432e864436aa11882",

		SignedSignatureProperties:  signedSP,
		SignedDataObjectProperties: signedDOP,
	}

	qualify := qualifyingProperties{
		XADES:            "http://uri.etsi.org/01903/v1.3.2#",
		Target:           "#id-34ffbff277e8d1432e864436aa11882",
		SignedProperties: signPro,
	}

	data := object{QualifyingProperties: qualify}

	return data

}
