package main

import (
	"context"
	"github.com/austinmoody/austinapi_db/austinapi_db"
	"github.com/jackc/pgx/v5"
	"github.com/sqids/sqids-go"
	"log"
	"math/rand"
	"time"
)

var (
	IdHasher   sqids.Sqids
	SqidLength string
)

func init() {
	var sqidLength uint8 = 20
	s, _ := sqids.New(sqids.Options{
		MinLength: sqidLength,
		Alphabet:  "usr4Z5gvSKhqpIt3BTAYVnwH8FQixC6G0cLNJ7fd9b1mlWEkOXz2RyjPoeUMDa",
	})

	IdHasher = *s

}

func GetIdFromToken(token string) int64 {
	nextTokenSlice := IdHasher.Decode(token)
	return int64(nextTokenSlice[0])
}

func main() {
	//populateRandom()

	GetList("", "")
}

func populateRandom() {
	connString := austinapi_db.GetDatabaseConnectionString()

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer conn.Close(ctx)

	apiDb := austinapi_db.New(conn)

	// Loop & randomly generate data
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < 50; i++ {
		randomHours := rng.Intn(8674) + 1

		params := austinapi_db.SavePreparednessParams{
			Date:   time.Now().Add(time.Hour * (-1 * time.Duration(randomHours))),
			Rating: rng.Intn(100) - 1,
		}

		err = apiDb.SavePreparedness(ctx, params)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}

		mySleep, err := apiDb.GetPreparednessByDate(ctx, params.Date)
		if err != nil {
			log.Fatalf("Error Getting Preparedness By Date: %v", err)
		}

		log.Printf("\nPreparedness: %v\n", mySleep)

	}

}

func GetList(queryType string, queryToken string) {

	var ListRowLimit int32 = 5

	params := austinapi_db.ListPreparednessParams{
		QueryType: "",
		InputID:   0,
		RowLimit:  ListRowLimit,
	}

	switch queryType {
	case "next_token":
		params.QueryType = "NEXT"
		params.InputID = GetIdFromToken(queryToken)
	case "previous_token":
		params.QueryType = "PREVIOUS"
		params.InputID = GetIdFromToken(queryToken)
	}

	connStr := austinapi_db.GetDatabaseConnectionString()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	defer conn.Close(ctx)

	apiDb := austinapi_db.New(conn)

	results, err := apiDb.ListPreparedness(ctx, params)
	if err != nil {
		log.Fatalf("issue pulling preparedness: %v", err)
	}

	log.Printf("Total results %d\n", len(results))

}
