package main

import (
	"context"
	"fmt"
	"github.com/austinmoody/austinapi_db/austinapi_db"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	connStr := GetDatabaseConnectionString()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer conn.Close(ctx)

	apiDb := austinapi_db.New(conn)

	// Loop 5 times, randomly generate data
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < 5; i++ {
		randomHours := rng.Int()

		sleepParams := austinapi_db.SaveSleepParams{
			Date:          time.Now().Add(time.Hour * time.Duration(randomHours)),
			Rating:        rng.Int31(),
			TotalDuration: rng.Int31(),
			NumberSleeps:  rng.Int31(),
		}

		err = apiDb.SaveSleep(ctx, sleepParams)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}

		mySleep, err := apiDb.GetSleepByDate(ctx, sleepParams.Date)
		if err != nil {
			log.Fatalf("Error Getting Sleep By Date: %v", err)
		}

		log.Printf("\nSleep Rating: %v\n", mySleep.Rating)

	}

}

func GetDatabaseConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	sslMode := os.Getenv("DATABASE_SSLMODE")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		databaseHost,
		databasePort,
		databaseUser,
		databasePassword,
		databaseName,
		sslMode,
	)
}
