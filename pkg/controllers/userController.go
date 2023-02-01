package controllers

import (
	"database/sql"
	"encoding/json"
	"go-demoproject/pkg/models"
	"go-demoproject/pkg/repositories"
	"log"
	"net/http"
	"time"
)

type UserController struct {
	Db *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		Db: db,
	}
}

func (uc *UserController) GetUsers() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		UserRepository := repositories.NewUserRepository(uc.Db)
		users, err := UserRepository.Select()
		if err != nil {
			log.Fatal(err)
		}

		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	})
}

func (uc *UserController) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		user := new(models.User)
		json.NewDecoder(req.Body).Decode(&user)

		user.CreatedAt = time.Now()

		UserRepository := repositories.NewUserRepository(uc.Db)
		e := UserRepository.Insert(user)
		if e != nil {
			log.Fatal(e)
		}
	})
}
