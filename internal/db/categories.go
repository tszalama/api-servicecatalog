package db

import (
	"fmt"
	"log"
)

//Ticket category structure (matches database table structure)
type TicketCategories struct {
	Ticketid       string `json:"ticket_id"`
	Productid      string `json:"product_id"`
	CategoryIdLvl1 string `json:"category_id_lvl1"`
	CategoryIdLvl2 string `json:"category_id_lvl2"`
	CategoryIdLvl3 string `json:"category_id_lvl3"`
	CategoryIdLvl4 string `json:"category_id_lvl4"`
	CategoryIdLvl5 string `json:"category_id_lvl5"`
	CategoryIdLvl6 string `json:"category_id_lvl6"`
}

//Product structure (matches database table structure)
type ProductServiceCategories struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

//Service Category level category structure (matches database table structure for categories level 1-6)
type ServiceCatalogLvL struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}

type RowsAffected struct {
	RowsAffected int64
}

//Function defines and executes SQL querry that returns all Ticket Categories from database
func (s *Server) GetTicketCategories(id string) ([]TicketCategories, error) {
	tsql := fmt.Sprintf("SELECT * FROM TicketCategories WHERE ticket_id=@p1;")
	return s.queryTicketCategories(tsql, id)
}

//Function defines and executes SQL querry that saves Ticket data to the database and also returns the created entry
func (s *Server) AddTicketCategories(Ticketid string, Productid string, CategoryIdLvl1 string, CategoryIdLvl2 string,
	CategoryIdLvl3 string, CategoryIdLvl4 string, CategoryIdLvl5 string, CategoryIdLvl6 string) ([]TicketCategories, error) {

	tsqlQuery := fmt.Sprintf("INSERT INTO TicketCategories(ticket_id,product_id,category_id_lvl1,category_id_lvl2,category_id_lvl3,category_id_lvl4,category_id_lvl5,category_id_lvl6) VALUES(@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8); SELECT * FROM TicketCategories WHERE ticket_id=@p1;")

	return s.queryTicketCategories(tsqlQuery, Ticketid, Productid, CategoryIdLvl1, CategoryIdLvl2,
		CategoryIdLvl3, CategoryIdLvl4, CategoryIdLvl5, CategoryIdLvl6)
}

//Function defines and executes SQL querry that deletes a specific Ticket entry from the database
func (s *Server) DeleteTicketCategories(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM TicketCategories WHERE ticket_id=@p1")
	return s.deleteEntry(tsql, id)
}

//Function defines and executes SQL querry that gets a specific Product from the database
func (s *Server) GetProductServiceCategories(id string) ([]ProductServiceCategories, error) {
	tsql := fmt.Sprintf("SELECT * FROM ProductServiceCategories WHERE id=@p1;")
	return s.queryProductServiceCategories(tsql, id)
}

//Function defines and executes SQL querry that returns all Products from database
func (s *Server) GetAllProductServiceCategories() ([]ProductServiceCategories, error) {
	tsql := fmt.Sprintf("SELECT * FROM ProductServiceCategories;")
	return s.queryAllProductCategories(tsql)
}

//Function defines and executes SQL querry that gets a specific Level 1 Service Category from the database based on the provided Parent Id
func (s *Server) GetServiceCatalogLvL1(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL1 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

//Function defines and executes SQL querry that gets a specific Level 2 Service Category from the database based on the provided Parent Id
func (s *Server) GetServiceCatalogLvL2(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL2 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

//Function defines and executes SQL querry that gets a specific Level 3 Service Category from the database based on the provided Parent Id
func (s *Server) GetServiceCatalogLvL3(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL3 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

//Function defines and executes SQL querry that gets a specific Level 4 Service Category from the database based on the provided Parent Id
func (s *Server) GetServiceCatalogLvL4(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL4 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

//Function defines and executes SQL querry that gets a specific Level 5 Service Category from the database based on the provided Parent Id
func (s *Server) GetServiceCatalogLvL5(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL5 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

//Function defines and executes SQL querry that gets a specific Level 6 Service Category from the database based on the provided Parent Id
func (s *Server) GetServiceCatalogLvL6(id string) ([]ServiceCatalogLvL, error) {
	tsql := fmt.Sprintf("SELECT * FROM ServiceCatalogLvL6 WHERE parent_id=@p1;")
	return s.queryServiceCatalogLvL(tsql, id)
}

//Function defines and executes SQL querry that saves Product data to the database and returns the created entry
func (s *Server) AddProductServiceCategories(id string, description string) ([]ProductServiceCategories, error) {

	tsqlQuery := fmt.Sprintf("INSERT INTO ProductServiceCategories(id,description) VALUES(@p1,@p2); SELECT * FROM ProductServiceCategories WHERE id = @p1;")

	return s.queryProductServiceCategories(tsqlQuery, id, description)
}

//Function defines and executes SQL querry that saves Level 1 Service Catelog data to the database and returns the created entry
func (s *Server) AddServiceCatalogLvL1(parent_id string, description string) ([]ServiceCatalogLvL, error) {
	/* NOTE
	Selection has to be done via SCOPE_IDENTITY which returns the last used idenity in the current session
	This is beacause the ID is generated automaticaly in the backend by SQL IDENTITY and is not known upon creation
	*/
	tsqlQuerry := fmt.Sprintf("INSERT INTO ServiceCatalogLvL1(parent_id, description) VALUES(@p1,@p2); SELECT * FROM ServiceCatalogLvL1 WHERE id = SCOPE_IDENTITY();")

	return s.queryServiceCatalogLvL(tsqlQuerry, parent_id, description)
}

//Function defines and executes SQL querry that saves Level 2 Service Catelog data to the database and returns the created entry
func (s *Server) AddServiceCatalogLvL2(parent_id string, description string) ([]ServiceCatalogLvL, error) {

	tsqlQuerry := fmt.Sprintf("INSERT INTO ServiceCatalogLvL2(parent_id, description) VALUES(@p1,@p2); SELECT * FROM ServiceCatalogLvL2 WHERE id = SCOPE_IDENTITY();")

	return s.queryServiceCatalogLvL(tsqlQuerry, parent_id, description)
}

//Function defines and executes SQL querry that saves Level 3 Service Catelog data to the database and returns the created entry
func (s *Server) AddServiceCatalogLvL3(parent_id string, description string) ([]ServiceCatalogLvL, error) {
	tsqlQuerry := fmt.Sprintf("INSERT INTO ServiceCatalogLvL3(parent_id, description) VALUES(@p1,@p2); SELECT * FROM ServiceCatalogLvL3 WHERE id = SCOPE_IDENTITY();")

	return s.queryServiceCatalogLvL(tsqlQuerry, parent_id, description)
}

//Function defines and executes SQL querry that saves Level 4 Service Catelog data to the database and returns the created entry
func (s *Server) AddServiceCatalogLvL4(parent_id string, description string) ([]ServiceCatalogLvL, error) {
	tsqlQuerry := fmt.Sprintf("INSERT INTO ServiceCatalogLvL4(parent_id, description) VALUES(@p1,@p2); SELECT * FROM ServiceCatalogLvL4 WHERE id = SCOPE_IDENTITY();")

	return s.queryServiceCatalogLvL(tsqlQuerry, parent_id, description)
}

//Function defines and executes SQL querry that saves Level 5 Service Catelog data to the database and returns the created entry
func (s *Server) AddServiceCatalogLvL5(parent_id string, description string) ([]ServiceCatalogLvL, error) {
	tsqlQuerry := fmt.Sprintf("INSERT INTO ServiceCatalogLvL5(parent_id, description) VALUES(@p1,@p2); SELECT * FROM ServiceCatalogLvL5 WHERE id = SCOPE_IDENTITY();")

	return s.queryServiceCatalogLvL(tsqlQuerry, parent_id, description)
}

//Function defines and executes SQL querry that saves Level 6 Service Catelog data to the database and returns the created entry
func (s *Server) AddServiceCatalogLvL6(parent_id string, description string) ([]ServiceCatalogLvL, error) {
	tsqlQuerry := fmt.Sprintf("INSERT INTO ServiceCatalogLvL6(parent_id, description) VALUES(@p1,@p2); SELECT * FROM ServiceCatalogLvL6 WHERE id = SCOPE_IDENTITY();")

	return s.queryServiceCatalogLvL(tsqlQuerry, parent_id, description)
}

//Function defines and executes SQL querry that deletes a specific Product entry from the database
func (s *Server) DeleteProductServiceCategories(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM ProductServiceCategories WHERE id=@p1")
	return s.deleteEntry(tsql, id)
}

//Function defines and executes SQL querry that deletes a specific Level 1 Service Catalog entry from the database
func (s *Server) DeleteServiceCatalogLvL1(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM ServiceCatalogLvL1 WHERE id=@p1")
	log.Printf("Delete 1")
	return s.deleteEntry(tsql, id)
}

//Function defines and executes SQL querry that deletes a specific Level 2 Service Catalog entry from the database
func (s *Server) DeleteServiceCatalogLvL2(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM ServiceCatalogLvL2 WHERE id=@p1")
	return s.deleteEntry(tsql, id)
}

//Function defines and executes SQL querry that deletes a specific Level 3 Service Catalog entry from the database
func (s *Server) DeleteServiceCatalogLvL3(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM ServiceCatalogLvL3 WHERE id=@p1")
	return s.deleteEntry(tsql, id)
}

//Function defines and executes SQL querry that deletes a specific Level 4 Service Catalog entry from the database
func (s *Server) DeleteServiceCatalogLvL4(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM ServiceCatalogLvL4 WHERE id=@p1")
	return s.deleteEntry(tsql, id)
}

//Function defines and executes SQL querry that deletes a specific Level 5 Service Catalog entry from the database
func (s *Server) DeleteServiceCatalogLvL5(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM ServiceCatalogLvL5 WHERE id=@p1")
	return s.deleteEntry(tsql, id)
}

//Function defines and executes SQL querry that deletes a specific Level 6 Service Catalog entry from the database
func (s *Server) DeleteServiceCatalogLvL6(id string) (RowsAffected, error) {
	tsql := fmt.Sprintf("DELETE FROM ServiceCatalogLvL6 WHERE id=@p1")
	return s.deleteEntry(tsql, id)
}

//Function for executing DELETE SQL querries in the database
func (s *Server) deleteEntry(tsql string, args ...interface{}) (RowsAffected, error) {

	//Set connection with database
	s.getConnection()

	log.Printf("Executing SQL: %s \n", tsql)
	log.Printf("With args: %s \n", args...)

	rowsModified := RowsAffected{}
	rowsModified.RowsAffected = 0
	//Execute querry in database using the provided querry string + inserted variables (args)
	result, err := s.db.Exec(tsql, args...)
	//If database opperation failed, return error
	if err != nil {
		return rowsModified, err
	}
	num, _ := result.RowsAffected()

	rowsModified.RowsAffected = num
	//If opperation was succesful, return the modified row count
	return rowsModified, nil

}

//Function for executing CREATE and SELECT SQL querries for the TicketCategories table
func (s *Server) queryTicketCategories(tsqlQuery string, args ...interface{}) ([]TicketCategories, error) {

	//Set database connection
	s.getConnection()
	//Define matching data object structure
	category := TicketCategories{}
	categories := []TicketCategories{}

	log.Printf("SQL Start: %s \n", tsqlQuery)
	log.Printf("Script: %s \n", args...)
	//Execute querry in database using the provided querry string + inserted variables (args)
	rows, err := s.db.Query(tsqlQuery, args...)
	//If database opperation failed, return error
	if err != nil {
		log.Println("Query failed")
		return nil, err
	}
	defer rows.Close()
	//If opperation was succesful, scan the output for newly created entry and store the data in object
	for rows.Next() {
		err := rows.Scan(&category.Ticketid, &category.Productid, &category.CategoryIdLvl1, &category.CategoryIdLvl2, &category.CategoryIdLvl3, &category.CategoryIdLvl4, &category.CategoryIdLvl5, &category.CategoryIdLvl6)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	//Return created entry data
	return categories, nil
}

//Function for executing SELECT * SQL querries specificaly for the ProductCategories table
func (s *Server) queryAllProductCategories(tsqlQuery string) ([]ProductServiceCategories, error) {

	s.getConnection()

	category := ProductServiceCategories{}
	categories := []ProductServiceCategories{}

	log.Printf("SQL Start: %s \n", tsqlQuery)

	//Similar to the function above but has no variables/ args because we are doing a SELECT * opperation a table
	rows, err := s.db.Query(tsqlQuery)

	if err != nil {
		log.Println("Query failed")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

//Function for executing CREATE and SELECT SQL querries specificaly for the ProductCategories table
func (s *Server) queryProductServiceCategories(tsqlQuery string, args ...interface{}) ([]ProductServiceCategories, error) {

	s.getConnection()

	category := ProductServiceCategories{}
	categories := []ProductServiceCategories{}

	log.Printf("SQL Start: %s \n", tsqlQuery)
	log.Printf("Script: %s \n", args...)

	rows, err := s.db.Query(tsqlQuery, args...)

	if err != nil {
		log.Println("Query failed")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

//Function for executing CREATE and SELECT SQL querries specificaly for the ServiceCatatalogLvl 1-6 tables
func (s *Server) queryServiceCatalogLvL(tsqlQuery string, args ...interface{}) ([]ServiceCatalogLvL, error) {

	s.getConnection()

	category := ServiceCatalogLvL{}
	categories := []ServiceCatalogLvL{}

	log.Printf("SQL Start: %s \n", tsqlQuery)
	log.Printf("Script: %s \n", args...)

	rows, err := s.db.Query(tsqlQuery, args...)

	if err != nil {
		log.Println("Query failed")
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
