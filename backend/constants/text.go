package constants

const (
	AppName = "dot-conf"

	// Error Messages
	GeneralError              = "Oops! Something went wrong"
	UnableToSaveData          = "We weren't able to save the data. Please try again"
	CompanyAlreadyRegistered  = "Company has already been registered"
	UserAlreadyExists         = "User already exists"
	CouldNotFetchFromDatabase = "Could not fetch from database"
	FailedUpdatingData        = "Failed updating the data"
	NoPathVariableFound       = "No path variable found"
	TextToIntConversionError  = "Text to int conversion error"
	UserNotFound              = "User not found"
	CredentialsMissing        = "Credentials are missing"
	InvalidCredentials        = "Invalid credentials"
	HeaderMissing             = "Header is missing"
	AppNotFound               = "App not found"
	AppAlreadyExists          = "App already exists"
	Forbidden                 = "Forbidden"
	ConfigAlreadyExists       = "Config already exists"
	InvalidUpdateRequest      = "Invalid update request"

	// Success Messages
	CompanyCreated          = "Company has been created successfully"
	ValueFetched            = "Values are fetched for the request"
	UserCreated             = "User has been created successfully"
	DeactivatedSuccessfully = "Deactivated successfully"
	LoggedInSuccess         = "Logged in successfully"
	DeletedSuccessfully     = "Deleted successfully"

	// Dev Messages
	Created        = "Created"
	Fetched        = "Fetched"
	Updated        = "Updated"
	Deactivated    = "Deactivated"
	LoggedIn       = "Logged in"
	Deleted        = "Deleted"
	AlreadyExists  = "Already exists"
	InvalidRequest = "Invalid request"

	// Response Keys
	Company   = "company"
	Companies = "companies"
	Admin     = "admin"
	User      = "user"
	Token     = "token"

	// Default Error Responses
	MarshallErrorResponse    = "{\"statusCode\": 500, \"message\": \"Internal server error\", \"devMessage\": \"Error marshalling\"}"
	NoCompanyIdFoundResponse = "{\"statusCode\": 400, \"message\": \"No company id found\", \"devMessage\": \"No company id found\"}"
	DecodingErrorResponse    = "{\"statusCode\": 500, \"message\": \"Internal server error\", \"devMessage\": \"Error unmarshalling\"}"
	TextToIntErrorResponse   = "{\"statusCode\": 500, \"message\": \"Internal server error\", \"devMessage\": \"Failed converting string to int\"}"
	HeaderIsMissing          = "{\"statusCode\": 400, \"message\": \"Header is missing\", \"devMessage\": \"Header is missing\"}"
	NoConfigIdFoundResponse  = "{\"statusCode\": 400, \"message\": \"No config id found\", \"devMessage\": \"No config id found\"}"

	// General Strings
	Empty = ""
)
