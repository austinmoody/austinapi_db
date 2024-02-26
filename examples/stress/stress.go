package main

import (
	"context"
	"github.com/austinmoody/austinapi_db/austinapi_db"
	"github.com/jackc/pgx/v5"
	"log"
	"math/rand"
	"time"
)

func main() {
	populateRandom(50)
}

func populateRandom(numberOfRecords int) {
	connStr := austinapi_db.GetDatabaseConnectionString()
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
	for i := 0; i < numberOfRecords; i++ {
		randomHours := rng.Intn(8674) + 1

		params := austinapi_db.SaveStressParams{
			Date:               time.Now().Add(time.Hour * (-1 * time.Duration(randomHours))),
			HighStressDuration: rng.Int63n(1000) - 1,
		}

		err = apiDb.SaveStress(ctx, params)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}

		result, err := apiDb.GetStressByDate(ctx, params.Date)
		if err != nil {
			log.Fatalf("Error Getting By Date: %v", err)
		}

		log.Printf("\n%v\n", result)

	}
}
