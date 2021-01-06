package repositories

import (
	"database/sql"
	"log"
	"time"
	"user-crud/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) models.IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Get() ([]models.User, error) {

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	result := []models.User{}

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID,
			&user.Username,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}

func (u *UserRepository) Find(id int64) (models.User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}
	defer rows.Close()

	user := models.User{}
	var hasResult bool = false
	for rows.Next() {
		hasResult = true
		err := rows.Scan(&user.ID,
			&user.Username,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			log.Println(err)
			return models.User{}, err
		}
	}
	if !hasResult {
		return models.User{}, nil
	}

	return user, nil
}

func (u *UserRepository) Create(user *models.User) (int64, error) {
	stmt, err := u.db.Prepare("INSERT INTO users(username,password,first_name,last_name,avatar,created_at) VALUES($1,$2,$3,$4,$5,$6)")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	result, err := stmt.Exec(&user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Avatar, time.Now())
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func (u *UserRepository) Update(id int64, user *models.User) error {
	stmt, err := u.db.Prepare("UPDATE users SET username=$1, password=$2, avatar=$3, first_name=$4, last_name=$5 WHERE id=$6")
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(&user.Username, &user.Password, &user.Avatar, &user.FirstName, &user.LastName, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *UserRepository) Delete(id int64) error {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (u *UserRepository) FindUserByUsername(username string) (models.User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}
	defer rows.Close()

	user := models.User{}
	var hasResult bool = false
	for rows.Next() {
		hasResult = true
		err := rows.Scan(&user.ID,
			&user.Username,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			log.Println(err)
			return models.User{}, err
		}
	}
	if !hasResult {
		return models.User{}, nil
	}

	return user, nil

}
