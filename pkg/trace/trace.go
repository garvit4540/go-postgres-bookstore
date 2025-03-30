package trace

import "fmt"

func PushTrace(msg string, mp map[string]interface{}) {
	fmt.Println(msg, mp)
}

// Info
const (
	AppStarted                      = "APP_STARTED"
	SuccessfullyConnectedToDatabase = "SUCCESSFULLY_CONNECTED_TO_DATABASE"
)

// Errors
const (
	ErrorMigratingDatabases     = "ERROR_MIGRATING_DATABASES"
	ErrorReadingBody            = "ERROR_READING_BODY"
	ErrorUnMarshallingBody      = "ERROR_UN_MARSHALL_JSON"
	ErrorFailedToConnectToMySql = "ERROR_FAILED_TO_CONNECT_TO_MYSQL"
)
