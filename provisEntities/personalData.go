package provisEntities

import "net/url"

// PersonalDataParams represents filter parameters for the PersonalData endpoint
// All fields are pointers to allow optional/nullable values
type PersonalDataParams struct {
	IncludeAddresses   *bool `json:"includeAddresses,omitempty"`
	IncludePaymentInfo *bool `json:"includePaymentInfo,omitempty"`
	IncludeContactInfo *bool `json:"includeContactInfo,omitempty"`
}

// ToURLValues serializes the filter params to url.Values
// Only non-nil fields are added to the result
func (p *PersonalDataParams) ToURLValues() url.Values {
	values := url.Values{}
	if p == nil {
		return values
	}
	// PersonalData endpoint currently doesn't have optional query params
	// besides installationid and idPersona which are handled separately
	return values
}

type Persona struct {
	IDPersona                                   int                `json:"idPersona"`
	IDInstalacion                               int                `json:"idInstalacion"`
	IDTipoDeCliente                             int                `json:"idTipoDeCliente"`
	Codigo                                      int                `json:"codigo"`
	NIF                                         string             `json:"nif"`
	TarjetaAcceso                               interface{}        `json:"tarjetaAcceso"`
	Email                                       string             `json:"email"`
	Nombre                                      string             `json:"nombre"`
	Apellidos                                   string             `json:"apellidos"`
	FechaDeNacimiento                           string             `json:"fechaDeNacimiento"`
	Telefono1                                   string             `json:"telefono1"`
	Telefono2                                   string             `json:"telefono2"`
	Movil                                       string             `json:"movil"`
	Direccion                                   string             `json:"direccion"`
	Localidad                                   string             `json:"localidad"`
	CP                                          string             `json:"cp"`
	IDProvincia                                 int                `json:"idProvincia"`
	CodBanco                                    string             `json:"codBanco"`
	CodSucursal                                 string             `json:"codSucursal"`
	DC                                          string             `json:"dc"`
	NumeroDeCuenta                              string             `json:"numeroDeCuenta"`
	IBAN                                        string             `json:"iban"`
	BIC                                         string             `json:"bic"`
	Nick                                        string             `json:"nick"`
	Sexo                                        string             `json:"sexo"`
	EnviarMailings                              int                `json:"enviarMailings"`
	EnviarSMS                                   int                `json:"enviarSMS"`
	EnviarEmails                                int                `json:"enviarEmails"`
	Contrasenia                                 interface{}        `json:"contrasenia"`
	UltimaFechaDeAlta                           string             `json:"ultimaFechaDeAlta"`
	UltimaFechaDeBaja                           string             `json:"ultimaFechaDeBaja"`
	PagosPersonas                               interface{}        `json:"pagosPersonas"`
	EstadosPersonas                             interface{}        `json:"estadosPersonas"`
	CuotasPersonas                              interface{}        `json:"cuotasPersonas"`
	MatriculasPersonas                          interface{}        `json:"matriculasPersonas"`
	AccesosPersonas                             interface{}        `json:"accesosPersonas"`
	ClasesColectivasPersonas                    interface{}        `json:"clasesColectivasPersonas"`
	PersonasDirecciones                         []PersonaDireccion `json:"personasDirecciones"`
	FechaNotificacionLOPD                       interface{}        `json:"fechaNotificacionLOPD"`
	PermitirNotificacionesComercialesLOPD       bool               `json:"permitirNotificacionesComercialesLOPD"`
	TipoDeAceptacionLOPD                        int                `json:"tipoDeAceptacionLOPD"`
	FechaNotificacionLOPDEnBaja                 interface{}        `json:"fechaNotificacionLOPDEnBaja"`
	PermitirNotificacionesComercialesLOPDEnBaja bool               `json:"permitirNotificacionesComercialesLOPDEnBaja"`
	TipoDeAceptacionLOPDEnBaja                  int                `json:"tipoDeAceptacionLOPDEnBaja"`
	EnviarNotificaciones                        bool               `json:"enviarNotificaciones"`
	EnviarMailingsComerciales                   bool               `json:"enviarMailingsComerciales"`
	EnviarSMSComerciales                        bool               `json:"enviarSMSComerciales"`
	EnviarEmailsComerciales                     bool               `json:"enviarEmailsComerciales"`
	EnviarNotificacionesComerciales             bool               `json:"enviarNotificacionesComerciales"`
	PermitirCompartirEnRedesSociales            bool               `json:"permitirCompartirEnRedesSociales"`
	PermitirCompartirDatosATerceros             bool               `json:"permitirCompartirDatosATerceros"`
	ImagenPersona                               interface{}        `json:"imagenPersona"`
}

type PersonaDireccion struct {
	IDPersonaDireccion  int         `json:"idPersonaDireccion"`
	IDPersona           int         `json:"idPersona"`
	TipoDeDireccion     int         `json:"tipoDeDireccion"`
	IDTipoDeVia         string      `json:"idTipoDeVia"`
	NombreVia           string      `json:"nombreVia"`
	Numero              string      `json:"numero"`
	Piso                string      `json:"piso"`
	Puerta              string      `json:"puerta"`
	Bloque              string      `json:"bloque"`
	Escalera            string      `json:"escalera"`
	DireccionConvertida bool        `json:"direccionConvertida"`
	DireccionConErrores bool        `json:"direccionConErrores"`
	Geolocalizar        bool        `json:"geolocalizar"`
	Personas            interface{} `json:"personas"`
}
