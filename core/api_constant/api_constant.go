package api_constant

import "fmt"

type HeaderList struct {
	Authorization string `json:"Authorization"`
	TraceId       string `json:"X-Trace-Id"`
	SpanId        string `json:"X-Span-Id"`
	DeviceId      string `json:"device-id"`
	DeviceOS      string `json:"device-os"`
	DeviceModel   string `json:"device-model"`
	DeviceBrand   string `json:"device-brand"`
	UserAgent     string `json:"User-Agent"`
	AppVersion    string `json:"app-version"`
	Language      string `json:"language"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Channel       string `json:"channel"`
	Forward       string `json:"forward"`
	XAuthToken    string `json:"X-Auth-Token"`
}

var Header = HeaderList{
	Authorization: "Authorization",
	TraceId:       "X-Trace-Id",
	SpanId:        "X-Span-Id",
	DeviceId:      "device-id",
	DeviceOS:      "device-os",
	DeviceModel:   "device-model",
	DeviceBrand:   "device-brand",
	UserAgent:     "User-Agent",
	AppVersion:    "app-version",
	Language:      "language",
	Latitude:      "latitude",
	Longitude:     "longitude",
	Channel:       "channel",
	Forward:       "X-Forwarded-For",
	XAuthToken:    "X-Auth-Token",
}

var (
	LANG_INDONESIA         string = "id"
	LANG_ENGLISH           string = "en"
	LANG_ARABIC            string = "ar"
	LANG_ALBANIAN          string = "sq"
	LANG_BELARUSIAN        string = "by"
	LANG_BULGARIAN         string = "bg"
	LANG_CATALAN           string = "ca"
	LANG_CHINESE           string = "za"
	LANG_CROATIAN          string = "hr"
	LANG_CZECH             string = "cs"
	LANG_DANISH            string = "da"
	LANG_DUTCH             string = "nl"
	LANG_ESTONIAN          string = "et"
	LANG_FINNISH           string = "fi"
	LANG_FRENCH            string = "fr"
	LANG_GERMAN            string = "de"
	LANG_GREEK             string = "el"
	LANG_HEBREW            string = "iw"
	LANG_HINDI             string = "hi"
	LANG_HUNGARIAN         string = "hu"
	LANG_IRISH             string = "ga"
	LANG_ITALIAN           string = "it"
	LANG_ICELANDIC         string = "is"
	LANG_JAPANESE          string = "ja"
	LANG_KOREAN            string = "ko"
	LANG_LATVIAN           string = "lv"
	LANG_LITHUANIAN        string = "lt"
	LANG_MACEDONIAN        string = "mk"
	LANG_MALAY             string = "ms"
	LANG_MALTESE           string = "mt"
	LANG_NORWEGIAN         string = "no"
	LANG_NORWEGIAN_BOKMAL  string = "nb"
	LANG_NORWEGIAN_NYNORSK string = "nn"
	LANG_POLISH            string = "pl"
	LANG_PORTUGUESE        string = "pt"
	LANG_ROMANIAN          string = "ro"
	LANG_RUSSIAN           string = "ru"
	LANG_SERBIAN           string = "sr"
	LANG_SLOVAK            string = "sk"
	LANG_SLOVENIAN         string = "sl"
	LANG_SPANISH           string = "es"
	LANG_SWEDISH           string = "sv"
	LANG_THAI              string = "th"
	LANG_TURKISH           string = "tr"
	LANG_UKRAINIAN         string = "uk"
	LANG_VIETNAMESE        string = "vi"
)

type ErrorCode struct {
	code        string
	description string
}

// NewErrorCode is a factory function to create a new ErrorCode instance.
func NewErrorCode(code, description string) *ErrorCode {
	return &ErrorCode{
		code:        code,
		description: description,
	}
}

// GetCode returns the code of the ErrorCode.
func (e *ErrorCode) GetCode() string {
	return e.code
}

// GetDescription returns the description of the ErrorCode.
func (e *ErrorCode) GetDescription() string {
	return e.description
}

// String returns a readable format of the ErrorCode.
func (e *ErrorCode) String() string {
	return fmt.Sprintf("ErrorCode{code='%s', description='%s'}", e.code, e.description)
}

// GeneralErrorCache General Error Cache General Error Cache
var (
	GeneralErrorCache = NewErrorCode("CacheErrorCode", "CacheErrorCode")
)

// General Error Codes
var (
	GeneralCodeSystem = NewErrorCode("IST-MVC", "IST-MVC System")
	GeneralSuccess    = NewErrorCode("00", "Success")
	GeneralError      = NewErrorCode("0019", "General error occurred")
)

// Client Request Error Codes
var (
	InvalidRequest      = NewErrorCode("4000", "Invalid request parameters")
	NotFoundRequest     = NewErrorCode("4040", "Requested data not found")
	UnauthorizedRequest = NewErrorCode("4010", "Unauthorized access")
	ForbiddenRequest    = NewErrorCode("4030", "Access forbidden")
	InvalidParameter    = NewErrorCode("4002", "Invalid parameter value")
	InvalidFormat       = NewErrorCode("4003", "Invalid format")
)

// Server Error Codes
var (
	ServerInternalError  = NewErrorCode("5000", "Internal server error")
	ServiceUnavailable   = NewErrorCode("5030", "Server unavailable error")
	ServerGatewayTimeout = NewErrorCode("5040", "Internal server error")
)

// Database / Data Errors: 6000 – 6999
var (
	DataError            = NewErrorCode("6000", "General Data Error")
	DataParameterInvalid = NewErrorCode("6010", "Invalid Data Parameter")
	DataDuplicate        = NewErrorCode("6020", "Data Error due to duplicate")
	DataNotFound         = NewErrorCode("6040", "Data not found / no records")
	DataEmpty            = NewErrorCode("6041", "Data is empty")
	DataConflict         = NewErrorCode("6099", "Data Conflict")
)

// Communication Errors: 8000 – 8999
var (
	CommunicationError = NewErrorCode("8000", "Communication / Network error")
	ConnectionError    = NewErrorCode("8001", "Connection Read")
	ConnectionRefused  = NewErrorCode("8002", "Connection refused")
	ConnectTimeout     = NewErrorCode("8003", "Connect Timeout occurred")
	ReadTimeout        = NewErrorCode("8004", "Read Timeout occurred")
	APIError           = NewErrorCode("8010", "API call error")
)
