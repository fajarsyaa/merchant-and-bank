package repository

import (
	"database/sql"
	"fmt"
	"project/model"
)

type LoginRepo interface {
	GetCustomerByUsername(username string) (*model.CustomerModel, error)
}

type loginRepoImpl struct {
	db *sql.DB
}

func (usrRepo *loginRepoImpl) GetCustomerByUsername(username string) (*model.CustomerModel, error) {
	qry := "SELECT username, password FROM ms_customer WHERE username = $1"
	usr := &model.CustomerModel{}
	err := usrRepo.db.QueryRow(qry, username).Scan(&usr.Username, &usr.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on loginRepoImpl.GetCustomerByUsername() : %w", err)
	}
	return usr, nil
}

func NewLoginRepo(db *sql.DB) LoginRepo {
	return &loginRepoImpl{
		db: db,
	}
}
