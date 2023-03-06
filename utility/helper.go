package utility

import (
	"context"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Fetch all data
func FetchMany(w http.ResponseWriter, c *mongo.Collection, s interface{}) error {
	opts := options.Find().SetSort(bson.D{{Key: "created", Value: -1}})
	cursor, err := c.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	// Loop through the cursor and create a slice of items
	result := reflect.ValueOf(s).Elem()
	for cursor.Next(context.Background()) {
		item := reflect.New(result.Type().Elem()).Interface()
		err := cursor.Decode(item)
		if err != nil {
			return err
		}
		result.Set(reflect.Append(result, reflect.ValueOf(item).Elem()))
	}

	return err
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

//
