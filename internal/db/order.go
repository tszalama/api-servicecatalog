package db

import (
	"fmt"
	"log"
	"time"
)

//Order -
type Ticket struct {
	Ticketid    string `json:"ticket_id"`
	Description string `json:"description"`
	Status		string `json:"status"`
	Customername string `json:"customer_name"`
	ContactName string `json:"contact_name"`
	Created     time.Time `json:"created"`
}

type RowsAffected struct {
	RowsAffected int64
}

func (s *Server) GetTicket(ticket_id string) ([]Ticket, error) {
	tsql := fmt.Sprintf("SELECT * FROM Tickets WHERE ticket_id=@p1;")
	return s.query(tsql, ticket_id)
}

func (s *Server) GetTickets() ([]Ticket, error) {
	tsql := fmt.Sprintf("SELECT * FROM Tickets;")
	return s.query(tsql, nil)
}

func (s *Server) AddTicket(ticket_id string, description string, status string, customer_name string, contact_name string) ([]Ticket, error) {
	tsql := fmt.Sprintf("INSERT INTO Tickets(ticket_id, description, status, customer_name, contact_name) VALUES(@p1,@p2,@p3,@p4,@p5);")
	_, err := s.exec(tsql, ticket_id, description, status, customer_name, contact_name)
	if err != nil {
		return nil, err
	}

	tsql = fmt.Sprintf("SELECT * FROM Tickets WHERE ticket_id=@p1;")
	return s.query(tsql, ticket_id, description, status, customer_name, contact_name)
}

func (s *Server) EditTicket(ticket_id string, description string, status string, customer_name string, contact_name string) (RowsAffected, error) {
	tsql := fmt.Sprintf("UPDATE Tickets SET description=@p2 WHERE ticket_id=@p1")
	return s.exec(tsql, ticket_id, description)
}

func (s *Server) DeleteTicket(ticket_id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM Tickets WHERE ticket_id=@p1")
	return s.exec(tsql, ticket_id)
}

func (s *Server) exec(tsql string, args ...interface{}) (RowsAffected, error) {

	s.getConnection()

	rowsAffectedResult := RowsAffected{}
	rowsAffectedResult.RowsAffected = 0

	log.Printf("Executing SQL: %s \n", tsql)
	log.Printf("With args: %s \n", args...)

	result, err := s.db.Exec(tsql, args...)
	if err != nil {
		return rowsAffectedResult, err
	}
	num, _ := result.RowsAffected()

	rowsAffectedResult.RowsAffected = num

	return rowsAffectedResult, nil

}

func (s *Server) query(tsql string, args ...interface{}) ([]Ticket, error) {

	s.getConnection()

	ticket := Ticket{}
	tickets := []Ticket{}

	log.Printf("Executing SQL: %s \n", tsql)
	log.Printf("With args: %s \n", args...)

	rows, err := s.db.Query(tsql, args...)

	if err != nil {
		log.Println("failed...")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ticket.Ticketid, &ticket.Description, &ticket.Created)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}
