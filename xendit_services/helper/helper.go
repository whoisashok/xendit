package helper

var (
	HeaderKeyAuthorization = "Authorization"
	HeaderKeyPlayerID      = "X-Player"
)

var (
	GenerateCodeExp         = "[A-Z0-9]{10}"
	GenerateTokenExp        = "[0-9]{6}"
	DisbursementPrefix      = "DIS-"
	BatchReferencePrefix    = "BAT-"
	BatchDisbursementPrefix = "BDIS-"
)

var (
	UserTypeNameIndividual = "Individual"
	UserTypeNameCorporate  = "Corporate"
)

var (
	UserStatusActive = "Active"
)

var (
	DateTimeFormat       = "20060102150405"
	DateTimeMilSecFormat = "20060102150405123"
	DateTimeMacSecFormat = "20060102150405123456"
	DateTimeNanSecFormat = "20060102150405123456789"
)
