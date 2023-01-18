package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mahal007/restrauntService/db"
	"github.com/mahal007/restrauntService/handlers"
	"github.com/rakyll/statik/fs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	_ "github.com/mahal007/restrauntService/swaggerui" //swagger static files
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// DB connection
	db := db.Connect(mysqlUsername, mysqlPassword, mysqlDatabase)

	fmt.Println(db)
	// Restraunt routes Routes
	r.Route("/restraunt", func(r chi.Router) {

		//Menu Routes
		r.Route("/menu", func(r chi.Router) {
			r.Get("/", handlers.GetMenu(db))
			r.Post("/additem", handlers.AddItemToMenu(db))
		})

		// //Order Routes
		r.Route("/order", func(r chi.Router) {
			r.Get("/", handlers.GetAllOrder(db))
			r.Put("/{OrderId}", handlers.UpdateOrder(db))
		})
	})

	//Customer Routes
	r.Route("/customer", func(r chi.Router) {
		r.Post("/", handlers.CreateNewCustomer(db))
		r.Route("/{customerId}/order", func(r chi.Router) {
			r.Post("/", handlers.CreateNewOrder(db))
			r.Get("/", handlers.GetAllOrderByCustomerId(db))
			r.Get("/{orderId}", handlers.GetOrderById(db))
		})
	})

	//Added for swaggerUI
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	//swaggerDocs := http.FileServer(http.Dir("./swaggerui"))
	r.Handle("/swaggerui/*", http.StripPrefix("/swaggerui", staticServer))

	log.Printf("Server started on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalln("There's an error with the server", err)

	}
}
