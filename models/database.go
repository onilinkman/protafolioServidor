package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //este driver es para que pueda ejecutar con mysql
)

func init() {
	Conectar()
}

var username = "root"
var password = ""
var host = "localhost"
var port = 3306
var dbname = "Portafolio"

var db *sql.DB

//EjecutarExec ejecuta un query sin que devuelva filas de la tabla
func EjecutarExec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

//EjecutarQuery ejecuta un query y devuelve filas de la tabla
func EjecutarQuery(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rows, nil
}

func generarURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
}

//Conectar conecta con la base de datos
func Conectar() {
	if GetConnection() != nil {
		return
	}
	coneccion, err := sql.Open("mysql", generarURL())
	if err != nil {
		panic(err)
	} else {
		db = coneccion
	}

}

//Cerrar cierra la conexioncon la base de datos
func Cerrar() {
	db.Close()
}

//Ping hace ping a la base de datos
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//GetConnection devuelve la coneccion con la base de datos
func GetConnection() *sql.DB {
	return db
}
