package domain

import (
	"bankingApp/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBith, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database error")
		}
	}
	return &c, nil

}

// Return the customers by their status

func (s CustomerRepositoryDb) ByStat(stat string) ([]Customer, *errs.AppError) {
	if stat == " " {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err := s.client.Query(findAllSql)
		if err != nil {
			log.Println("Error while queriying Customers table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		customers := make([]Customer, 0)
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBith, &c.Status)
			if err != nil {
				log.Println("Error while scanning customers table " + err.Error())
				return nil, errs.NewUnexpectedTableError("Unexpected Table error")
			}
			customers = append(customers, c)
			//fmt.Println(customers)
		}
		return customers, nil

	} else {
		statSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"

		rows, err := s.client.Query(statSql, stat)
		if err != nil {
			log.Println("Error while queriying Customers table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers := make([]Customer, 0)
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBith, &c.Status)
			if err != nil {
				log.Println("Error while scanning customers table " + err.Error())
				return nil, errs.NewUnexpectedTableError("Unexpected Table error")
			}
			customers = append(customers, c)
		}
		return customers, nil
	}
}

// Helper function for establishing the DB connectivity

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
