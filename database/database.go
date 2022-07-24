package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	Role         string             `bson:"role"`
	PasswordHash string             `bson:"password_hash"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    time.Time          `bson:"deleted_at,omitempty"`
}

type Person struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type Car struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
}

type Wallet struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Currency string             `bson:"currency"`
	Amount   int                `bson:"amount"`
	Time     time.Time          `bson:"time"`
}

func StartConnection() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://monggo:monggo@localhost:27017"))
	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to database")
		return nil
	}
	fmt.Println("Success to connect to database")

	db := client.Database("userdb")

	return db
}

func InsertUser() {
	ctx := context.Background()
	db := StartConnection()

	var person1 Person
	person1.FirstName = "Muhammad"
	person1.LastName = "Arash"

	var person2 Person
	person2.FirstName = "Dimas"
	person2.LastName = "Ramadhan"

	var person3 Person
	person3.FirstName = "Muhammad"
	person3.LastName = "Agung"

	fmt.Println("Run 1")
	_, err := db.Collection("people").InsertOne(ctx, person1)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Run 2")
	_, err = db.Collection("people").InsertOne(ctx, person2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Run 3")
	_, err = db.Collection("people").InsertOne(ctx, person3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert person success!")
}

func InsertCar() {
	ctx := context.Background()
	db := StartConnection()

	var car1 Car
	car1.UserID, _ = primitive.ObjectIDFromHex("629c0c563a3957cc3589fa82")
	car1.Name = "Avanza"

	var car2 Car
	car2.UserID, _ = primitive.ObjectIDFromHex("629c0c563a3957cc3589fa82")
	car2.Name = "Xenia"

	_, err := db.Collection("cars").InsertOne(ctx, car1)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("cars").InsertOne(ctx, car2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert car success!")
}

func InsertWallet() {
	ctx := context.Background()
	db := StartConnection()

	var wallet1 Wallet
	wallet1.Currency = "Dollar"
	wallet1.Amount = 200
	wallet1.Time = time.Now()

	_, err := db.Collection("wallets").InsertOne(ctx, wallet1)
	if err != nil {
		log.Fatal(err.Error())
	}

	time.Sleep(2 * time.Second)

	var wallet2 Wallet
	wallet2.Currency = "Rupiah"
	wallet2.Amount = 300
	wallet2.Time = time.Now()

	_, err = db.Collection("wallets").InsertOne(ctx, wallet2)
	if err != nil {
		log.Fatal(err.Error())
	}

	time.Sleep(2 * time.Second)

	var wallet3 Wallet
	wallet3.Currency = "Dollar"
	wallet3.Amount = 100
	wallet3.Time = time.Now()

	_, err = db.Collection("wallets").InsertOne(ctx, wallet3)
	if err != nil {
		log.Fatal(err.Error())
	}

	time.Sleep(2 * time.Second)

	var wallet4 Wallet
	wallet4.Currency = "Rupiah"
	wallet4.Amount = 500
	wallet4.Time = time.Now()

	_, err = db.Collection("wallets").InsertOne(ctx, wallet4)
	if err != nil {
		log.Fatal(err.Error())
	}

	var wallet5 Wallet
	wallet5.Currency = "Dollar"
	wallet5.Amount = 400
	wallet5.Time = time.Now()

	_, err = db.Collection("wallets").InsertOne(ctx, wallet5)
	if err != nil {
		log.Fatal(err.Error())
	}

	time.Sleep(2 * time.Second)

	var wallet6 Wallet
	wallet6.Currency = "Rupiah"
	wallet6.Amount = 600
	wallet6.Time = time.Now()

	_, err = db.Collection("wallets").InsertOne(ctx, wallet6)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert wallet success!")
}

func GetAllWallets() {
	ctx := context.Background()
	db := StartConnection()

	cursor, err := db.Collection("wallets").Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	var wallets []Wallet
	err = cursor.All(ctx, &wallets)
	if err != nil {
		panic(err)
	}

	for _, wallet := range wallets {
		fmt.Println(wallet.ID)
		fmt.Println(wallet.Currency)
		fmt.Println(wallet.Amount)
		fmt.Println(wallet.Time)
	}
}
func GetWalletsByCurrency() {
	ctx := context.Background()
	db := StartConnection()

	options := options.Find()
	options.SetLimit(2)
	options.SetSort(bson.D{
		{"amount", -1},
		{"time", -1},
		})	//	1 = Ascending & -1 = Descending

	filter := bson.M{"currency": "Dollar"}
	cursor, err := db.Collection("wallets").Find(ctx, filter, options)
	if err != nil {
		panic(err)
	}

	var wallets []Wallet
	err = cursor.All(ctx, &wallets)
	if err != nil {
		panic(err)
	}

	for _, wallet := range wallets {
		fmt.Println(wallet.ID)
		fmt.Println(wallet.Currency)
		fmt.Println(wallet.Amount)
		fmt.Println(wallet.Time)
	}
}

func UpdateWallet() {
	ctx := context.Background()
	db := StartConnection()
	
	var wallet Wallet
	wallet.ID, _ = primitive.ObjectIDFromHex("62ab12b10afa5dfd9b9e231f")
	wallet.Currency = "Dollar"
	wallet.Amount = -600
	wallet.Time = time.Now()

	filter := bson.M{"_id": wallet.ID}
	_, err := db.Collection("wallets").ReplaceOne(ctx, filter, wallet)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update wallet success!")
}

func GetUser() {
	ctx := context.Background()
	db := StartConnection()

	objectId, _ := primitive.ObjectIDFromHex("6289a403e94fe309ca893de5")

	filter := bson.M{"_id": objectId}
	cursor := db.Collection("customers").FindOne(ctx, filter)

	var data Person
	err := cursor.Decode(&data)
	if err != nil {
		fmt.Printf("error occured %v", err)
	}
	fmt.Println(data)
	fmt.Println(data.FirstName)
	fmt.Println(data.LastName)
}

func GetCarsByUserID() {
	ctx := context.Background()
	db := StartConnection()

	objectId, _ := primitive.ObjectIDFromHex("629c0c563a3957cc3589fa82")

	filter := bson.M{"user_id": objectId}
	cursor, err := db.Collection("cars").Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var cars []Car
	err = cursor.All(ctx, &cars)
	if err != nil {
		panic(err)
	}

	for _, car := range cars {
		fmt.Println(car.ID)
		fmt.Println(car.UserID)
		fmt.Println(car.Name)
	}
}

func GetUserByDeletedAt() {
	ctx := context.Background()
	db := StartConnection()

	deletedAt, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")

	filter := bson.M{"deleted_at": deletedAt}
	cursor := db.Collection("customers").FindOne(ctx, filter)

	var data User
	err := cursor.Decode(&data)
	if err != nil {
		fmt.Printf("error occured %v", err)
	}
	fmt.Println(data)
	fmt.Println(data.ID)
	fmt.Println(data.Name)
	fmt.Println(data.Email)
	fmt.Println(data.PasswordHash)
	fmt.Println(data.CreatedAt)
	fmt.Println(data.UpdatedAt)
	fmt.Println(data.DeletedAt)
}

func GetAllUsersByDeletedAt() {
	ctx := context.Background()
	db := StartConnection()

	// Filter By Single Parameter
	filter := bson.M{"deleted_at": nil}

	// Filter By Multiple Parameter
	// filter := bson.D{
	// 	{"deleted_at", nil},
	// 	{"firstname", "Muhammad"},
	// }
	// Filter By Multiple Parameter
	// filter := bson.D{
	// 	{"deleted_at", bson.D{{"$eq", nil}}},
	// 	{"firstname", bson.D{{"$eq", "Muhammad"}}},
	// }

	// // Filter By Multiple AND Parameter
	// filter := bson.M{"$and": bson.A{
	// 	bson.M{"firstname": "Muhammad"},
	// 	bson.M{"lastname": "Arash"},
	// }}

	// Filter By Multiple OR Parameter
	// filter := bson.M{"$or": bson.A{
	// 	bson.M{"firstname": "Muhammad"},
	// 	bson.M{"lastname": "Arash"},
	// }}

	cursor, err := db.Collection("customers").Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var people []Person
	err = cursor.All(ctx, &people)
	if err != nil {
		panic(err)
	}

	for _, person := range people {
		fmt.Println(person.ID)
		fmt.Println(person.FirstName)
		fmt.Println(person.LastName)
		fmt.Println("")
	}
}

func GetUserByEmail() {
	ctx := context.Background()
	db := StartConnection()

	var data Person
	filter := bson.M{"email": "admin@gmail.com"}
	err := db.Collection("customers").FindOne(ctx, filter).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Ini adalah admin", data)
}
