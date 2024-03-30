package db

import (
	"clickshopUser/models"
	"clickshopUser/tools"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Register init")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentence := "INSERT INTO users(User_Email, User_UUID, USER_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println(sentence)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
