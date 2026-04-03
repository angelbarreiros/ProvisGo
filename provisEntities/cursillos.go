package provisEntities

import (
	"net/url"
	"strconv"
)

// CursillosParams represents filter parameters for the Cursillos endpoint
// All fields are pointers to allow optional/nullable values
type CursillosParams struct {
	Origen                           *string  `json:"origen,omitempty"`
	IDPersona                        *int     `json:"idpersona,omitempty"`
	FechaInicio                      *string  `json:"fechainicio,omitempty"`
	Edad                             *int     `json:"edad,omitempty"`
	DiaSemana                        *int     `json:"diasemana,omitempty"`
	FranjaHoraria                    *string  `json:"franjahoraria,omitempty"`
	PrecioMaximo                     *float64 `json:"preciomaximo,omitempty"`
	Agrupacion                       *string  `json:"agrupacion,omitempty"`
	TipoDeCliente                    *string  `json:"tipodecliente,omitempty"`
	Cursillo                         *string  `json:"cursillo,omitempty"`
	Centro                           *string  `json:"centro,omitempty"`
	Plazas                           *int     `json:"plazas,omitempty"`
	IncluirVisiblesSoloDesdePrograma *bool    `json:"incluirVisiblesSoloDesdePrograma,omitempty"`
}

// ToURLValues serializes the filter params to url.Values
// Only non-nil fields are added to the result
func (p *CursillosParams) ToURLValues() url.Values {
	values := url.Values{}
	if p == nil {
		return values
	}
	if p.Origen != nil {
		values.Set("origen", *p.Origen)
	}
	if p.IDPersona != nil {
		values.Set("idpersona", strconv.Itoa(*p.IDPersona))
	}
	if p.FechaInicio != nil {
		values.Set("fechainicio", *p.FechaInicio)
	}
	if p.Edad != nil {
		values.Set("edad", strconv.Itoa(*p.Edad))
	}
	if p.DiaSemana != nil {
		values.Set("diasemana", strconv.Itoa(*p.DiaSemana))
	}
	if p.FranjaHoraria != nil {
		values.Set("franjahoraria", *p.FranjaHoraria)
	}
	if p.PrecioMaximo != nil {
		values.Set("preciomaximo", strconv.FormatFloat(*p.PrecioMaximo, 'f', -1, 64))
	}
	if p.Agrupacion != nil {
		values.Set("agrupacion", *p.Agrupacion)
	}
	if p.TipoDeCliente != nil {
		values.Set("tipodecliente", *p.TipoDeCliente)
	}
	if p.Cursillo != nil {
		values.Set("cursillo", *p.Cursillo)
	}
	if p.Centro != nil {
		values.Set("centro", *p.Centro)
	}
	if p.Plazas != nil {
		values.Set("plazas", strconv.Itoa(*p.Plazas))
	}
	if p.IncluirVisiblesSoloDesdePrograma != nil {
		values.Set("incluirVisiblesSoloDesdePrograma", strconv.FormatBool(*p.IncluirVisiblesSoloDesdePrograma))
	}
	return values
}

// CuotasParams represents filter parameters for the Cuotas endpoint
// All fields are pointers to allow optional/nullable values
type CuotasParams struct {
	FechaInicio *string `json:"fechaInicio,omitempty"`
	FechaFin    *string `json:"fechaFin,omitempty"`
	PersonID    *string `json:"personid,omitempty"`
}

// ToURLValues serializes the filter params to url.Values
// Only non-nil fields are added to the result
func (p *CuotasParams) ToURLValues() url.Values {
	values := url.Values{}
	if p == nil {
		return values
	}
	if p.FechaInicio != nil {
		values.Set("fechaInicio", *p.FechaInicio)
	}
	if p.FechaFin != nil {
		values.Set("fechaFin", *p.FechaFin)
	}
	if p.PersonID != nil {
		values.Set("personid", *p.PersonID)
	}
	return values
}

type Horario struct {
	DiaDeLaSemana int    `json:"diadelasemana"`
	HoraInicio    string `json:"horainicio"`
	HoraFin       string `json:"horafin"`
	IDZona        int    `json:"idzona"`
	Zona          string `json:"zona"`
	IDMonitor     int    `json:"idmonitor"`
	Monitor       string `json:"monitor"`
}

type Grupo struct {
	IDCursilloGrupo      int       `json:"idcursillogrupo"`
	Descripcion          string    `json:"descripcion"`
	Capacidad            int       `json:"capacidad"`
	IDSubVarios          int       `json:"idsubvarios"`
	SubVarios            string    `json:"subvarios"`
	ListaEspera          bool      `json:"listaEspera"`
	Completo             bool      `json:"completo"`
	Ocupacion            int       `json:"ocupacion"`
	OcupacionListaEspera int       `json:"ocupacionListaEspera"`
	Horario              []Horario `json:"horario"`
}

type LabelsCursillosList struct {
	ApplicationPath                  any  `json:"applicationPath"`
	Horario                          any  `json:"horario"`
	Inscribirse                      any  `json:"inscribirse"`
	Cursillo                         any  `json:"cursillo"`
	Grupo                            any  `json:"grupo"`
	ImporteCursilloDosPuntos         any  `json:"importeCursilloDosPuntos"`
	Matricula                        any  `json:"matricula"`
	TotalDosPuntos                   any  `json:"totalDosPuntos"`
	NombreCentro                     any  `json:"nombreCentro"`
	EresUsuario                      any  `json:"eresUsuario"`
	AccedeAlCentro                   any  `json:"accedeAlCentro"`
	QuieresUnirteAlCentro            any  `json:"quieresUnirteAlCentro"`
	DateDeAlta                       any  `json:"dateDeAlta"`
	Cancelar                         any  `json:"cancelar"`
	MesPunto                         any  `json:"mesPunto"`
	MesesPunto                       any  `json:"mesesPunto"`
	CursilloCompletoPeriodicidad     any  `json:"cursilloCompletoPeriodicidad"`
	ObtenerCursillos                 any  `json:"obtenerCursillos"`
	SeHaProducidoUnError             any  `json:"seHaProducidoUnError"`
	DesdeEl                          any  `json:"desdeEl"`
	HastaEl                          any  `json:"hastaEl"`
	PuedesApuntarteALaListaDeEspera  any  `json:"puedesApuntarteALaListaDeEspera"`
	InscripcionCubreLaTotalidad      any  `json:"inscripcionCubreLaTotalidad"`
	EstaInscripciónSeraValidaDurante any  `json:"estaInscripciónSeraValidaDurante"`
	ReservarCursillo                 any  `json:"reservarCursillo"`
	Ocupacion                        any  `json:"ocupacion"`
	Confirmar                        any  `json:"confirmar"`
	Submenu                          any  `json:"submenu"`
	PuedePagarConTarjeta             bool `json:"puedePagarConTarjeta"`
	PuedePagarConCredito             bool `json:"puedePagarConCredito"`
	PuedePagarConBizum               bool `json:"puedePagarConBizum"`
	PagoPorReferencia                bool `json:"pagoPorReferencia"`
	ListaTokens                      any  `json:"listaTokens"`
	TarjetaTerminadaEn               any  `json:"tarjetaTerminadaEn"`
	PagarConTarjeta                  any  `json:"pagarConTarjeta"`
	PagarConCredito                  any  `json:"pagarConCredito"`
	PagarEnElCentro                  any  `json:"pagarEnElCentro"`
	PagarConBizum                    any  `json:"pagarConBizum"`
	PagarConBono                     any  `json:"pagarConBono"`
	Reservado                        any  `json:"reservado"`
	Error                            any  `json:"error"`
}

type CursillosFiltroModel struct {
	Origen                           any     `json:"origen"`
	IDPersona                        int     `json:"idpersona"`
	FechaInicio                      string  `json:"fechainicio"`
	Edad                             any     `json:"edad"`
	DiaSemana                        any     `json:"diasemana"`
	FranjaHoraria                    any     `json:"franjahoraria"`
	PrecioMaximo                     float64 `json:"preciomaximo"`
	Agrupacion                       any     `json:"agrupacion"`
	TipoDeCliente                    any     `json:"tipodecliente"`
	Cursillo                         any     `json:"cursillo"`
	Centro                           any     `json:"centro"`
	Plazas                           int     `json:"plazas"`
	CursillosFiltrados               any     `json:"cursillosfiltrados"`
	GruposFiltrados                  any     `json:"gruposfiltrados"`
	DiasSemanaFiltrados              any     `json:"diasSemanafiltrados"`
	FranjasHorariasFiltrados         any     `json:"franjasHorariasfiltrados"`
	TiposDeClientesFiltrados         any     `json:"tiposDeclientesfiltrados"`
	IncluirVisiblesSoloDesdePrograma bool    `json:"incluirVisiblesSoloDesdePrograma"`
}

type Cursillo struct {
	CursillosFiltroModel                         CursillosFiltroModel `json:"cursillosFiltroModel"`
	CursillosFiltroLista                         any                  `json:"cursillosFiltroLista"`
	LabelsCursillosList                          LabelsCursillosList  `json:"labelsCursillosList"`
	IDInstalacion                                int                  `json:"idinstalacion"`
	IDTipoDeCliente                              int                  `json:"idtipodecliente"`
	EdadDesde                                    int                  `json:"edaddesde"`
	EdadHasta                                    int                  `json:"edadhasta"`
	IDCursillo                                   int                  `json:"idcursillo"`
	Nombre                                       string               `json:"nombre"`
	Descripcion                                  string               `json:"descripcion"`
	IDAgrupacion                                 int                  `json:"idagrupacion"`
	Agrupacion                                   string               `json:"agrupacion"`
	ColorAgrupacion                              any                  `json:"coloragrupacion"`
	ImagenAgrupacion                             any                  `json:"imagenagrupacion"`
	FechaInicio                                  string               `json:"fechainicio"`
	FechaFin                                     string               `json:"fechafin"`
	IDVarios                                     any                  `json:"idvarios"`
	Varios                                       string               `json:"varios"`
	IDCursilloImporte                            int                  `json:"idcursilloimporte"`
	Periodicidad                                 int                  `json:"periodicidad"`
	Importe                                      float64              `json:"importe"`
	ImporteCalculado                             float64              `json:"importeCalculado"`
	IDMatricula                                  int                  `json:"idmatricula"`
	IDMatriculaImporte                           int                  `json:"idmatriculaimporte"`
	ImporteMatricula                             float64              `json:"importematricula"`
	Public                                       bool                 `json:"public"`
	SoloListaDeEspera                            bool                 `json:"soloListaDeEspera"`
	PermitirPagosPendientes                      bool                 `json:"permitirPagosPendientes"`
	Grupos                                       []Grupo              `json:"grupos"`
	TipoDeListaDeEspera                          int                  `json:"tipoDeListaDeEspera"`
	LimitePlazasListaDeEspera                    int                  `json:"limitePlazasListaDeEspera"`
	GenerarElPrimerMes                           bool                 `json:"generarElPrimerMes"`
	GenerarPrimerPeriodoDeCursilloSoloEnPreventa bool                 `json:"generarPrimerPeriodoDeCursilloSoloEnPreventa"`
	GenerarElUltimoMes                           bool                 `json:"generarElUltimoMes"`
	OpcionesDivisiones                           int                  `json:"opcionesDivisiones"`
	DuracionDivisiones                           int                  `json:"duracionDivisiones"`
	Imagen                                       any                  `json:"imagen"`
	ImagenHash                                   any                  `json:"imagenHash"`
	Suplementos                                  []any                `json:"suplementos"`
	MostrarEnWebObservaciones                    bool                 `json:"mostrarEnWebObservaciones"`
	MostrarEnWebQuienRecoge                      bool                 `json:"mostrarEnWebQuienRecoge"`
	PosicionWeb                                  any                  `json:"posicionWeb"`
	PosicionWebAgrupacion                        any                  `json:"posicionWebAgrupacion"`
	NombreTipoDeCliente                          any                  `json:"nombreTipoDeCliente"`
	GruposCursillosPendientes                    []any                `json:"gruposCursillosPendientes"`
}
type CursillosResponse struct {
	Cursillos []Cursillo
}
