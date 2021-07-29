package shotner

// Redirect Model Containing annotations for JSON, MongoDB, msgPack and Validator
type Redirect struct {
	Code      string `json:"code" bson:"code"`
	URL       string `json:"url" bson:"url" validate:"empty=false & format=url"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
}
