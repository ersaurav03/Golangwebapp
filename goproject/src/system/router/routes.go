package router

import(
	HomeHandler "JSMPJ/goproject/src/controllers/home"
	"JSMPJ/goproject/pkg/types/routes"
	"net/http"
)
func Middleware(next http.Handler) http.Handler{
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	next.ServeHTTP(w,r)
})
}
func GetRoutes()  routes.Routes{

   return routes.Routes{
	   routes.Route{"Home","GET","/",HomeHandler.Index},
	//    routes.Route{"AuthStore","POST","/auth/login",AuthHandler.Login},
	//    routes.Route{"AuthCheck","GET","/auth/check",AuthHandler.Check},
   }
}