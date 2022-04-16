package models

type Filter struct {
	Title     string `json:"title,omitempty"`
	Price     PriceFilter
	Target    string `json:",omitempty"`
	Category  string `json:",omitempty"`
	SortBy    string `json:",omitempty"`
	Condition string `json:",omitempty"`
}

type PriceFilter struct {
	Operator  string `json:"operator,omitempty"`
	FromValue int    `json:",omitempty"`
	ToValue   int    `json:",omitempty"`
}
