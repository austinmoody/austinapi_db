package main

import (
	"context"
	"fmt"
	"github.com/austinmoody/austinapi_db/austinapi_db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"log"
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

	dt := time.Now()

	params := austinapi_db.AddSleepRatingParams{
		Date: pgtype.Date{
			Time:             dt,
			InfinityModifier: 0,
			Valid:            true,
		},
		Rating: pgtype.Int4{
			Int32: 55,
			Valid: true,
		},
	}

	err = apiDb.AddSleepRating(ctx, params)
	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}

	durationParams := austinapi_db.AddSleepDurationParams{
		Date: pgtype.Date{
			Time:  dt,
			Valid: true,
		},
		TotalDuration: pgtype.Int8{
			Int64: 123,
			Valid: true,
		},
	}

	err = apiDb.AddSleepDuration(ctx, durationParams)
	if err != nil {
		log.Fatalf("Insert error: %v", err)
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
