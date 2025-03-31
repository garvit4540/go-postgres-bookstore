package trace

import "fmt"

func PushTrace(msg string, mp map[string]interface{}) {
	fmt.Println(msg, mp)
}

// Info
const (
	AppStarted                      = "APP_STARTED"
	SuccessfullyConnectedToDatabase = "SUCCESSFULLY_CONNECTED_TO_DATABASE"
	CreateBookRequestReceived       = "CREATE_BOOK_REQUEST_RECEIVED"
	BookSuccessFullyCreated         = "BOOK_SUCCESSFULLY_CREATED"
	FetchAllBooksRequestReceived    = "FETCH_ALL_BOOKS_REQUEST_RECEIVED"
	BooksFetchedSuccessfully        = "BOOKS_FETCHED_SUCCESSFULLY"
	FetchBookRequestReceived        = "FETCH_BOOK_REQUEST_RECEIVED"
	FetchBookSuccessfull            = "FETCH_BOOK_SUCCESSFULL"
	UpdateBookRequestReceived       = "UPDATE_BOOK_REQUEST_RECEIVED"
	BookUpdatedSuccessfully         = "BOOK_UPDATED_SUCCESSFULLY"
	DeleteBookRequestReceived       = "DELETE_BOOK_REQUEST_RECEIVED"
	BookSuccessFullyDeleted         = "BOOK_SUCCESSFULLY_DELETED"
)

// Errors
const (
	ErrorMigratingDatabases        = "ERROR_MIGRATING_DATABASES"
	ErrorReadingBody               = "ERROR_READING_BODY"
	ErrorJsonMarshalUnMarshall     = "ERROR_JSON_MARSHALL_UN_MARSHALL"
	ErrorFailedToConnectToPostgres = "ERROR_FAILED_TO_CONNECT_TO_POSTGRES"
	ErrorStartingApp               = "ERROR_STARTING_APP"
)
