package db

import (
	"fmt"
	"log"
	"time"
)

//Order -
type Ticket struct {
	Ticketid     string    `json:"ticket_id"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	Customername string    `json:"customer_name"`
	ContactName  string    `json:"contact_name"`
	Created      time.Time `json:"created"`
}

type ProductServiceCategories struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
}
type ServiceCatalogLvL struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}

/*
type ServiceCatalogLvL2 struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL3 struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL4 struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL5 struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}
type ServiceCatalogLvL6 struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}
*/

type RowsAffected struct {
	RowsAffected int64
}

func (s *Server) GetProductServiceCategories(id string) ([]ProductServiceCategories, error) {
	tsql := fmt.Sprintf("SELECT * FROM ProductServiceCategories WHERE id=@p1;")
	return s.queryProductServiceCategories(tsql, id)
}

func (s *Server) GetServiceCatalogLvL1(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL1 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

func (s *Server) GetServiceCatalogLvL2(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL2 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

func (s *Server) GetServiceCatalogLvL3(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL3 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

func (s *Server) GetServiceCatalogLvL4(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL4 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

func (s *Server) GetServiceCatalogLvL5(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL5 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

func (s *Server) GetServiceCatalogLvL6(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL6 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
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

func (s *Server) queryProductServiceCategories(tsql string, args ...interface{}) ([]ProductServiceCategories, error) {

	s.getConnection()

	category := ProductServiceCategories{}
	categories := []ProductServiceCategories{}

	log.Printf("Executing SQL: %s \n", tsql)
	log.Printf("With args: %s \n", args...)

	rows, err := s.db.Query(tsql, args...)

	if err != nil {
		log.Println("failed...")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.ProductId)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *Server) queryServiceCatalogLvL(tsql string, args ...interface{}) ([]ServiceCatalogLvL, error) {

	s.getConnection()

	category := ServiceCatalogLvL{}
	categories := []ServiceCatalogLvL{}

	log.Printf("Executing SQL: %s \n", tsql)
	log.Printf("With args: %s \n", args...)

	rows, err := s.db.Query(tsql, args...)

	if err != nil {
		log.Println("failed...")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.ParentId, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
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
		err := rows.Scan(&ticket.Ticketid, &ticket.Description, &ticket.Status, &ticket.Customername, &ticket.ContactName, &ticket.Created)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}