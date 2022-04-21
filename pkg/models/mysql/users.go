package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/hideaki10/web-application-tool/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name,email,hashed_password,created) VALUES(?,?,?,UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, hashPassword)

	if err != nil {

		var mySqlError *mysql.MySQLError

		if errors.As(err, &mySqlError) {
			if mySqlError.Number == 1062 && strings.Contains(mySqlError.Message, "user_uc_email") {
				return models.ErrDuplidateEmail
			}
		}
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
