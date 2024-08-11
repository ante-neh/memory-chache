package main

import (
	"flag"
	"log"
	"os"
	"github.com/ante-neh/memory-chache/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	//Load environmental varaibles 
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Unable to load environmental varaibles")
	}

    //access port address and set it a default value
	port := os.Getenv("PORT")
	address := flag.String("address", port, "server address")
	flag.Parse()

	//create info and error loggers to handle errors gracefully
	infoLog := log.New(os.Stdout, "INFO: ", log.Ltime | log.Ldate) 
	errorLog := log.New(os.Stdout, "ERROR: ", log.Ltime | log.Ldate | log.Lshortfile) 

	//create a server type 
	app := server.NewServer(infoLog, errorLog, *address)

	//start the server
	server := app.Start()

	app.InfoLogger.Println("Server is running on port: ", *address)
	err = server.ListenAndServe()
	
	if err != nil{
		app.ErrorLogger.Fatal("Unable to start a server")
	}
}