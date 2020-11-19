package main
import(
	"SampleApp-CookBook/Server/helper"
	"SampleApp-CookBook/models"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"net/http"
)



func main(){


	router := mux.NewRouter()

	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("api/books/{id}",updateBook).Methods("PUT")
	router.HandleFunc("api/books/{id}",deleteBook).Methods("DELETE")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(router)))


}

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content/Type", "application/json")
	var bookset []models.Book

	collection := helper.DBConnection()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur,err := collection.Find(context.TODO(),bson.M{})
	if err != nil {
		helper.GetError(err,w)
		return
	}

	defer  cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var book models.Book

		// & character returns the memory address of the following variable.
		err := cur.Decode(&book) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		bookset = append(bookset, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(bookset) // encode similar to serialize process.

}

func getBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var book models.Book
	params :=mux.Vars(r)

	id,_ :=primitive.ObjectIDFromHex(params["id"])
	collection := helper.DBConnection()

	filter :=bson.M{"_id": id}
	err := collection.FindOne(context.TODO(),filter).Decode(&book)
	if err !=nil{
		helper.GetError(err,w)
		return
	}
	json.NewEncoder(w).Encode(book)
}


func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&book)

	// connect db
	collection := helper.DBConnection()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}


func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book models.Book

	collection := helper.DBConnection()

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	book.ID = id

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := helper.DBConnection()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}