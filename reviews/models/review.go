package models
import(
	"time"
)
const maxLengthInComments = 400

// Review represent an anon review from some website
type Review struct{
	Id             int 64
	ProductName    string 
	Port           string 
	Price          string
	TypeCoin       string 
	TypeBusiness   string 
	Observations   string
	Status         string 
	DeliveryPeriod string 
	Quality        string 
	WayPay         string 
	IsPort         string
	Expiration     time.Time
	CreatedAt      time.Time 
}

// CreateReviewCMD command to create a new review
type CreateReviewCMD struct{
	ProductName    string `json:"productname,omitempty"`
	Port           string `json:"port,omitempty"`
	Price          string `json:"price,omitempty"`
	TypeCoin       string `json:"typecoin,omitempty"`
	TypeBusiness   string `json:"typebusiness,omitempty"`
	Observations   string `json:"observations,omitempty"`
	Status         string `json:"status,omitempty"`
	DeliveryPeriod string `json:"deliveryperiod,omitempty"`
	Quality        string `json:"quality,omitempty"`
	WayPay         string `json:"waypay,omitempty"`
	IsPort         string `json:"isport,omitempty"`
	Expiration     time.Time `json:"expiration,omitempty"`
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