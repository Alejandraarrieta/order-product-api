package models
import(
	"time"
)
const maxLengthInComments = 400

// Review represent an anon review from some website
type Review struct{
	Id int 64
	ProductName    string 
	Port           string 
	Price          string 
	Date           time.Time 
	TypeBusiness   string 
	Status         string 
	DeliveryPeriod string 
	Quality        string 
	WayPay         string 
	Expiration     string
	IsPort         string
}

// CreateReviewCMD command to create a new review
type CreateReviewCMD struct{
	ProductName    string `json:"productname,omitempty" bson:"productname,omitempty"`
	Port           string `json:"port,omitempty" bson:"portname,omitempty"`
	Price          string `json:"price,omitempty" bson:"price,omitempty"`
	TypeCoin       string `json:"typecoin,omitempty" bson:"typecoin,omitempty"`
	Date           string `json:"date,omitempty" bson:"date,omitempty"`
	TypeBusiness   string `json:"typebusiness,omitempty" bson:"typebusiness,omitempty"`
	Observations   string `json:"observations,omitempty" bson:"observations,omitempty"`
	Status         string `json:"status,omitempty" bson:"status,omitempty"`
	DeliveryPeriod string `json:"deliveryperiod,omitempty" bson:"deliveryperiod,omitempty"`
	Quality        string `json:"quality,omitempty" bson:"quality,omitempty"`
	WayPay         string `json:"waypay,omitempty" bson:"waypay,omitempty"`
	Expiration     string `json:"expiration,omitempty" bson:"expiration,omitempty"`
	IsPort         string `json:"isport,omitempty" bson:"isport,omitempty"`

}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 1 - 5")
	}

	if len(cmd.Observations) > maxLengthInComments {
		return errors.New("comment must be less than 400 chars")
	}
	return nil
}