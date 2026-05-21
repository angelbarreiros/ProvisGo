package provisentities

type LoginRequest struct {
	IdInstallation int    `json:"idInstallation"`
	NifNickEmail   string `json:"NifNickEmail"`
	Password       string `json:"password"`
}
