package domain

import (
	"bankingApp/errs"
	"bankingApp/logger"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database error")
		}
	}
	return &c, nil

}

// Return the customers by their status

func (s CustomerRepositoryDb) ByStat(stat string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	if stat == " " {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err := s.client.Select(&customers, findAllSql)
		if err != nil {
			logger.Error("Error while queriying Customers table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		return customers, nil
	} else {
		statSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err := s.client.Select(&customers, statSql, stat)
		if err != nil {
			logger.Error("Error while queriying Customers table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		return customers, nil
	}
}

// Helper function for establishing the DB connectivity

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
