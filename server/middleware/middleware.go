package middleware

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	loadEnv()
	creteDBInit()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading the env file")
	}
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func TaskComplete(w http.ResponseWriter, r *http.Request) {

}

func UndoTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

}
