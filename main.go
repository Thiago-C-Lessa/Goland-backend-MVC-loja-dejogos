package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	//carrega .env
	errGetEnv := godotenv.Load();
	if errGetEnv != nil{
		log.Fatal(errGetEnv)
	}
	
	var PORT = os.Getenv("PORT");
	fmt.Println(PORT)
}