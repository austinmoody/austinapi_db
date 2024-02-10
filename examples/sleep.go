package main

import (
	"context"
	"fmt"
	"github.com/austinmoody/austinapi_db/austinapi_db"
	"github.com/google/uuid"
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

	myUuid := uuid.New()
	mySleep, err := apiDb.GetSleep(ctx, myUuid)
	if err != nil {
		log.Fatalf("Error Getting Sleep By Date: %v", err)
	}

	log.Println(mySleep)

}

func oldmain() {
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
	for i := 0; i < 50; i++ {
		randomHours := rng.Intn(8674) + 1

		sleepParams := austinapi_db.SaveSleepParams{
			Date:       time.Now().Add(time.Hour * (-1 * time.Duration(randomHours))),
			Rating:     rng.Int63n(100) - 1,
			TotalSleep: rng.Intn(50000) - 100,
			LightSleep: rng.Intn(50000) - 200,
			DeepSleep:  rng.Intn(50000) - 300,
			RemSleep:   rng.Intn(50000) - 400,
		}

		err = apiDb.SaveSleep(ctx, sleepParams)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}

		mySleep, err := apiDb.GetSleepByDate(ctx, sleepParams.Date)
		if err != nil {
			log.Fatalf("Error Getting Sleep By Date: %v", err)
		}

		log.Printf("\nSleep: %v\n", mySleep)

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
