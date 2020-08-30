package database

import (
	"context"
	"log"
	"oblique/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertWallet gets wallet pointer and returns pointer to mongo.InsertOneResult
func InsertWallet(wallet *model.Wallet) *mongo.InsertOneResult {
	log.Println("Database: InsertWallet")
	collection := client.Database("oblique-dev").Collection("wallets")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, wallet)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

// GetWallet gets id of wallet and pointer to wallet for decode result into it
func GetWallet(id primitive.ObjectID) (*error, *model.Wallet) {
	log.Println("Database: GetWallet")

	collection := client.Database("oblique-dev").Collection("wallets")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var wallet *model.Wallet
	err := collection.FindOne(ctx, model.Wallet{ID: id}).Decode(&wallet)
	if err != nil {
		log.Println(err)
		return &err, nil
	}

	return nil, wallet
}

// GetWallets fetch all wallets from db and returns them and error
func GetWallets() (*error, *[]model.Wallet) {
	log.Println("Database: GetWallets")

	collection := client.Database("oblique-dev").Collection("wallets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return &err, nil
	}
	defer cursor.Close(ctx)

	var wallets *[]model.Wallet
	for cursor.Next(ctx) {
		var wallet model.Wallet
		err = cursor.Decode(&wallet)
		if err != nil {
			log.Println(err)
			return &err, nil
		}

		*wallets = append(*wallets, wallet)
	}

	err = cursor.Err()
	if err != nil {
		log.Println(err)
		return &err, nil
	}

	return nil, wallets
}

// UpdateWallet gets the id and bson object that contains the updates fields. Returns the *mongo.UpdateResult
func UpdateWallet(id primitive.ObjectID, update bson.D) *mongo.UpdateResult {
	log.Println("Database: UpdateWallet")

	collection := client.Database("oblique-dev").Collection("wallets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}
