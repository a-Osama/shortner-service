package shotner

// Redirect Model Containing annotations for JSON, MongoDB, msgPack and Validator
type Redirect struct {
	Code      string `json:"code" bson:"code" msgpack:"code"`
	URL       string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt" msgpack:"createdAt"`
}
