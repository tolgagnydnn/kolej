package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Okul model
type Okul struct {
	ID      int    `json:"id,omitempty"`
	Adi     string `json:"adi,omitempty"`
	Adres   string `json:"adres,omitempty"`
	Telefon string `json:"telefon,omitempty"`
	Website string `json:"website,omitempty"`
	Sehir   string `json:"sehir,omitempty"`
	Semt    string `json:"semt,omitempty"`
}

var mongoClient *mongo.Client
var liste = make([]Okul, 0)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@mongo:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client

	log.Println("Connected to MongoDB!")
}

// Liste tüm okulların listesini getirir.
// Dönüş değeri Okul modeli dizisidir.
func (o *Okul) Liste() []*Okul {
	collection := mongoClient.Database("kolej").Collection("okullar")

	var liste []*Okul

	cur, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var okul Okul
		err := cur.Decode(&okul)
		if err != nil {
			log.Fatal(err)
		}

		liste = append(liste, &okul)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", liste)
	return liste
}

// Ekle veritabanına yeni bir okul bilgisi ekler.
func (o *Okul) Ekle(okul Okul) bool {
	collection := mongoClient.Database("kolej").Collection("okullar")

	res, err := collection.InsertOne(context.TODO(), okul)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a Single Document: ", res.InsertedID)
	return true
}
