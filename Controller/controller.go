package Controller

import (
	"../models"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)
//Connect to mongo and initialize the client
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
var client, _  = mongo.Connect(context.TODO(), clientOptions)
var cPlayers = client.Database("test").Collection("players")
var cGame = client.Database("test").Collection("game")

// send a payload of JSON content
func respondWithJSON(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)
	if _, err := w.Write(response); err != nil {
		msg := fmt.Sprintf("Can't insert result in DB. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
}

// send a JSON error message
func respondWithError(w http.ResponseWriter, message string) {
	respondWithJSON(w, map[string]string{"error": message})
}

//PLAYER REQUEST
//Add a new player in the list
func NewPlayer(w http.ResponseWriter, r *http.Request) {
	// Get a handle for your collection
	var newPlayer models.Player
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newPlayer); err != nil {
		msg := fmt.Sprintf("Invalid JSON. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	// Insert a single document
	if _, err := cPlayers.InsertOne(context.TODO(), newPlayer); err != nil {
		msg := fmt.Sprintf("Can't insert result in DB. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//Get data for a single player
func SeePlayer(w http.ResponseWriter, r *http.Request) {
	// Find a single document
	var result models.Player
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&result); err != nil {
		msg := fmt.Sprintf("Invalid JSON. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	filter := bson.D{{"nickname", result.Nickname}}
	if err := cPlayers.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		msg := fmt.Sprintf("Can't found result in DB. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		msg := fmt.Sprintf("Can't send JSON. Error: %s", err.Error())
		respondWithError(w, msg)
	}
}

//Get data for all the player's list
func SeePlayers(w http.ResponseWriter, r *http.Request) {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(5)
	var results []*models.Player
	// Finding multiple documents returns a cursor
	cur, err := cPlayers.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		msg := fmt.Sprintf("Can't find result in DB. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem models.Player
		if err := cur.Decode(&elem); err != nil {
			msg := fmt.Sprintf("Error: %s", err.Error())
			respondWithError(w, msg)
			return
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		msg := fmt.Sprintf("Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	// Close the cursor once finished
	if err := cur.Close(context.TODO()); err != nil {
		msg := fmt.Sprintf("Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	for _, p := range results {
		fmt.Printf("%+v\n", p)
	}
	w.Header().Set("Content-Type", "application/json")
	respondWithJSON(w, json.NewEncoder(w).Encode(results))
}
//Delete a player
func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	// Find a single document
	var result models.Player
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&result); err != nil {
		msg := fmt.Sprintf("Invalid JSON. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	filter := bson.D{{"nickname", result.Nickname}}

	if err := cPlayers.FindOneAndDelete(context.TODO(), filter).Decode(&result); err != nil {
		msg := fmt.Sprintf("Can't insert result in DB. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	respondWithJSON(w, json.NewEncoder(w).Encode(result))
}

//GAME REQUEST
//Add a new game
func NewGame(w http.ResponseWriter, r *http.Request) {
	// Get a handle for your collection
	var newGame models.Games
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newGame)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	// Insert a single document
	if _, err := cGame.InsertOne(context.TODO(), newGame);err != nil {
		msg := fmt.Sprintf("Can't insert item. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

//Get data for a single game
func SeeGame(w http.ResponseWriter, r *http.Request) {
	// Find a single document
	var result models.Games
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	filter := bson.D{{"idgame", result.IdGame}}
	err2 := cGame.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil {
		msg := fmt.Sprintf("Can't find item. Error: %s", err2.Error())
		respondWithError(w, msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		msg := fmt.Sprintf("Can't send JSON. Error: %s", err.Error())
		respondWithError(w, msg)
	}
}

//Get data for all the player's list
func SeeGames(w http.ResponseWriter, r *http.Request) {
	// Pass these options to the Find method
	findOptions := options.Find()
	var results []*models.Games
	// Finding multiple documents returns a cursor
	cur, err := cGame.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		msg := fmt.Sprintf("Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem models.Games
		err := cur.Decode(&elem)
		if err != nil {
			msg := fmt.Sprintf("Error: %s", err.Error())
			respondWithError(w, msg)
			return
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		msg := fmt.Sprintf("Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	// Close the cursor once finished
	if err := cur.Close(context.TODO()); err != nil {
		msg := fmt.Sprintf("Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	for _, p := range results {
		fmt.Printf("%+v\n", p)
	}
	respondWithJSON(w, json.NewEncoder(w).Encode(results))
}

//Delete a single game
func DeleteGame(w http.ResponseWriter, r *http.Request) {
	// Find a single document
	var result models.Games
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON. Error: %s", err.Error())
		respondWithError(w, msg)
		return
	}
	filter := bson.D{{"idgame", result.IdGame}}
	err2 := cGame.FindOneAndDelete(context.TODO(), filter).Decode(&result)
	if err2 != nil {
		msg := fmt.Sprintf("Can't find item. Error: %s", err2.Error())
		respondWithError(w, msg)
		return
	}
	fmt.Printf("Found and delete a single document: %+v\n", result)
	respondWithJSON(w, json.NewEncoder(w).Encode(result))
}
