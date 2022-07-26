package main

type Root struct {
	Limit float64 `json:"limit"`
	Count float64 `json:"count"`
	Total float64 `json:"total"`
	Pages float64 `json:"pages"`
	Deals []Deals `json:"deals"`
}

type Deals struct {
	ModelYear             string   `json:"modelYear"`
	FuelDescription       string   `json:"fuelDescription"`
	Images                []string `json:"images"`
	GearDescription       string   `json:"gearDescription"`
	Km                    float64  `json:"km"`
	ManufacturingYear     string   `json:"manufacturingYear"`
	DealText              string   `json:"dealText"`
	Price                 float64  `json:"price"`
	ColorDescription      string   `json:"colorDescription"`
	Equipments            []string `json:"equipments"`
	PlateLastNumber       string   `json:"plateLastNumber"`
	SellerCityDescription string   `json:"sellerCityDescription"`
	SellerName            string   `json:"sellerName"`
}
