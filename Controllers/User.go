package Controllers

import (
	"Curd-Api-Http/Database"
	"Curd-Api-Http/Models"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)
var collection=Database.ConnectDB()

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var movie Models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	result, err := collection.InsertOne(context.TODO(),movie)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func GetMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	Id := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(Id, 10, 64)
	if checkMovie(id)==false{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Movie Doesnot exit!!")
		return
		}
	var movie Models.Movie
	filter := bson.M{"_id": id}
	collection.FindOne(context.TODO(), filter).Decode(&movie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func GetMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie []Models.Movie
	cur,_:=collection.Find(context.TODO(), bson.M{})
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var m Models.Movie
		cur.Decode(&m) 

		movie = append(movie, m)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func checkMovie(Id int64) bool{
	var movie Models.Movie
	filter := bson.M{"_id": Id}
	collection.FindOne(context.TODO(), filter).Decode(&movie)
	if movie.ID==0{
		return false
	}
	return true
}
func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.ParseInt(key, 10, 64)
	if checkMovie(id)==false{
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Movie Doesnot exit!!")
	return
	}
	collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User Deleted succesfully")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.ParseInt(key, 10, 64)
	if checkMovie(id)==false{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Movie Doesnot exit!!")
		return
		}
	var updatemovie Models.Movie
	json.NewDecoder(r.Body).Decode(&updatemovie)
	update:= bson.M{"year": updatemovie.Year, "title": updatemovie.Title, "director": updatemovie.Dricetor,"rating":updatemovie.Rating}
    collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Record Updated")

}