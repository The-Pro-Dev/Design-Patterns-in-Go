package SingletonPattern

import "fmt"

func displayConnection(connection DatabaseConnection) {
	fmt.Println("Database URL:", connection.getUrl(), "Connection Status:", connection.getState())
}

func Main() {
	fmt.Println("** Singleton Pattern **")

	url := "mongodb://localhost:27017/myproject"
	connection := createDatabaseConnection(url)
	displayConnection(connection)

	connection.connect()
	displayConnection(connection)

	connection.close()
	displayConnection(connection)
}
