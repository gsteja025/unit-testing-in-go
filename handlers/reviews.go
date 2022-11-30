package handlers
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


func (s Dbserver) Getallreviews(w http.ResponseWriter, r *http.Request) {

	db := createDB()

	productid := mux.Vars(r)["id"]

	rating_reference := []Rating{}
	db.Find(&rating_reference)
	var review_slice []string
	db.Where("id = ?", productid).Find(&rating_reference)
	for _, val := range rating_reference {
		review_slice = append(review_slice, val.Review)
	}

func (s Dbserver) CreateReview(w http.ResponseWriter, r *http.Request) {

		db := createDB()
			
		var rev Rating
			
		w.Header().Set("content-type", "application/json")
			
		json.NewDecoder(r.Body).Decode(&rev)
		db.Save(&rev)
		json.NewEncoder(w).Encode(rev)
		return
}

func (s Dbserver) Updatereview(w http.ResponseWriter, r *http.Request) {
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