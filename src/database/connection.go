package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func getStringConnection() (connectionString string, erro error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("src/config")

	if erro = viper.ReadInConfig(); erro != nil {
		log.Fatalf("Erro ao tentar ler o árquivo de conexão: Erro: %v", erro)
	}

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	hostname := viper.GetString("database.hostname")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database_name")

	connectionString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
	return

}
func GetConnection() (db *sql.DB, erro error) {
	var once sync.Once

	once.Do(func() {

		connectionString, erro := getStringConnection()

		db, erro = sql.Open("postgres", connectionString)

		if erro != nil {
			log.Fatal("Error: dados de conexão com o banco invalidos.  ", erro)
			return
		}

		erro = db.Ping()
		if erro != nil {
			log.Fatal("Error: Não foi possível se conectar com o banco de dados. ", erro)
			return
		}
	})

	return

}
