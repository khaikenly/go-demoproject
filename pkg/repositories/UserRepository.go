package repositories

import (
	"database/sql"
	"go-demoproject/pkg/models"
	"log"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) Select() ([]models.User, error) {
	rows, err := u.Db.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		user := models.User{}

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func (u *UserRepository) Insert(user *models.User) error {
	query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
	result, err := u.Db.Exec(query, user.Username, user.Password, user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	uID, _ := result.LastInsertId()
	user.Id = int(uID)

	log.Print("Record added: ", user)
	return nil
}
