package gateway

import "github.com/Alejandraarrieta/order-product-api/reviews/models"

type ReviewStorage interface {
	Add(cmd *models.CreateReviewCMD) (string, error)
}

type ReviewStg struct {
	*database.Mongo
}

func (s *ReviewStg) Add(cmd *models.CreateReviewCMD) (string, error) {
	coll := s.Client.Database("productDB").Collection("products")

	res, err := coll.InsertOne(context.Background(),
		bson.D{ //documento a insertar con bson ordenado
			{"productname", cmd.Productname},
			{"port", cmd.Port},
			{"price", cmd.Price},
			{"typecoin", cmd.TypeCoin},
			{"typebusiness", cmd.TypeBusiness},
			{"observations", cmd.Observations},
			{"status", cmd.Status},
			{"deliveryperiod", cmd.DeliveryPeriod},
			{"quality", cmd.Quality},
			{"waypay", cmd.WayPay},
			{"isport", cmd.IsPort},
			{"createdAt", time.Now()},
		})

	if err != nil {
		logs.Log().Error("cannot insert review")
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID)

	return id.String(), nil
}

