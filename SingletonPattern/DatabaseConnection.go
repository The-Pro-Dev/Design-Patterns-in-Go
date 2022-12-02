package SingletonPattern

type db_ConnectionState string

const (
	db_INITIALISED = db_ConnectionState("INITIALISED")
	db_OPEN        = db_ConnectionState("OPEN")
	db_CLOSED      = db_ConnectionState("CLOSED")
)

type DatabaseConnection struct {
	state db_ConnectionState
	url   string
}

func (dbConnection *DatabaseConnection) close() {
	dbConnection.state = db_CLOSED
}

func (dbConnection *DatabaseConnection) connect() {
	dbConnection.state = db_OPEN
}

func (dbConnection *DatabaseConnection) getState() db_ConnectionState {
	return dbConnection.state
}

func (dbConnection *DatabaseConnection) getUrl() string {
	return dbConnection.url
}

var dbConnection *DatabaseConnection = nil

func createDatabaseConnection(url string) DatabaseConnection {
	if dbConnection == nil {
		dbConnection = &DatabaseConnection{state: db_INITIALISED, url: url}
	}

	return *dbConnection
}
