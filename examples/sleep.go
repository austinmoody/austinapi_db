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

	//populateRandom()

	sleepList := newListSleepExample()

	log.Printf("Sleep List %v\n", sleepList.Data[0].ID)

	nextSleepList := listSleepNext(sleepList.NextToken)

	log.Printf("Sleep List %v\n", sleepList.Data[0].ID)

	previousSleepList := listSleepPrevious(nextSleepList.PreviousToken)

	log.Printf("Sleep List %v\n", sleepList.Data[0].ID)

	log.Println(previousSleepList.PreviousToken)
}

func listSleepPrevious(previousToken string) SleepList {
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
	numbers := s.Decode(previousToken)

	params := austinapi_db.ListSleepPreviousParams{
		ID:    int64(numbers[0]),
		Limit: 10,
	}
	sleeps, err := apiDb.ListSleepPrevious(ctx, params)
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

	mySleeps := austinapi_db.ListSleepPreviousRowToSleep(sleeps)

	sleepList := SleepList{
		Data:          mySleeps,
		NextToken:     nextId,
		PreviousToken: previousId,
	}

	return sleepList
}

func listSleepNext(nextToken string) SleepList {
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
	numbers := s.Decode(nextToken)

	params := austinapi_db.ListSleepNextParams{
		ID:    int64(numbers[0]),
		Limit: 10,
	}
	sleeps, err := apiDb.ListSleepNext(ctx, params)
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

	mySleeps := austinapi_db.ListSleepNextRowToSleep(sleeps)

	sleepList := SleepList{
		Data:          mySleeps,
		NextToken:     nextId,
		PreviousToken: previousId,
	}

	return sleepList
}

func newListSleepExample() SleepList {
	connStr := GetDatabaseConnectionString()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer conn.Close(ctx)

	apiDb := austinapi_db.New(conn)

	sleeps, err := apiDb.ListSleep(ctx, 10)
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

	log.Printf("Previous ID: %v\n", previousId)
	log.Printf("Next ID: %v\n", nextId)

	mySleeps := austinapi_db.ListSleepRowToSleep(sleeps)

	sleepList := SleepList{
		Data:          mySleeps,
		NextToken:     nextId,
		PreviousToken: previousId,
	}

	return sleepList

	//jsonBytes, err := json.Marshal(sleepList)
	//if err != nil {
	//	log.Fatalf("error marshaling JSON response: %v", err)
	//	return
	//}
	//log.Println(string(jsonBytes))

}

func ConvertListSleepNextToSleep(rows []austinapi_db.ListSleepNextRow) []austinapi_db.Sleep {
	sleeps := make([]austinapi_db.Sleep, len(rows))
	for i, row := range rows {
		sleeps[i] = austinapi_db.Sleep{
			ID:               row.ID,
			Date:             row.Date,
			Rating:           row.Rating,
			TotalSleep:       row.TotalSleep,
			DeepSleep:        row.DeepSleep,
			LightSleep:       row.LightSleep,
			RemSleep:         row.RemSleep,
			CreatedTimestamp: row.CreatedTimestamp,
			UpdatedTimestamp: row.UpdatedTimestamp,
		}
	}
	return sleeps
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
