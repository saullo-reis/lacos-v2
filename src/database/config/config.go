package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	User string
	PostgresDriver string
	Host string
	Port string
	Password string
	DbName string
	DataSourceName string
) 

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error ao carregar variáveis de ambiente na configuração do banco "+err.Error())
	}

	User = os.Getenv("USERPOSTGRES")
	PostgresDriver = os.Getenv("POSTGRESDRIVER")
	Host = os.Getenv("HOSTPOSTGRES")
	Port = os.Getenv("PORTPOSTGRES")
	Password = os.Getenv("PASSWORDPOSTGRES")
	DbName = os.Getenv("DBNAME")

	DataSourceName = fmt.Sprintf("host=%s port=%s user=%s " + "password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
}

func ConnectDB() (*sql.DB, error){
	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db,err = sql.Open(PostgresDriver, DataSourceName)
		if err != nil {
			log.Println("Falha ao conectar com o banco de dados tentando novamente "+ err.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		err = db.Ping()
		if err == nil {
			return db, nil
		}
		log.Printf("Falha ao pingar o banco de dados "+ err.Error())
		time.Sleep(2 * time.Second)
		continue
	}
	return nil, fmt.Errorf("Falha ao conectar o banco de dados "+err.Error())
}
