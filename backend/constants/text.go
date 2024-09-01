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

	// Success Messages
	CompanyCreated          = "Company has been created successfully"
	ValueFetched            = "Values are fetched for the request"
	UserCreated             = "User has been created successfully"
	DeactivatedSuccessfully = "Deactivated successfully"
	LoggedInSuccess         = "Logged in successfully"

	// Dev Messages
	Created     = "Created"
	Fetched     = "Fetched"
	Updated     = "Updated"
	Deactivated = "Deactivated"
	LoggedIn    = "Logged in"

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

	// General Strings
	Empty = ""
)
