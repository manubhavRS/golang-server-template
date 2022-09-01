package helper

import (
	"database/sql"
	"github.com/tejashwikalptaru/tutorial/database"
	"github.com/tejashwikalptaru/tutorial/models"
)

func CreateUser(name, email string) (string, error) {
	// language=SQL
	SQL := `INSERT INTO users(name, email) 
			VALUES ($1, $2) 
			RETURNING id;`
	var userID string
	err := database.Tutorial.Get(&userID, SQL, name, email)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func GetUser(userID string) (*models.User, error) {
	// language=SQL
	SQL := `SELECT 
    		id, name, email, created_at, archived_at 
			FROM users 
			WHERE archived_at IS NULL 
			AND id = $1`
	var user models.User
	err := database.Tutorial.Get(&user, SQL, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func UpdateUser(user *models.User) error {
	//language=SQL
	SQL := `UPDATE
                users
            SET name = coalesce(trim($1), name),
                email = coalesce(trim($2), email)
            WHERE id = $3
            AND  archived_at IS NULL`
	_, err := database.Tutorial.Exec(SQL, user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID string) error {
	//language=SQL
	SQL := `UPDATE
                users
            SET archived_at=now()
            WHERE id = $3`
	_, err := database.Tutorial.Exec(SQL, userID)
	if err != nil {
		return err
	}
	return nil
}
