package errors

const (
	// Al llamar al constructor de la estructura para proveer la dependencia en el contenedor
	InvalidDiProvideError = "0002"
	// Al tratar de obtener una dependencia del contenedor dig
	DependencyNotFoundError = "0003"
	// Hacer el binding
	BindingError = "0004"
	//No se pudo generar el comando
	CommandNotFoundError = "0005"
	// Error en consulta con BD
	DbQueryError = "0006"
	// No se encontraron datos en consulta
	DataNotFoundError = "0007"
	// Error al obtener numero de solicitud
	ObtenerNumeroSolicitudError = "0008"
	// Error de cabecera
	InvalidHeaderError = "0009"
	// Error error de comunicación con BD
	DbErrorError = "0010"
	// Error de concurrencia
	ConcurrencyError = "0011"
	// Al validar el request del body de creacion
	SolicitudValidateTipo      = "0012"
	SolicitudValidateBodyField = "0013"
	// Error general al consumir grpc
	GrpcError = "0014"
	// Al validar un parametro de un método
	InvalidArgumentError = "0015"
	// Al validar una variable de entorno
	EnvironmentVariableError = "0016"
	ThirdPartyError          = "0017"
	UnmarshalError           = "0018"
	DbConfigError            = "0019"
	MarshalError             = "0020"
	ConsistencyError         = "0021"
	DuplicateConstraintError = "0022"
)
