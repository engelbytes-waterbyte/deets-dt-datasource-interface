package deets

type AdditionalProperty struct {
	Name string      `json:"name"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (p AdditionalProperty) GetID() string {
	return p.Name + "-" + p.Type
}

const (
	PropertyTypeText   = "text"
	PropertyTypeNumber = "number"
	PropertyTypeImage  = "image"
)

type TextPropertyData string

type NumberPropertyData float64

type MediaPropertyData struct {
	URL    string  `json:"url"`
	Width  *int    `json:"width,omitempty"`
	Height *int    `json:"height,omitempty"`
	Alt    *string `json:"alt,omitempty"`
}
