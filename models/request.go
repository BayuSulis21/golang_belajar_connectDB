package models

// rooter with parameter json
type Login struct {
	User     string `json:"user"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

// validation request
type HeaderDenomList struct {
	AppsID string `json:"appsID" binding:"required"`
}
type DenomList struct {
	BilleryType string `json:"type"`
}
