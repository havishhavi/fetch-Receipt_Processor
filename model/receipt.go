package model

// item struct to represent different item names on the list
type Item struct {
	ShortDescription string `json:"shortDescrption"`
	Price            string `json:"price"`
}

// receipt that has all the req fields as mentioned in the example
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchasetime"`

	Total string `json:"total"`
	Items []Item `json:"items"`
}
