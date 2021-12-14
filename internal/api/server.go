package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tz19003/KymaTickets/tree/master/internal/config"
	"github.com/tz19003/KymaTickets/tree/master/internal/db"
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

func (s *server) AuthUser(w http.ResponseWriter, r *http.Request) {

	config := config.GetConfig()

	adminKeyConfig := config.AdminKey
	userKeyConfig := config.UserKey

	apiKeyAdmin := r.Header.Get("apiKeyAdmin")
	apiKeyUser := r.Header.Get("apiKeyUser")

	if apiKeyAdmin != "" {

		var signingKey = []byte(adminKeyConfig)

		log.Printf(apiKeyAdmin)
		log.Printf(string(signingKey))

		isAdmin := "true"

		if apiKeyAdmin == string(signingKey) {
			token, err := GenerateJWT(signingKey, isAdmin)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			js, _ := json.Marshal(token)

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

		} else {

			http.Error(w, "auth failed", http.StatusUnauthorized)
		}

	} else if apiKeyUser != "" {

		var signingKey = []byte(userKeyConfig)

		log.Printf(apiKeyUser)
		log.Printf(string(signingKey))

		isAdmin := "false"

		if apiKeyUser == string(signingKey) {
			token, err := GenerateJWT(signingKey, isAdmin)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			js, _ := json.Marshal(token)

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

		} else {

			http.Error(w, "auth failed", http.StatusUnauthorized)
		}
	} else {

		http.Error(w, "auth failed", http.StatusUnauthorized)

	}
}

func GenerateJWT(key []byte, isAdmin string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Minute * 25).Unix()

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *server) GetTicketCategories(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]
	categories, err := s.db.GetTicketCategories(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
	categories, err := s.db.GetServiceCatalogLvL4(id)

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

func (s *server) AddProductServiceCategories(w http.ResponseWriter, r *http.Request) {

	var category ProductServiceCategories

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddProductServiceCategories(category.Id, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

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

func (s *server) AddServiceCatalogLvL1(w http.ResponseWriter, r *http.Request) {

	var category ServiceCatalogLvL

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddServiceCatalogLvL1(category.ParentId, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) AddServiceCatalogLvL2(w http.ResponseWriter, r *http.Request) {

	var category ServiceCatalogLvL

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddServiceCatalogLvL2(category.ParentId, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) AddServiceCatalogLvL3(w http.ResponseWriter, r *http.Request) {

	var category ServiceCatalogLvL

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddServiceCatalogLvL3(category.ParentId, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) AddServiceCatalogLvL4(w http.ResponseWriter, r *http.Request) {

	var category ServiceCatalogLvL

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddServiceCatalogLvL4(category.ParentId, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) AddServiceCatalogLvL5(w http.ResponseWriter, r *http.Request) {

	var category ServiceCatalogLvL

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddServiceCatalogLvL5(category.ParentId, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) AddServiceCatalogLvL6(w http.ResponseWriter, r *http.Request) {

	var category ServiceCatalogLvL

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := s.db.AddServiceCatalogLvL6(category.ParentId, category.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *server) DeleteProductServiceCategories(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path

	id := strings.Split(url, "/")[2]

	var rowsEffected db.RowsAffected
	var err error

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

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, _ := json.Marshal(rowsEffected)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
