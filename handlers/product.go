package handlers

import (
	"encoding/json"

	"fmt"
	"net/http"
	"temp/schema"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Dbserver struct {
	Db *gorm.DB
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

func (s Dbserver) GetproductByid(w http.ResponseWriter, r *http.Request) {

	productid := mux.Vars(r)["id"]

	var prod schema.Product

	s.Db.Model(&schema.Product{}).First(&prod, productid)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(prod)
	return

}

func (s Dbserver) Createproduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var prod schema.Product
	json.NewDecoder(r.Body).Decode(&prod)

	s.Db.Save(&prod)
	json.NewEncoder(w).Encode(prod)
	return
}
