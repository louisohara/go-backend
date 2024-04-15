// this is where the database operations / SQL queries go

package repositories

import (
	"database/sql"

	"example.com/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}


func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := rdb.Query("SELECT * FROM users")

	if err != nil {
		return nil, fmt.Errorf("GetUsers %q: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID,&u.FirstName, &u.Surname, &u.Email, &u.DOB); err != nil {
            return nil, fmt.Errorf("GetUsers %q: %v", err)
        }
        users = append(users, u)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("GetUsers %q: %v", err)
    }
    return users, nil
}


func (r *UserRepository) userByID(id int64) (models.User, error) {

	var u models.User
	
	row := db.QueryRow("SELECT * from users where ID = ?", id)

	if err := row.Scan(&u.ID,&u.FirstName, &u.Surname, &u.Email, &u.DOB); err != nil{
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("usersById %d: no such user", id)
		}
		return u, fmt.Errorf("usersById %d: %v", id, err)
	}
	return u, nil
}

func (r *UserRepository) addUser(u models.User) (int64, error) {
	result, err := db.Exec("INSERT INTO users (first_name, surname, email, dob) VALUES (?, ?, ?, ?)",u.FirstName, u.Surname, u.Email, u.DOB)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}

	return id, nil
}