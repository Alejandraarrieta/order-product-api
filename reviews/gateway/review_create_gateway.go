package gateway

import "github.com/Alejandraarrieta/order-product-api/reviews/models"

type ReviewGateway interface {
	AddReview(cmd models.CreateReviewCMD)(string, error)
}

type ReviewGtw struct{
	ReviewStorage
}

//func para devolver la interface
func NewReviewGateway(mongoClient *database.Mongo) ReviewGateway {
	return &ReviewGtw{&ReviewStg{mongoClient}}
}