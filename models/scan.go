package models

/*
Scan represents a scan api response
*/
type Scan struct {
	ID       string `json:"scan_id"`
	Resource string
	Code     int    `json:"response_code"`
	Message  string `json:"verbose_msg"`
}
