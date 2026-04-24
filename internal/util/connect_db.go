package util

import "fmt"

func BuildPostgreSQLDSN(host string, port int, user string, password string, dbName string, sslMode string) string {
	// postgres://admin:nnphan2126@localhost:5432/Simple_Bank?sslmode=disable
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbName,
		sslMode,
	)

}
