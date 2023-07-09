package main
import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-bookstore-mysql/pkg/routes"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookstoreBooks(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}