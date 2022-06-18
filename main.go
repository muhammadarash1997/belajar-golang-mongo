package main

import (
	"belajar-golang-mongo/database"
	// "context"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "golang.org/x/crypto/bcrypt"
)

// Create Administrator if not already created
// func init() {
// 	ctx := context.Background()
// 	db := database.StartConnection()

// 	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("01010101"), bcrypt.MinCost)
// 	Administrator := database.User{
// 		Name:         "Administrator",
// 		Email:        "admin@gmail.com",
// 		Role:         "admin",
// 		PasswordHash: string(passwordHash),
// 		CreatedAt:    time.Now(),
// 		UpdatedAt:    time.Now(),
// 	}

// 	cursor, _ := db.Collection("customers").CountDocuments(ctx, bson.M{"role": "admin"})
// 	if cursor > 0 {
// 		return
// 	}

// 	db.Collection("customers").InsertOne(ctx, Administrator)
// }

func main() {
	// database.StartConnection()
	// database.InsertUser()
	// database.InsertCar()
	// database.InsertWallet()
	// database.GetAllWallets()
	// database.GetWalletsByCurrency()
	database.UpdateWallet()
	// database.GetUser()
	// database.GetCarsByUserID()
	// database.GetUserByDeletedAt()
	// database.GetAllUsersByDeletedAt()
	// database.GetUserByEmail()
}
