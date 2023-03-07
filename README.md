# deets-dt-datasource-interface

This is a swagger specification that acts as an interface for developing custom data sources for [Deets](https://deets.waterbyte.club)

## Getting Started

You can implement this swagger specification in any framework and programming language of your choice.

For the properties of a product, stick to the swagger documentation, special properties are:

- images: This property is required. If your datasource does not provide product images, use a stringified empty array as values (e.g. `[]`). If your datasource does provide product images, use a stringified array of MediaPropertyData objects.
- additional_properties: The additional properties of the product, you may use any supported property type here, adhere to the AdditionalProperty Type.

## Types

### MediaPropertyData

<!-- type MediaPropertyData struct {
URL string `json:"url"`
Width *int `json:"width,omitempty"`
Height *int `json:"height,omitempty"`
Alt \*string `json:"alt,omitempty"`
} -->

```json
{
  "url": "string,required",
  "width": 0,
  "height": 0,
  "alt": "string"
}
```

### AdditionalProperty

```json
{
  "name": "string,required",
  "type": "string,required",
  "data": "any,required"
}
```
