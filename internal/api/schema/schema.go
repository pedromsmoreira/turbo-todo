package schema

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"
)

func CreateSchema(cfg *configs.Config) error {
	if cfg.Database.SkipSchema {
		log.Println("schema >>> schema creation skipped...")
		return nil
	}
	connStr := fmt.Sprintf("postgresql://root@%s/defaultdb?sslmode=disable", cfg.Database.Host)
	config, err := pgx.ParseConfig(connStr)
	config.Database = "defaultdb"
	if err != nil {
		return errors.New(fmt.Sprintf("error schema configuration: %v", err))
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return errors.New(fmt.Sprintf("error connecting to the database: %v", err))
	}
	defer conn.Close(context.Background())

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		log.Println("schema >>> creating Database...")
		if _, err := tx.Exec(context.Background(),
			"CREATE DATABASE IF NOT EXISTS turbotodo;"); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return errors.New(fmt.Sprintf("error creating database: %v", err))
	}

	log.Println("schema >>> changing to turbotodo database...")
	config.Database = "turbotodo"
	conn, err = pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return errors.New(fmt.Sprintf("error connecting to turbotodo database: %v", err))
	}
	defer conn.Close(context.Background())

	log.Println("schema >>> creating turbotodo tables...")
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		log.Println("schema >>> creating todos table...")
		if _, err := tx.Exec(context.Background(),
			"CREATE TABLE IF NOT EXISTS todos (id STRING(36) PRIMARY KEY, datecreated TIMESTAMP, title STRING, content STRING, tags STRING[], status STRING, version INT8);"); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return errors.New(fmt.Sprintf("error creating todos table: %v", err))
	}

	return nil
}
