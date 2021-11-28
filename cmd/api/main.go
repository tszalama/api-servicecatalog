package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	//"github.com/SAP-samples/kyma-runtime-extension-samples/api-mssql-go/internal/api"
	"github.com/tz19003/KymaTickets/tree/master/internal/api"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	apiServer := api.InitAPIServer()

	/*
			router.HandleFunc("/orders", apiServer.GetOrders).Methods("GET")
			router.HandleFunc("/orders/{id}", apiServer.GetOrder).Methods("GET")
			router.HandleFunc("/orders/{id}", apiServer.DeleteOrder).Methods("DELETE")
			router.HandleFunc("/orders/{id}", apiServer.EditOrder).Methods("PUT")
			router.HandleFunc("/orders", apiServer.AddOrder).Methods("POST")

		router.HandleFunc("/tickets", apiServer.GetTickets).Methods("GET")
		router.HandleFunc("/tickets/{id}", apiServer.GetTicket).Methods("GET")
		router.HandleFunc("/tickets/{id}", apiServer.DeleteTicket).Methods("DELETE")
		router.HandleFunc("/tickets/{id}", apiServer.EditTicket).Methods("PUT")
		router.HandleFunc("/tickets", apiServer.AddTicket).Methods("POST")
	*/

	router.HandleFunc("/productservicecategories/{id}", apiServer.GetProductServiceCategories).Methods("GET")
	router.HandleFunc("/productservicecategories", apiServer.GetAllProductServiceCategories).Methods("GET")
	router.HandleFunc("/servicecataloglvl1/{id}", apiServer.GetServiceCatalogLvL1).Methods("GET")
	router.HandleFunc("/servicecataloglvl2/{id}", apiServer.GetServiceCatalogLvL2).Methods("GET")
	router.HandleFunc("/servicecataloglvl3/{id}", apiServer.GetServiceCatalogLvL3).Methods("GET")
	router.HandleFunc("/servicecataloglvl4/{id}", apiServer.GetServiceCatalogLvL4).Methods("GET")
	router.HandleFunc("/servicecataloglvl5/{id}", apiServer.GetServiceCatalogLvL5).Methods("GET")
	router.HandleFunc("/servicecataloglvl6/{id}", apiServer.GetServiceCatalogLvL6).Methods("GET")

	router.HandleFunc("/productservicecategories", apiServer.AddProductServiceCategories).Methods("POST")
	router.HandleFunc("/servicecataloglvl1", apiServer.AddServiceCatalogLvL1).Methods("POST")
	router.HandleFunc("/servicecataloglvl2", apiServer.AddServiceCatalogLvL2).Methods("POST")
	router.HandleFunc("/servicecataloglvl3", apiServer.AddServiceCatalogLvL3).Methods("POST")
	router.HandleFunc("/servicecataloglvl4", apiServer.AddServiceCatalogLvL4).Methods("POST")
	router.HandleFunc("/servicecataloglvl5", apiServer.AddServiceCatalogLvL5).Methods("POST")
	router.HandleFunc("/servicecataloglvl6", apiServer.AddServiceCatalogLvL6).Methods("POST")

	router.HandleFunc("/productservicecategories/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl1/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl2/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl3/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl4/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl5/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")
	router.HandleFunc("/servicecataloglvl6/{id}", apiServer.DeleteProductServiceCategories).Methods("DELETE")

	//router.HandleFunc("/orderCodeEvent", apiServer.ConsumeOrderCode).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
