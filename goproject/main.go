package main

import(
	"log"
	DB "JSMPJ/goproject/src/system/db"
	"JSMPJ/goproject/src/system/app"
	"flag"
	"github.com/joho/godotenv"
	"os"
)
var port string
func init(){
  flag.StringVar(&port, "port", "8000", "Assiging port to server")
  flag.Parse()

  if  err := godotenv.Load("config.ini");err!=nil{
	  panic(err)
  }
  envPort := os.Getenv("PORT")
  if len(envPort)>0{
   port=envPort
  }
}
func main(){
	log.Println("JSMPJ Corporation")
	db, err := DB.Connect()
	if err != nil{
		panic(err)
	}
	s := app.NewServer()
	s.Init(port, db)
	s.Start()
}