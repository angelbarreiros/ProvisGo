package provisentities

import "net/url"

type PersonImageParams struct{}

func (p *PersonImageParams) ToURLValues() url.Values {
	return url.Values{}
}

type PersonImageResponse struct {
	Image string `json:"image"`
}
