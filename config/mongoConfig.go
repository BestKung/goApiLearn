package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func InitMongo(connection string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// เชื่อมต่อ MongoDB
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))

	fmt.Printf("db : %T", db)

	if err != nil {
		log.Fatal(err)
	}

	// ตรวจสอบการเชื่อมต่อ
	err = db.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Connect MongoDB failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	// เลือก Database และ Collection
	collection := db.Database("myshop").Collection("myshop")
	fmt.Println(collection)

	insertResult, err := collection.InsertOne(ctx, map[string]interface{}{
		"name":  "Best",
		"email": "best@example.com",
		"age":   25,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🟢 Inserted ID:", insertResult.InsertedID)

	var result map[string]interface{}
	err = collection.FindOne(ctx, map[string]interface{}{"name": "Best"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🔍 Found document:", result)
}
