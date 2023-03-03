package utility

import (
	"context"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Fetch all data
func FetchMany(w http.ResponseWriter, c *mongo.Collection, s interface{}) interface{} {
	cursor, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		ServerErr(w, err)
		return nil
	}
	defer cursor.Close(context.Background())

	// Loop through the cursor and create a slice of items
	result := reflect.ValueOf(s).Elem()
	for cursor.Next(context.Background()) {
		item := reflect.New(result.Type().Elem()).Interface()
		err := cursor.Decode(item)
		if err != nil {
			ServerErr(w, err)
			return nil
		}
		result.Set(reflect.Append(result, reflect.ValueOf(item).Elem()))
	}

	return result.Interface()
}

// Delete a item by id
func DeleteById(w http.ResponseWriter, r *http.Request, c *mongo.Collection) error {
	id := chi.URLParam(r, "id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = c.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
