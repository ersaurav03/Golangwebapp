package app

import(
	"JSMPJ/goproject/src/system/router"
	"github.com/go-xorm/xorm"
	"net/http"//https://golang.org/pkg/net/http/
	"log"
	"github.com/gorilla/handlers"
	"os"
	"time"
)
type Server struct{
	port string  //Creating type of servcer struct
	Db *xorm.Engine
}
func NewServer() Server{
   return Server{}//Return instance of Server
}
func (s *Server) Init(port string, db *xorm.Engine){
	log.Println("Initilizing a server")//method to initlizing server at port 8000
	s.port = ":" + port
	s.Db = db
}

func (s *Server) Start(){
	//Method to start server
	log.Println("Starting a server on port",s.port)
	r := router.NewRouter()
	r.Init()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET","PUT","PATCH","POST","DELETE","OPTIONS"}),
		handlers.AllowedHeaders([]string{"content-Type","Origin","Cache-Control","X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler =handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler: handler,
		Addr: "0.0.0.0" + s.port,
		WriteTimeout: 15 + time.Second,
		ReadTimeout: 15 * time.Second,
	}
	// http.ListenAndServe(s.port,r.Router)
	log.Fatal(newServer.ListenAndServe())
}