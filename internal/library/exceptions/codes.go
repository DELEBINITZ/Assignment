package exceptions

const (
	InvalidRequestStruct ServerErrorCode = 1001
	MarshallingError     ServerErrorCode = 1002
	NoClaimsFound        ServerErrorCode = 1004
	ParsingError         ServerErrorCode = 1005
	StrconvError         ServerErrorCode = 1006
	UnmarshallingError   ServerErrorCode = 1007

	//DB and Cache Errors
	DBSaveError           ServerErrorCode = 2001
	DBReadError           ServerErrorCode = 2002
	DBUpdateError         ServerErrorCode = 2003
	DBDeleteError         ServerErrorCode = 2004
	DBConnectionError     ServerErrorCode = 2005
	DBRecordNotFoundError ServerErrorCode = 2006
	DBDuplicateKeyError   ServerErrorCode = 2007
	DBInvalidSQLError     ServerErrorCode = 2008

	//External Errors
	ImageDownloadError         ServerErrorCode = 3001
	ImageUploadError           ServerErrorCode = 3002
	BucketCreationError        ServerErrorCode = 3003
	GcpClientCreationError     ServerErrorCode = 3004
	URLParsingError            ServerErrorCode = 3005
	HTTPCallError              ServerErrorCode = 3006
	HTTPInvalidStatusCodeError ServerErrorCode = 3007

	// Auth Error
	JWTTokenNotFound         ServerErrorCode = 4001
	JWTTokenExpired          ServerErrorCode = 4002
	JWTTokenInvalid          ServerErrorCode = 4003
	VerificationTokenExpired ServerErrorCode = 4006
	VerificationTokenInvalid ServerErrorCode = 4007
	InvalidCredentials       ServerErrorCode = 4004
	UserNotFOund             ServerErrorCode = 4005
)

func GenerateNewServerError(code ServerErrorCode, actualErr error, msg string, httpcode int) ServerError {
	return ServerError{
		Code:      code,
		ActualErr: actualErr,
		Msg:       msg,
		HttpCode:  httpcode,
	}
}
