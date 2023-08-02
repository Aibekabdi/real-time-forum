package ws

import (
	"encoding/json"
	"errors"
	"forum/internal/models"
)

func unmarshalEventBody(e *models.WSEvent, v interface{}) error {
	body, ok := e.Body.(map[string]interface{})
	if !ok {
		return errors.New("invalid event body")
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, &v)
}
