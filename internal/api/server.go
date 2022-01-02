package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tszalama/api-servicecatalog/tree/main/internal/config"
	"github.com/tszalama/api-servicecatalog/tree/main/internal/db"
)

//Ticket category structure (matches DB table structure)
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

//Product structure (matches DB table structure)
type ProductServiceCategories struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

//Service category level structure (matches DB table structure)
type ServiceCatalogLvL struct {
	Id          string `json:"id"`
	ParentId    string `json:"parent_id"`
	Description string `json:"description"`
}

type server struct {
	db *db.Server
}

//Initialize server
func InitAPIServer() *server {
	server := &server{}
	server.db = db.InitDatabase()
	return server
}

//Function checks provided creditentials and returns a JWT token that can be used for further requests
func (s *server) AuthUser(w http.ResponseWriter, r *http.Request) {

	config := config.GetConfig()

	//Get admin and regular user keys from Enviroment variables
	adminKeyConfig := config.AdminKey
	userKeyConfig := config.UserKey

	//Get provided api key from request header
	apiKeyAdmin := r.Header.Get("apiKeyAdmin")
	apiKeyUser := r.Header.Get("apiKeyUser")

	//If the client is attepting to log in as "Admin", check that the provided Admin key matches the Admin key defined in conig
	if apiKeyAdmin != "" {

		var signingKey = []byte(adminKeyConfig)

		log.Printf(apiKeyAdmin)
		log.Printf(string(signingKey))

		isAdmin := "true"

		if apiKeyAdmin == string(signingKey) {
			//If provided key matches, user is authorized to access the resource
			//Generete a JWT token for the Admin user
			token, err := GenerateJWT(signingKey, isAdmin)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			js, _ := json.Marshal(token)

			//Return generated JWT token
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

		} else {
			//If the provided JWT token failed validation, return 401 status "not authorized"
			http.Error(w, "auth failed", http.StatusUnauthorized)
		}

		//If the client is attepting to log in as "User", check that the provided User key matches the User key defined in conig
	} else if apiKeyUser != "" {

		var signingKey = []byte(userKeyConfig)

		log.Printf(apiKeyUser)
		log.Printf(string(signingKey))

		isAdmin := "false"

		if apiKeyUser == string(signingKey) {
			//If provided key matches, user is authorized to access the resource
			//Generete a JWT token for the Admin user
			token, err := GenerateJWT(signingKey, isAdmin)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			js, _ := json.Marshal(token)

			//Return generated JWT token
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

		} else {
			//If the provided JWT token failed validation, return 401 status "not authorized"
			http.Error(w, "auth failed", http.StatusUnauthorized)
		}
	} else {
		//If neither Admin nor User key was provided, return 401 status "not authorized"
		http.Error(w, "auth failed", http.StatusUnauthorized)

	}
}

//Function that generates a JWT token based on the provided siging key, admin indicator and expiry date
func GenerateJWT(key []byte, isAdmin string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = isAdmin
	//JWT token expires 25 minutes after its generated
	claims["exp"] = time.Now().Add(time.Minute * 25).Unix()

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//Function gets a specific Ticket entry from Database based on the id in the URL path and returns it to the client
func (s *server) GetTicketCategories(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetTicketCategories(id)

	//If database retrieve failed, return internal server error status
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)
	//Send data back to client
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Function gets a specific Product entry from Database based on the id in the URL path and returns it to the user
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

//Function gets a all Product entries from Database and returns them to the user
func (s *server) GetAllProductServiceCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := s.db.GetAllProductServiceCategories()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Function gets a specific Service Catalog Level N entry from Database based on the id in the URL path and returns it to the user
func (s *server) GetServiceCatalogLvL(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	url := r.URL.Path
	var categories []db.ServiceCatalogLvL
	var err error

	if strings.Contains(url, "servicecataloglvl1") {
		categories, err = s.db.GetServiceCatalogLvL1(id)
	}
	if strings.Contains(url, "servicecataloglvl2") {
		categories, err = s.db.GetServiceCatalogLvL2(id)
	}
	if strings.Contains(url, "servicecataloglvl3") {
		categories, err = s.db.GetServiceCatalogLvL3(id)
	}
	if strings.Contains(url, "servicecataloglvl4") {
		categories, err = s.db.GetServiceCatalogLvL4(id)
	}
	if strings.Contains(url, "servicecataloglvl5") {
		categories, err = s.db.GetServiceCatalogLvL5(id)
	}
	if strings.Contains(url, "servicecataloglvl6") {
		categories, err = s.db.GetServiceCatalogLvL6(id)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Function deletes a category in Service Category Level N from database and return the affected row count to client
func (s *server) DeleteProductServiceCategories(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path

	//Get target Category ID from URL params
	id := strings.Split(url, "/")[2]

	var rowsEffected db.RowsAffected
	var err error

	//Find the the correct deletion function based on the route and execute it
	if strings.Contains(url, "productservicecategories") {
		rowsEffected, err = s.db.DeleteProductServiceCategories(id)
	}
	if strings.Contains(url, "ticketcategories") {
		rowsEffected, err = s.db.DeleteTicketCategories(id)
	}
	if strings.Contains(url, "servicecataloglvl1") {
		rowsEffected, err = s.db.DeleteServiceCatalogLvL1(id)
	}
	if strings.Contains(url, "servicecataloglvl2") {
		rowsEffected, err = s.db.DeleteServiceCatalogLvL2(id)
	}
	if strings.Contains(url, "servicecataloglvl3") {
		rowsEffected, err = s.db.DeleteServiceCatalogLvL3(id)
	}
	if strings.Contains(url, "servicecataloglvl4") {
		rowsEffected, err = s.db.DeleteServiceCatalogLvL4(id)
	}
	if strings.Contains(url, "servicecataloglvl5") {
		rowsEffected, err = s.db.DeleteServiceCatalogLvL5(id)
	}
	if strings.Contains(url, "servicecataloglvl6") {
		rowsEffected, err = s.db.DeleteServiceCatalogLvL6(id)
	}

	//If data deletion failed, return status - internal server error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//If data deletion was succesful, return the affected row count
	js, _ := json.Marshal(rowsEffected)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Function saves Product data to the database
func (s *server) AddProductServiceCategories(w http.ResponseWriter, r *http.Request) {

	//Define expected message struct
	var category ProductServiceCategories

	//Decode message body into struct
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	//In case of error, return internal error status
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Call functions that saves Product data to database
	categories, err := s.db.AddProductServiceCategories(category.Id, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Return a copy of the saved data if the save was succesful
	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Function saves Ticket data to the database and returns a copy of the saved data to the client
func (s *server) AddTicketCategories(w http.ResponseWriter, r *http.Request) {

	var category TicketCategories

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddTicketCategories(category.Ticketid, category.Productid, category.CategoryIdLvl1, category.CategoryIdLvl2, category.CategoryIdLvl3, category.CategoryIdLvl4, category.CategoryIdLvl5, category.CategoryIdLvl6)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Function saves Service Catalog Level N data to the database and returns a copy of the saved data to the client
func (s *server) AddServiceCatalogLvL(w http.ResponseWriter, r *http.Request) {

	var category ServiceCatalogLvL

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := r.URL.Path
	var categories []db.ServiceCatalogLvL

	if strings.Contains(url, "servicecataloglvl1") {
		categories, err = s.db.AddServiceCatalogLvL1(category.ParentId, category.Description)
	}
	if strings.Contains(url, "servicecataloglvl2") {
		categories, err = s.db.AddServiceCatalogLvL2(category.ParentId, category.Description)
	}
	if strings.Contains(url, "servicecataloglvl3") {
		categories, err = s.db.AddServiceCatalogLvL3(category.ParentId, category.Description)
	}
	if strings.Contains(url, "servicecataloglvl4") {
		categories, err = s.db.AddServiceCatalogLvL4(category.ParentId, category.Description)
	}
	if strings.Contains(url, "servicecataloglvl5") {
		categories, err = s.db.AddServiceCatalogLvL5(category.ParentId, category.Description)
	}
	if strings.Contains(url, "servicecataloglvl6") {
		categories, err = s.db.AddServiceCatalogLvL6(category.ParentId, category.Description)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
