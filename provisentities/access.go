package provisentities

import (
	"encoding/json"
	"net/url"
)

type AccessByDateParams struct {
	StartDatetime string `json:"startDatetime,omitempty"`
	EndDateTime   string `json:"endDateTime,omitempty"`
}

func (p *AccessByDateParams) ToURLValues() url.Values {
	values := url.Values{}
	if p == nil {
		return values
	}
	if p.StartDatetime != "" {
		values.Set("startDatetime", p.StartDatetime)
	}
	if p.EndDateTime != "" {
		values.Set("endDateTime", p.EndDateTime)
	}
	return values
}

type AccessByDateResponse struct {
	Code          int           `json:"code,omitempty"`
	Message       string        `json:"message,omitempty"`
	CustomMessage CustomMessage `json:"customMessage,omitempty"`
	Details       string        `json:"details,omitempty"`
	Accesses      []Access      `json:"result"`
}

func (r *AccessByDateResponse) UnmarshalJSON(data []byte) error {
	var accesses []Access
	if err := json.Unmarshal(data, &accesses); err == nil {
		r.Accesses = accesses
		return nil
	}

	type alias AccessByDateResponse
	var wrapped alias
	if err := json.Unmarshal(data, &wrapped); err != nil {
		return err
	}
	*r = AccessByDateResponse(wrapped)
	return nil
}

type Access struct {
	IDAcceso      int    `json:"idAcceso"`
	FechaHora     string `json:"fechaHora"`
	IDPersona     int    `json:"idPersona"`
	Nombre        string `json:"nombre"`
	IDNivel       int    `json:"idNivel"`
	IDZona        any    `json:"idZona"`
	Sentido       bool   `json:"sentido"`
	EntradaSalida string `json:"entradaSalida"`
}
