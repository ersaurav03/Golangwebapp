package router
import(
	"JSMPJ/goproject/pkg/types/routes"
	StatusHandler "JSMPJ/goproject/src/controllers/v1/status"
	"net/http"
	"log"
)
func Middleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		token:=r.Header.Get("X-App-Token")
	    if len(token) < 1 {
			http.Error(w,"Not Atuhorized",http.StatusUnauthorized)
			return
		}
		log.Println("Inide v1 middle ware")
		next.ServeHTTP(w,r)
	})
}
func GetRoutes()(SubRoute map[string]routes.SubRoutePackage){
SubRoute = map[string]routes.SubRoutePackage{
	"/v1": routes.SubRoutePackage{
		Routes: routes.Routes{
			routes.Route{"Home","GET","/",StatusHandler.Index},
		},
		Middleware:Middleware,
	},
}	
return
}