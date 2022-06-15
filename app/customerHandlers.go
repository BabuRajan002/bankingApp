package app

import (
	"bankingApp/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func (ch *CustomerHandlers) getCustomersByStat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("status") == "active" {
		stat := "1"
		log.Println("Boolean set in stat" + stat)
		customers, err := ch.service.GetAllCustomersByStat(stat)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customers)
		}

	} else if r.URL.Query().Get("status") == "inactive" {
		stat := "0"
		log.Println("Boolean set in stat" + stat)
		customers, err := ch.service.GetAllCustomersByStat(stat)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customers)
		}
	} else {
		stat := " "
		log.Println("Boolean set in stat" + stat)
		customers, err := ch.service.GetAllCustomersByStat(stat)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customers)
		}

	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
