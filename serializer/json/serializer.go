package json

import (
	"encoding/json"

	shortner "github.com/a-Osama/urlshortner/shortner"
	"github.com/pkg/errors"
)

type Redirect struct{}

func (R *Redirect) Decode(input []byte) (*shortner.Redirect, error) {
	redirect := &shortner.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "Serializer.Redirect.Decode")
	}
	return redirect, nil
}
func (R *Redirect) Encode(input *Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "Serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
