package main

import (
	"context"
	"fmt"
	"github.com/austinmoody/austinapi_db/austinapi_db"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sqids/sqids-go"
	"log"
	"math/rand"
	"os"
	"time"
)

type SleepList struct {
	Data          []austinapi_db.Sleep `json:"data"`
	NextToken     string               `json:"next_token"`
	PreviousToken string               `json:"previous_token"`
}

func main() {

	// TODO how to rename SleepsRow to something better?
	//populateRandom()

	//sleepList := newListSleepExample()
	sleepList := AustinTest("", "")

	log.Printf("Sleep List %v\n", sleepList.Data[0].ID)

	//nextSleepList := listSleepNext(sleepList.NextToken)
	nextSleepList := AustinTest("NEXT", sleepList.NextToken)

	log.Printf("Sleep List %v\n", sleepList.Data[0].ID)

	//previousSleepList := listSleepPrevious(nextSleepList.PreviousToken)
	previousSleepList := AustinTest("PREVIOUS", nextSleepList.PreviousToken)

	log.Printf("Sleep List %v\n", sleepList.Data[0].ID)

	log.Println(previousSleepList.PreviousToken)
}

func AustinTest(queryType string, token string) SleepList {
	connStr := GetDatabaseConnectionString()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer conn.Close(ctx)

	apiDb := austinapi_db.New(conn)

	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})
	number := s.Decode(token)
	var idFromToken int64
	if len(number) == 0 {
		idFromToken = 0
	} else {
		idFromToken = int64(number[0])
	}

	params := austinapi_db.SleepsParams{
		QueryType: queryType,
		InputID:   idFromToken,
		RowLimit:  10,
	}

	sleeps, err := apiDb.Sleeps(ctx, params)
	if err != nil {
		log.Fatalf("Error getting list of sleep %v\n", err)
	}

	var previousId string
	if sleeps[0].PreviousID < 1 {
		previousId = ""
	} else {
		s, _ := sqids.New(sqids.Options{
			MinLength: 10,
		})
		id, _ := s.Encode([]uint64{uint64(sleeps[0].PreviousID)})
		previousId = id
	}

	var nextId string
	if sleeps[len(sleeps)-1].NextID < 1 {
		nextId = ""
	} else {
		s, _ := sqids.New(sqids.Options{
			MinLength: 10,
		})
		id, _ := s.Encode([]uint64{uint64(sleeps[len(sleeps)-1].NextID)})
		nextId = id
	}

	mySleeps := austinapi_db.AustinRowToSleep(sleeps)

	sleepList := SleepList{
		Data:          mySleeps,
		NextToken:     nextId,
		PreviousToken: previousId,
	}

	return sleepList
}

func getSleepExample() {
	//connStr := GetDatabaseConnectionString()
	//ctx := context.Background()
	//
	//conn, err := pgx.Connect(ctx, connStr)
	//if err != nil {
	//	log.Fatalf("DB Connection error: %v", err)
	//}
	//defer conn.Close(ctx)
	//
	//apiDb := austinapi_db.New(conn)
	//
	//myUuid := uuid.New()
	//mySleep, err := apiDb.GetSleep(ctx, myUuid)
	//if err != nil {
	//	log.Fatalf("Error Getting Sleep By Date: %v", err)
	//}
	//
	//log.Println(mySleep)
}

func populateRandom() {
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
