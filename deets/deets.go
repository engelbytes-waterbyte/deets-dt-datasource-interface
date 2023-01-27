package deets

type AdditionalProperty struct {
	Name string      `json:"name"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type TextPropertyData string

type NumberPropertyData float64

type ImagePropertyData struct {
	URL    string  `json:"url"`
	Width  *int    `json:"width,omitempty"`
	Height *int    `json:"height,omitempty"`
	Alt    *string `json:"alt,omitempty"`
}
