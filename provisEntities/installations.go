package provisEntities

import "net/url"

// InstallationsParams represents filter parameters for the Installations endpoint
// All fields are pointers to allow optional/nullable values
type InstallationsParams struct {
	Extended        *bool `json:"extended,omitempty"`
	IncludeInactive *bool `json:"includeInactive,omitempty"`
}

// ToURLValues serializes the filter params to url.Values
// Only non-nil fields are added to the result
func (p *InstallationsParams) ToURLValues() url.Values {
	values := url.Values{}
	if p == nil {
		return values
	}
	// Installations endpoint currently uses hardcoded query params
	// Extended is handled in the URL path
	return values
}
