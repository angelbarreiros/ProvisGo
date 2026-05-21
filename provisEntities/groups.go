package provisEntities

import "net/url"

// GroupsParams represents filter parameters for the Groups endpoint
// All fields are pointers to allow optional/nullable values
type GroupsParams struct {
	AcceptLanguage *string `json:"acceptLanguage,omitempty"`
}

// ToURLValues serializes the filter params to url.Values
// Only non-nil fields are added to the result
func (p *GroupsParams) ToURLValues() url.Values {
	values := url.Values{}
	if p == nil {
		return values
	}

	return values
}

type GroupsResponse struct {
	Code          int           `json:"code"`
	Message       string        `json:"message"`
	CustomMessage CustomMessage `json:"customMessage"`
	Details       string        `json:"details"`
	Persons       []GroupPerson `json:"result"`
}

type CustomMessage struct {
	ID         interface{}            `json:"id"`
	Parameters map[string]interface{} `json:"parameters"`
}

type GroupPerson struct {
	IdPersona     int         `json:"idPersona"`
	IdInstalacion int         `json:"idInstalacion"`
	Codigo        int         `json:"codigo"`
	Nombre        string      `json:"nombre"`
	Apellidos     string      `json:"apellidos"`
	Nivel         interface{} `json:"nivel"`
	TipoCliente   string      `json:"tipoCliente"`
	Asistencia    bool        `json:"asistencia"`
}
