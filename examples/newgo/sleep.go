package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/austinmoody/austinapi_db/austinapi_db"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
)

type Sleeps struct {
	Data      []austinapi_db.Sleep `json:"data"`
	NextToken string               `json:"next_token"`
}

// ObfuscateInt obfuscates an int32 into a base64-encoded string
func ObfuscateInt(num int32) string {

	num = num + 1000
	// Convert int32 to string
	str := strconv.FormatInt(int64(num), 10)

	// Convert string to byte slice
	bytes := []byte(str)

	// Base64 encode byte slice
	encoded := base64.StdEncoding.EncodeToString(bytes)

	return encoded
}

// DecodeObfuscatedInt decodes a base64-encoded string back to an int32
func DecodeObfuscatedInt(obfuscatedStr string) int32 {
	// Decode base64-encoded string
	decoded, err := base64.StdEncoding.DecodeString(obfuscatedStr)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return 0
	}

	// Convert byte slice to string
	str := string(decoded)

	// Parse string to int32
	decodedInt, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to int32:", err)
		return 0
	}

	decodedInt = decodedInt - 1000

	return int32(decodedInt)
}

func main() {
	getSleep(56)

	getSleeps("MTAwNQ==")
}

func getSleep(sleepId int64) {
	connStr := austinapi_db.GetDatabaseConnectionString()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer conn.Close(ctx)

	apiDb := austinapi_db.New(conn)

	result, err := apiDb.GetSleep(ctx, sleepId)
	if err != nil {
		log.Fatalf("Issue getting sleep from database: %v", err)
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("issue with creating json: %v", err)
	}

	log.Print(string(jsonBytes))
}

func getSleeps(nextToken string) {
	connStr := austinapi_db.GetDatabaseConnectionString()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer conn.Close(ctx)

	apiDb := austinapi_db.New(conn)

	params := austinapi_db.GetSleepsParams{
		RowLimit:  5,
		RowOffset: 0,
	}

	if nextToken != "" {
		params.RowOffset = DecodeObfuscatedInt(nextToken)
	}

	results, err := apiDb.GetSleeps(ctx, params)
	if err != nil {
		log.Fatalf("GetSleeps error: %v", err)
	}

	sleeps := Sleeps{
		Data:      results,
		NextToken: ObfuscateInt(params.RowOffset + params.RowLimit),
	}

	jsonBytes, err := json.Marshal(sleeps)
	if err != nil {
		log.Fatalf("JSON Marshal error: %v", err)
	}

	log.Print(string(jsonBytes))
}
