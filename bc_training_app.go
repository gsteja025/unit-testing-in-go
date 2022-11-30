package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	prod "temp/handlers"
	rev "temp/handlers"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Variant struct {
	gorm.Model
	ProductID uint   `json:"productid"`
	Color     string `json:"color"`
	Image     string `sql:"type:VARCHAR(255)" json:"image"`
}
type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	Image       string `sql:"type:VARCHAR(255)" json:"image"`
	Variants    []Variant
	Myratings   []Rating
}

type Rating struct {
	gorm.Model
	ProductID uint   `json:"productid"`
	Review    string `json:"review"`
	Ratin     uint
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	db, err := gorm.Open("postgres", "user=postgres password=root sslmode=disable")
	f := flag.String("dbname", "postgres", "added by gst")

	if err != nil {
		panic(err.Error())
	}

	router := mux.NewRouter()
	flag.Parse()
	fmt.Println(*f)

	var s prod.Dbserver = prod.Dbserver{db}
	router.HandleFunc("/products", s.Getproducts).Methods("GET")

	router.HandleFunc("/products/{id}", s.GetproductByid).Methods("GET")

	router.HandleFunc("/products/{id}/reviews", s.Getallreviews).Methods("GET")

	router.HandleFunc("/products/create", s.Createproduct).Methods("POST")

	router.HandleFunc("/products/{id}/reviews/create", rev.CreateReview).Methods("POST")

	router.HandleFunc("/products/{productid}/reviews/{reviewid}", rev.Updatereview).Methods("PUT")

	// fmt.Println("Getting products")
	// db.query

	// router := mux.NewRouter()

	// router.HandleFunc('/products/',Getproducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

	fmt.Println("CREATED")

}
