package category

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oblique/database"
	"oblique/model"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllCategories")

	w.Header().Set("Content-Type", "application/json")

	err, categories := database.GetCategories()
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCategory")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	err, category := database.GetCategory(id)
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("AddCategory")

	params := r.URL.Query()

	var category model.Category

	category.Title = params.Get("title")
	category.ImageName = params.Get("imageName")
	category.Color = params.Get("color")

	// model.Categories = append(model.Categories, category)

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(model.Categories)
}

func GetMostUsedCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMostUsedCategories")

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(MostUsedCategories())
}

func GetCategoriesStatistic(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMostUsedCategories")

	w.Header().Set("Content-Type", "application/json")

	// var categories []model.Category
	// for _, category := range model.Categories {
	// 	category.Percantage = CalculatePercantage(category)
	// 	categories = append(categories, category)
	// }
	// json.NewEncoder(w).Encode(categories)

}
