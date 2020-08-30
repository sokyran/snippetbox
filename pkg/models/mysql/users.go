package mysql

import (
	"avgustavgust/snippetbox/pkg/models"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// UserModel is used to access Users table from MySQL base
type UserModel struct {
	DB *sql.DB
}

// Insert to insert records to db with name, email, password
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

// Authenticate is used to verify user's credentials and to return his ID
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get is used to get user data from db using his id
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
