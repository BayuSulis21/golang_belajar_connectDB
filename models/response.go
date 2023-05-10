package models

type DenomListresponse struct {
	AppsID          string   `json:"appsID"`
	BillerType      string   `json:"type"`
	ResponseCode    string   `json:"responseCode"`
	ResponseMessage string   `json:"responseMessage"`
	ListData        []Biller `json:"listData"`
}

type ErrorResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
