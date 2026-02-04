package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DecodeJSON(r *http.Request, payload any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(payload)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return fmt.Errorf("malformed JSON at position %d", syntaxErr.Offset)

		case errors.As(err, &unmarshalTypeErr):
			return fmt.Errorf(
				"invalid value for field '%s' (expected %s)",
				unmarshalTypeErr.Field,
				unmarshalTypeErr.Type,
			)

		case err == io.EOF:
			return errors.New("request body is empty")

		default:
			return err
		}
	}

	return nil
}
