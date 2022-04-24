package main
import(
	"github.com/Alejandraarrieta/order-product-api/internal/database"
	"github.com/Alejandraarrieta/order-product-api/internal/logs"
	reviews "github.com/Alejandraarrieta/order-product-api/reviews/web"

)

func main(){
	_= logs.InitLogger()

	mongoClient := database.NewMongoClient("localhost")

	reviewHandler := reviews.NewReviewHandler(mongoClient)

	mux:= Routes(handler, reviewHandler)
	server := NewServer(mux)
	server.Run()
}