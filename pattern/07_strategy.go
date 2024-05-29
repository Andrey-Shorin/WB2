package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type DBconnecter interface {
	Connect()
}

type PostgresConnection struct {
	ConnectionString string
}

func (p *PostgresConnection) Connect() {

}

type MySQLConnection struct {
	ConnectionString string
}

func (m *MySQLConnection) Connect() {

}

type SQLLiteConnection struct {
	ConnectionString string
}

func (s *SQLLiteConnection) Connect() {

}

type DBConnection struct {
	db DBconnecter
}

func (con DBConnection) DBConnect() {
	con.db.Connect()
}
