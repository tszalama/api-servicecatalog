package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"

	"github.com/tszalama/api-servicecatalog/tree/main/internal/api"
	"github.com/tszalama/api-servicecatalog/tree/main/internal/config"
)

// All incoming requests have to go trough this midlleware function
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		config := config.GetConfig()

		adminKeyConfig := config.AdminKey
		userKeyConfig := config.UserKey
		// Do stuff here
		log.Println(r.RequestURI)

		// If path is not /auth, validate provided JWT token
		if r.RequestURI != "/auth" {

			// Parse token using ecription alg HMAC and signing key
			keyString := ""

			if r.Header["Token"] != nil {

				//First parse JWT token without verification to get the user role claim
				claims := jwt.MapClaims{}
				_, _, err := new(jwt.Parser).ParseUnverified(r.Header["Token"][0], claims)
				if err != nil {
					fmt.Println(err)
					return
				}

				userType := ""

				// Get user type from JWT token claims and determine if the token has to be checked agains the admin key or the user key
				for key, val := range claims {
					fmt.Printf("Key: %v, value: %v\n", key, val)
					if key == "admin" && val == "false" {
						keyString = userKeyConfig
						userType = "user"
					} else if key == "admin" && val == "true" {
						keyString = adminKeyConfig
						userType = "admin"
					}
				}

				signingKey := []byte(keyString)

				// If user type in JWT token is "user" return 401 (Regular users should only be able to access GET methods)
				if userType == "user" && r.Method != "GET" {
					http.Error(w, "not authorized", http.StatusUnauthorized)
					return
				}

				//Validate the provided JWT token with either the admin or user key (based on the determined claim)
				token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There is an error")
					}

					return signingKey, nil
				})

				//If token verification failed, return 401
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
				}

				// if token is valid, user is authorized to access API.
				// Proceed to original endpoint
				if token.Valid {
					next.ServeHTTP(w, r)
				}

			} else { //If there is no token provided at all, return 401
				http.Error(w, "not authorized", http.StatusUnauthorized)
			}

		} else { //If request path is /auth, verify login credentials and return JWT token that can be used for further requests
			next.ServeHTTP(w, r)
			log.Printf("auth required")
		}
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	apiServer := api.InitAPIServer()

	// Route all incoming requests trough the middleware function
	router.Use(loggingMiddleware)

	/*
		---ROUTER DEFINITION---

		Defines the allowed routes, methods and which handler to call for each request
	*/
	router.HandleFunc("/auth", apiServer.AuthUser).Methods("GET")

	router.HandleFunc("/productservicecategories/{id}", apiServer.GetProductServiceCategories).Methods("GET")
	router.HandleFunc("/productservicecategories", apiServer.GetAllProductServiceCategories).Methods("GET")
	router.HandleFunc("/ticketcategories/{id}", apiServer.GetTicketCategories).Methods("GET")
	router.HandleFunc("/servicecataloglvl1/{id}", apiServer.GetServiceCatalogLvL1).Methods("GET")
	router.HandleFunc("/servicecataloglvl2/{id}", apiServer.GetServiceCatalogLvL2).Methods("GET")
	router.HandleFunc("/servicecataloglvl3/{id}", apiServer.GetServiceCatalogLvL3).Methods("GET")
	router.HandleFunc("/servicecataloglvl4/{id}", apiServer.GetServiceCatalogLvL4).Methods("GET")
	router.HandleFunc("/servicecataloglvl5/{id}", apiServer.GetServiceCatalogLvL5).Methods("GET")
	router.HandleFunc("/servicecataloglvl6/{id}", apiServer.GetServiceCatalogLvL6).Methods("GET")

	router.HandleFunc("/productservicecategories", apiServer.AddProductServiceCategories).Methods("POST")
	router.HandleFunc("/ticketcategories", apiServer.AddTicketCategories).Methods("POST")
	router.HandleFunc("/servicecataloglvl1", apiServer.AddServiceCatalogLvL1).Methods("POST")
	router.HandleFunc("/servicecataloglvl2", apiServer.AddServiceCatalogLvL2).Methods("POST")
	router.HandleFunc("/servicecataloglvl3", apiServer.AddServiceCatalogLvL3).Methods("POST")
	router.HandleFunc("/servicecataloglvl4", apiServer.AddServiceCatalogLvL4).Methods("POST")
	router.HandleFunc("/servicecataloglvl5", apiServer.AddServiceCatalogLvL5).Methods("POST")
	router.HandleFunc("/servicecataloglvl6", apiServer.AddServiceCatalogLvL6).Methods("POST")

	router.HandleFunc("/productservicecategories/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/ticketcategories/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl1/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl2/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl3/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl4/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl5/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl6/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
