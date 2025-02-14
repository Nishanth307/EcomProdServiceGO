package clickhouse

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouseRepo struct{
	Conn *sql.DB
}

func Connect(clickhouseURI string) (*ClickHouseRepo, error) {
	fmt.Println("ClickHouse Connection")

	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{clickhouseURI},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60, // Query timeout in seconds
		},
		DialTimeout: 10 * time.Second, // Connection timeout
	})

	if err := conn.Ping(); err != nil {
		log.Printf("Failed to ping ClickHouse: %v", err)
		return nil, err
	}

	fmt.Println("Connected to ClickHouse!")
	return &ClickHouseRepo{Conn: conn}, nil
}