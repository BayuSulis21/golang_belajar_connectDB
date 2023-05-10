package models

// object biller
type Biller struct {
	AggregatorProductID     string `json:"aggregatorProductID"  gorm:"column:aggregatorProductId"`
	ProductName             string `json:"productName"  gorm:"column:productName"`
	Denom                   string `json:"denom" gorm:"column:denom"`
	AggregatorName          string `json:"aggregatorName" gorm:"column:aggregatorName"`
	AggregatorProductCode   string `json:"aggregatorProductCode" gorm:"column:aggregatorProductCode"`
	ApplicationProductPrice string `json:"applicationProductPrice" gorm:"column:applicationProductPrice"`
	AggregatorPrice         string `json:"aggregatorPrice" gorm:"column:aggregatorPrice"`
	StatusBiller            string `json:"statusBiller" gorm:"column:statusBiller"`
	StatusProduct           string `json:"statusProduct" gorm:"column:statusProduct"`
}
