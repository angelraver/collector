package dataBase

import (
	"coleccionista/config"
	"context"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// Returns an open connection
func Conectar() *sql.DB {
	var (
		dbUser                 = "postgres"                                // e.g. 'my-db-user'
		dbPwd                  = "Pachanda1!"                              // e.g. 'my-db-password'
		dbName                 = "postgres"                                // e.g. 'my-database'
		instanceConnectionName = "coleccionista:us-central1:coleccionista" // e.g. 'project:region:instance'
		usePrivate             = ""
	)

	dsn := fmt.Sprintf("user=%s password=%s database=%s", dbUser, dbPwd, dbName)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil
	}
	var opts []cloudsqlconn.Option
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
	}
	d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
	if err != nil {
		return nil
	}
	// Use the Cloud SQL connector to handle connecting to the instance.
	// This approach does *NOT* require the Cloud SQL proxy.
	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}
	dbURI := stdlib.RegisterConnConfig(config)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil
	}
	return dbPool
}

func ConectarOld() *sql.DB {
	var (
		host     = config.Get("POSTGRE_SERVER")
		port     = config.Get("POSTGRE_PORT")
		user     = config.Get("POSTGRE_USER")
		password = config.Get("POSTGRE_PASS")
		dbname   = config.Get("POSTGRE_DBNAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port="+port+" user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	dbPool, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return dbPool
}

func ConnectTCPSocket() *sql.DB {
	var (
		dbUser    = config.Get("POSTGRE_USER")
		dbPwd     = config.Get("POSTGRE_PASS")
		dbTCPHost = config.Get("POSTGRE_SERVER")
		dbPort    = config.Get("POSTGRE_PORT")
		dbName    = config.Get("POSTGRE_DBNAME")
	)

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTCPHost, dbUser, dbPwd, dbPort, dbName)
	dbPool, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		fmt.Println("errorrrrr!!!!!")
		fmt.Println(err)
		panic(err.Error())
	}

	return dbPool
}
