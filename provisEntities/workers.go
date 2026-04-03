package provisEntities

import "net/url"

// WorkersParams represents filter parameters for the Workers endpoint
// All fields are pointers to allow optional/nullable values
type WorkersParams struct {
	IncludeInactive *bool `json:"includeInactive,omitempty"`
	IncludeStaff    *bool `json:"includeStaff,omitempty"`
}

// ToURLValues serializes the filter params to url.Values
// Only non-nil fields are added to the result
func (p *WorkersParams) ToURLValues() url.Values {
	values := url.Values{}
	if p == nil {
		return values
	}
	// Workers endpoint currently doesn't have optional query params
	// besides installationId which is handled separately
	return values
}

type ProvisWorker struct {
	IDTrabajador     int     `json:"idTrabajador"`
	Nombre           string  `json:"nombre"`
	Apellidos        string  `json:"apellidos"`
	AccesoStaffEmail *string `json:"accesoStaffEmail"`
	Email            *string `json:"email"`
	NIF              *string `json:"nif"`
	IDInstalacion    int     `json:"idInstalacion"`
	Imagen           *string `json:"imagen"`
}
type ProvisWorkers struct {
	Workers []*ProvisWorker `json:"workers"`
}
