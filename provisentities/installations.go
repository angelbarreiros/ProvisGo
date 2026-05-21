package provisentities

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

type Installation struct {
	IDInstalacion                       int     `json:"idInstalacion"`
	Catalogo                            int     `json:"catalogo"`
	Region                              string  `json:"region"`
	URLInstalacion                      string  `json:"urlInstalacion"`
	Longitud                            float64 `json:"longitud"`
	Latitud                             float64 `json:"latitud"`
	NomInstalacion                      string  `json:"nomInstalacion"`
	ExportacionURL                      *string `json:"exportacionUrl"`
	URLAppNoticias                      *string `json:"urlAppNoticias"`
	EsPasaporte                         bool    `json:"esPasaporte"`
	TieneModuloDeClasesColectivasCadena bool    `json:"tieneModuloDeClasesColectivasCadena"`
}

type InstallationsResponse struct {
	Installations []Installation `json:"installations"`
}
