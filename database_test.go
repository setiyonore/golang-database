package belajar_golang_database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnections(t *testing.T) {
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/belajar-golang-database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Db berhasil konek")
}
