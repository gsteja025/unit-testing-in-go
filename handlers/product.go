package handlers

import (
	"encoding/json"

	"fmt"
	"net/http"
	"temp/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type prods interface {
	//Createproduct(w http.ResponseWriter, r *http.Request)
	get(schema.Product) bool
}

type Dbserver struct {
	Db *gorm.DB
}

// type mystruct struct{

// }
type empty struct {
}

var s Dbserver

var url string = "http://localhost:8000/products/{name}"

func (e empty) get(prod schema.Product) bool {

	s.Db.Save(&prod)
	return true
}

func (s Dbserver) Getproducts(w http.ResponseWriter, r *http.Request) {

	prods := []schema.Product{}
	fmt.Println("Getting products.....")
	s.Db.Find(&prods)
	// for _, u := range prods {
	// 	fmt.Printf("\n%v\n", u)
	// }

	w.Header().Set("Content-type", "appilication/json")
	json.NewEncoder(w).Encode(prods)

}

func (e empty) Createproduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var prod schema.Product
	json.NewDecoder(r.Body).Decode(&prod)

	found := empty{}.get(prod)
	json.NewEncoder(w).Encode(prod)

	if found {
		
	}
	return
}
