package model

type AdListing struct {
	AdID       int64  `json:"ad_id"`
	ListID     int64  `json:"list_id"`
	CategoryID int64  `json:"category"`
	Body       string `json:"body"`
	Subject    string `json:"subject"`
	//AdParams   []AdListingParam `json:"ad_params"`
}

type AdListingParam struct {
	Address  string `json:"address"`
	Area     string `json:"area"`
	Ward     string `json:"ward"`
	Region   string `json:"region"`
	Price    string `json:"price"`
	CarBrand string `json:"car_brand"`
	CarModel string `json:"car_model"`
	GearBox  string `json:"gear_box"`
	MfDate   string `json:"mf_date"`
}
