package provisEntities

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
