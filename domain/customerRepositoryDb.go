package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, error) {
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth from customers"
		err = d.db.Select(&customers, findAllSql)

	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers where status=1"
		err = d.db.Select(&customers, findAllSql, status)

	}

	if err != nil {
		logger.Error("Error while querying customers" + err.Error())
		return nil, err
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth,status from customers where customer_id=?"
	var c Customer
	err := d.db.Get(&c, customerSql, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(db *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{db}
}
