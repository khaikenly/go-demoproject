package main

import (
	"go-demoproject/pkg/drivers"
	"go-demoproject/pkg/routers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	mysql_driver := &drivers.MysqlDB{}
	db := mysql_driver.Connect()

	// postgres_driver := &drivers.PostgresDB{}
	// db := postgres_driver.Connect()

	if err := db.SQL.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully created connection to database")

	defer db.SQL.Close()

	// UserRepository := repositories.NewUserRepository(db.SQL)
	// users, err := UserRepository.Select()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(users)

	// u := &models.User{
	// 	Username:  "khai123",
	// 	Password:  "12312312",
	// 	CreatedAt: time.Now(),
	// }

	// e := UserRepository.Insert(u)
	// if e != nil {
	// 	log.Fatal(e)
	// }

	routes := routers.MyRouter{}
	r := routes.Routes(mux.NewRouter())
	log.Fatal(http.ListenAndServe(":8082", r))
}
