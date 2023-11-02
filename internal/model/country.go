package model

type Country struct {
	Ip          string `json:"ip"`
	ISOCode     string `json:"iso_code"`
	CountryName string `json:"country_name"`
	// and so on
}
