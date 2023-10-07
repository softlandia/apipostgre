package server

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, username, password, host, port, database string) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)

}

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps < 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}

// 	Conn, err := pgx.Connect(context.Background(), "postgres://postgres:123@localhost:5432/postgres")
// if err != nil{
// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 	os.Exit(1)
// }
// defer Conn.Close(context.Background())
// var greeting string
// err = Conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(greeting)

// Conn.QueryRow(context.Background(), "select Hello Sobaka").Scan(&greeting)
// fmt.Println(greeting)
