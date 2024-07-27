package constants

const (
	// Error Messages
	GeneralError              = "Oops! Something went wrong"
	UnableToSaveData          = "We weren't able to save the data. Please try again"
	CompanyAlreadyRegistered  = "Company has already been registered"
	UserAlreadyExists         = "User already exists"
	CouldNotFetchFromDatabase = "Could not fetch from database"
	FailedUpdatingData        = "Failed updating the data"

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
)
