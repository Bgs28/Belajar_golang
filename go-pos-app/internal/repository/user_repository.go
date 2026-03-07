package repository

import (
	"database/sql"
	"go-pos-app/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}


// menampilkan semua user
func (r *UserRepository) GetUsers() ([]model.User, error) {
	rows, err := r.DB.Query("SELECT id, name, username, role FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User

		err := rows.Scan(
			&u.ID,  
			&u.Name, 
			&u.Username, 
			&u.Role,
		)

		if err != nil{
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) CreateUsers(user model.User) error {
	query := "INSERT INTO users (name, username, password, role) VALUES (?,?,?,?)"

	_, err := r.DB.Exec(query,
	user.Name,
	user.Username,
	user.Password,
	user.Role,
	)
	
	return err
}