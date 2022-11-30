package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

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

	router := mux.NewRouter()

	router.HandleFunc("/products", Getproducts).Methods("GET")

	router.HandleFunc("/products/{id}", GetproductByid).Methods("GET")

	router.HandleFunc("/products/{id}/reviews", Getallreviews).Methods("GET")

	router.HandleFunc("/products/create", Createproduct).Methods("POST")

	router.HandleFunc("/products/{id}/reviews/create", CreateReview).Methods("POST")

	router.HandleFunc("/products/{productid}/reviews/{reviewid}", Updatereview).Methods("PUT")

	// fmt.Println("Getting products")
	// db.query

	// router := mux.NewRouter()

	// router.HandleFunc('/products/',Getproducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

	fmt.Println("CREATED")

}

func Updatereview(w http.ResponseWriter, r *http.Request) {
	db := createDB()

	productid := mux.Vars(r)["productid"]

	reviewid := mux.Vars(r)["reviewid"]

	var rev Rating

	json.NewDecoder(r.Body).Decode(&rev)
	reviewstring := rev.Review
	rating_reference := []Rating{}
	db.Find(&rating_reference)
	db.Where("id = ? and product_id = ?", reviewid, productid).Find(&rating_reference).Update("review", reviewstring)
	return

}

func Getproducts(w http.ResponseWriter, r *http.Request) {
	db := createDB()

	prods := []Product{}
	fmt.Println("Getting products.....")
	db.Find(&prods)
	// for _, u := range prods {
	// 	fmt.Printf("\n%v\n", u)
	// }

	w.Header().Set("Content-type", "appilication/json")
	json.NewEncoder(w).Encode(prods)

}

func GetproductByid(w http.ResponseWriter, r *http.Request) {

	db := createDB()
	productid := mux.Vars(r)["id"]

	var prod Product

	db.Model(&Product{}).First(&prod, productid)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(prod)
	return

}

func Getallreviews(w http.ResponseWriter, r *http.Request) {

	db := createDB()

	productid := mux.Vars(r)["id"]

	rating_reference := []Rating{}
	db.Find(&rating_reference)
	var review_slice []string
	db.Where("id = ?", productid).Find(&rating_reference)
	for _, val := range rating_reference {
		review_slice = append(review_slice, val.Review)
	}
	//fmt.Println(review_slice)

	// db.Preload("Myratings").Find(&prod)
	// for _, val := range prod {
	// 	fmt.Println(val.Myratings)
	// }

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(review_slice)

}

func CreateReview(w http.ResponseWriter, r *http.Request) {

	db := createDB()

	var rev Rating

	w.Header().Set("content-type", "application/json")

	json.NewDecoder(r.Body).Decode(&rev)
	db.Save(&rev)
	json.NewEncoder(w).Encode(rev)
	return
}

func Createproduct(w http.ResponseWriter, r *http.Request) {

	db := createDB()

	w.Header().Set("content-type", "application/json")

	var prod Product
	json.NewDecoder(r.Body).Decode(&prod)

	db.Save(&prod)
	json.NewEncoder(w).Encode(prod)
	return
}
func createDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=root sslmode=disable")
	f := flag.String("dbname", "postgres", "added by gst")
	checkErr(err)

	// db.DropTable(&Product{})
	// db.CreateTable(&Product{})
	// db.DropTable(&Rating{})
	// db.CreateTable(&Rating{})
	// db.DropTable(&Variant{})
	// db.CreateTable(&Variant{})

	flag.Parse()
	fmt.Println(*f)
	// db.Save(&Product{
	// 	Name:        "Guitar",
	// 	Description: "Water resistent and made with fine wood",
	// 	Category:    "Music Instrument",
	// 	Quantity:    100,
	// 	Price:       5430,
	// 	Variants: []Variant{
	// 		{Color: "Black"},
	// 		{Color: "Brown"},
	// 	},
	// 	Myratings: []Rating{
	// 		{Name: "GST",
	// 			Review: "This guitar strings are so good",
	// 			Ratin:  6,
	// 		},
	// 	},
	// })
	return db

}

//
