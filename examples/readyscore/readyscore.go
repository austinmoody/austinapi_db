package main

import (
	"context"
	"encoding/json"
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

type ReadyScores struct {
	Data          []ReadyScore `json:"data"`
	NextToken     string       `json:"next_token"`
	PreviousToken string       `json:"previous_token"`
}

type ReadyScore struct {
	*austinapi_db.GetReadyScoresRow
}

type ReadyScoresResult []austinapi_db.GetReadyScoresRow
type ReadyScoreItem austinapi_db.GetReadyScoresRow

func (rsr ReadyScoresResult) ToReadyScores() ReadyScores {
	readyScores := ReadyScores{
		Data: make([]ReadyScore, len(rsr)),
	}
	for i, item := range rsr {
		readyScores.Data[i] = ReadyScore{&item}
	}
	return readyScores
}

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
	//`populateRandom()

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

		params := austinapi_db.SaveReadyScoreParams{
			Date:  time.Now().Add(time.Hour * (-1 * time.Duration(randomHours))),
			Score: rng.Intn(100) - 1,
		}

		err = apiDb.SaveReadyScore(ctx, params)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}

		dbResult, err := apiDb.GetReadyScoreByDate(ctx, params.Date)
		if err != nil {
			log.Fatalf("Error Getting Ready Score By Date: %v", err)
		}

		log.Printf("\nReady Score: %v\n", dbResult)

	}

}

func GetList(queryType string, queryToken string) {

	var ListRowLimit int32 = 5

	params := austinapi_db.GetReadyScoresParams{
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

	results, err := apiDb.GetReadyScores(ctx, params)
	if err != nil {
		log.Fatalf("issue pulling ready scores: %v", err)
	}

	log.Printf("Total results %d\n", len(results))

	rsr := ReadyScoresResult(results)
	readyScores := rsr.ToReadyScores()

	jsonBytes, err := json.Marshal(readyScores)
	if err != nil {
		log.Fatalf("error with json marshall: %v", err)
	}

	log.Print(string(jsonBytes))

}
