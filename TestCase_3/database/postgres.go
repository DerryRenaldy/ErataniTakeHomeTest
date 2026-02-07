package database

import (
	"eratani_assesment_test/TestCase_3/configs"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConn struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

func ProvidePostgresConn(config *configs.Config) *PostgresConn {
	return &PostgresConn{
		Read:  CreatePostgresReadConn(*config),
		Write: CreatePostgresWriteConn(*config),
	}
}

func CreatePostgresReadConn(config configs.Config) *sqlx.DB {
	return CreatePostgresDBConnection(
		"read",
		config.DB.Postgres.Read.User,
		config.DB.Postgres.Read.Password,
		config.DB.Postgres.Read.Host,
		config.DB.Postgres.Read.Port,
		config.DB.Postgres.Read.Name,
		config.DB.Postgres.Read.SSLMode,
		config.DB.Postgres.Read.MaxConnLifetime,
		config.DB.Postgres.Read.MaxIdleConn,
		config.DB.Postgres.Read.MaxOpenConn)
}

func CreatePostgresWriteConn(config configs.Config) *sqlx.DB {
	return CreatePostgresDBConnection(
		"write",
		config.DB.Postgres.Write.User,
		config.DB.Postgres.Write.Password,
		config.DB.Postgres.Write.Host,
		config.DB.Postgres.Write.Port,
		config.DB.Postgres.Write.Name,
		config.DB.Postgres.Write.SSLMode,
		config.DB.Postgres.Write.MaxConnLifetime,
		config.DB.Postgres.Write.MaxIdleConn,
		config.DB.Postgres.Write.MaxOpenConn)
}

func CreatePostgresDBConnection(connType, username, password, host, port, dbName, sslmode string, maxConnLifetime time.Duration, maxIdleConn, maxOpenConn int) *sqlx.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		dbName,
		sslmode)

	if password == "" {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s",
			host,
			port,
			username,
			dbName,
			sslmode)
	}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed connecting to Postgres database (%s): %s", connType, err))
	}

	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return db
}
