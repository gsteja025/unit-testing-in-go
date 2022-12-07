package handlers

import (
	"encoding/json"
	"net/http"
	"temp/schema"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (s Dbserver) Getallreviews(w http.ResponseWriter, r *http.Request) {

	productid := mux.Vars(r)["id"]

	rating_reference := []schema.Rating{}
	s.Db.Find(&rating_reference)
	var review_slice []string
	s.Db.Where("id = ?", productid).Find(&rating_reference)
	for _, val := range rating_reference {
		review_slice = append(review_slice, val.Review)
	}
}

func (s Dbserver) CreateReview(w http.ResponseWriter, r *http.Request) {

	var rev schema.Rating

	w.Header().Set("content-type", "application/json")

	json.NewDecoder(r.Body).Decode(&rev)
	s.Db.Save(&rev)
	json.NewEncoder(w).Encode(rev)
	return
}

func (s Dbserver) Updatereview(w http.ResponseWriter, r *http.Request) {

	productid := mux.Vars(r)["productid"]

	reviewid := mux.Vars(r)["reviewid"]

	var rev schema.Rating

	json.NewDecoder(r.Body).Decode(&rev)
	reviewstring := rev.Review
	rating_reference := []schema.Rating{}
	s.Db.Find(&rating_reference)
	s.Db.Where("id = ? and product_id = ?", reviewid, productid).Find(&rating_reference).Update("review", reviewstring)
	return

}
