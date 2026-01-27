package provisEntities

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
