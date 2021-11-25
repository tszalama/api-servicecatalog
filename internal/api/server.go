package api

import (
	"encoding/json"
	"net/http"
	"strings"

	//"github.com/SAP-samples/kyma-runtime-extension-samples/api-mssql-go/internal/db"
	"github.com/tz19003/KymaTickets/tree/master/internal/db"
)

type ticketData struct {
	Ticketid     string `json:"ticket_id"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	Customername string `json:"customer_name"`
	ContactName  string `json:"contact_name"`
}

/*
type ProductServiceCategories struct {
	Id    string `json:"id"`
	ProductId string `json:"product_id"`
}
type ServiceCatalogLvL1 struct {
	Id    string `json:"id"`
	ParentId string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL2 struct {
	Id    string `json:"id"`
	ParentId string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL3 struct {
	Id    string `json:"id"`
	ParentId string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL4 struct {
	Id    string `json:"id"`
	ParentId string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL5 struct {
	Id    string `json:"id"`
	ParentId string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL6 struct {
	Id    string `json:"id"`
	ParentId string `json:"parent_id"`
	Description string `json:"description"`
}
*/

type server struct {
	db *db.Server
}

func InitAPIServer() *server {
	server := &server{}
	server.db = db.InitDatabase()
	return server
}

func (s *server) GetProductServiceCategories(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetProductServiceCategories(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) GetServiceCatalogLvL1(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetServiceCatalogLvL1(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) GetServiceCatalogLvL2(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetServiceCatalogLvL2(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func (s *server) GetServiceCatalogLvL3(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetServiceCatalogLvL3(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func (s *server) GetServiceCatalogLvL4(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetServiceCatalogLvL5(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func (s *server) GetServiceCatalogLvL5(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetServiceCatalogLvL5(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func (s *server) GetServiceCatalogLvL6(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetServiceCatalogLvL6(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//old
func (s *server) GetTicket(w http.ResponseWriter, r *http.Request) {

	ticket_id := strings.Split(r.URL.Path, "/")[2]
	tickets, err := s.db.GetTicket(ticket_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(tickets)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) GetTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := s.db.GetTickets()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(tickets)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) EditTicket(w http.ResponseWriter, r *http.Request) {

	var ticket ticketData

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&ticket)

	rowsEffected, err := s.db.EditTicket(ticket.Ticketid, ticket.Description, ticket.Status, ticket.Customername, ticket.ContactName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(rowsEffected)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) AddTicket(w http.ResponseWriter, r *http.Request) {

	var ticket ticketData

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&ticket)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tickets, err := s.db.AddTicket(ticket.Ticketid, ticket.Description, ticket.Status, ticket.Customername, ticket.ContactName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(tickets)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	ticket_id := strings.Split(r.URL.Path, "/")[2]
	rowsEffected, err := s.db.DeleteTicket(ticket_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(rowsEffected)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
