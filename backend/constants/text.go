package constants

const (
	// Error Messages
	GeneralError              = "Oops! Something went wrong"
	UnableToSaveData          = "We weren't able to save the data. Please try again"
	CompanyAlreadyRegistered  = "Company has already been registered"
	UserAlreadyExists         = "User already exists"
	CouldNotFetchFromDatabase = "Could not fetch from database"
	FailedUpdatingData        = "Failed updating the data"
	NoCompanyIdFound          = "No company id found"
	TextToIntConversionError  = "Text to int conversion error"

	// Success Messages
	CompanyCreated = "Company has been stored successfully, will send a verification mail in 24 hours"
	ValueFetched   = "Values are fetched for the request"

	// Dev Messages
	Created = "Created"
	Fetched = "Fetched"
	Updated = "Updated"

	// Response Keys
	Company   = "company"
	Companies = "companies"
	Admin     = "admin"

	// Default Error Responses
	MarshallErrorResponse    = "{\"statusCode\": 500, \"message\": \"Internal server error\", \"devMessage\": \"Error marshalling\"}"
	NoCompanyIdFoundResponse = "{\"statusCode\": 400, \"message\": \"No company id found\", , \"devMessage\": \"No company id found\"}"
	DecodingErrorResponse    = "{\"statusCode\": 500, \"message\": \"Internal server error\", \"devMessage\": \"Error unmarshalling\"}"
	TextToIntErrorResponse   = "{\"statusCode\": 500, \"message\": \"Internal server error\", \"devMessage\": \"Failed converting string to int\"}"

	// General Strings
	Empty = ""
)
