package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SAP-samples/kyma-runtime-extension-samples/api-mssql-go/internal/db"
)

type ticketData struct {
	Ticketid    string `json:"ticket_id"`
	Description string `json:"description"`
	Status		string `json:"status"`
	Customername string `json:"customer_name"`
	ContactName string `json:"contact_name"`
}

type server struct {
	db *db.Server
}

func InitAPIServer() *server {
	server := &server{}
	server.db = db.InitDatabase()
	return server
}

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
