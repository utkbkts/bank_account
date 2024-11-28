package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var conn *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	// pgxpool ile bağlantı oluştur
	conn, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)

	// Testleri çalıştır
	code := m.Run()

	// Bağlantıyı kapat
	conn.Close()

	// Çıkış
	os.Exit(code)
}
