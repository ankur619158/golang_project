package serverconst

var (
	// Errors
	NotFoundErr           = "Not Found Error: "
	NotFoundUserErr       = "Not Found Error: User not found for given id"
	BadRequest            = "Bad Request: "
	BadRequestParams      = "Bad Request: request parameter required and was not provided"
	BadRequestToParam     = "Bad Request: Invalid 'To' parameter"
	BadRequestFromParam   = "Bad Request: Invalid 'From' parameter"
	BadRequestEmailParam  = "Bad Request: Invalid 'Email' parameter"
	ValidationErr         = "Validation Error: "
	VErrIdNotProvided     = "Validation Error: id is required and was not provided"
	VErrZeroIdProvided    = "Validation Error: non zero id is required and a zero id was provided"
	VErrUserIdNotProvided = "Validation Error: user id is required and was not provided"
	UserCreated           = "Person Created Succesfully"
)
