package models

//CertificateVariation ...
type CertificateVariation struct {
	Model
	CertificateID int64   `json:"certificate_id"`
	VariationID   int64   `json:"variation_id"`
	Value         float64 `json:"value"`
}
