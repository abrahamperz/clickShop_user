package db

import (
	"clickshopUser/models"
	"clickshopUser/secretm"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Succesfull database conection ")
	return nil
}

func ConnStr(clave models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = clave.Username
	dbEndpoint = clave.Host
	dbName = "clickshopp"
	authToken = clave.Password
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true&parseTime=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn

}