package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

var validate = validator.New()

const MaxBodySize = 1 << 20

func DecodeAndValidate(r http.Request, v any) error {
	d := json.NewDecoder(io.LimitReader(r.Body, MaxBodySize))
	d.DisallowUnknownFields()

	if err := d.Decode(v); err != nil {
		var se *json.SyntaxError
		switch {
		case errors.As(err, &se):
			return fmt.Errorf("bad JSON syntax at byte %d", se.Offset)

		case errors.Is(err, io.EOF):
			return fmt.Errorf("empty request body")
		default:
			return fmt.Errorf("invalid JSON: %w", err)
		}
	}

	if d.More() {
		return errors.New("body must contain a single JSON object")
	}

	if err := validate.Struct(v); err != nil {
		if verrs, ok := err.(validator.ValidationErrors); ok {
			fe := verrs[0]
			return fmt.Errorf("field %s %s", fe.Field(), validationMsg(fe))
		}
		return err
	}
	return nil
}

func validationMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is a required field"
	case "uuid":
		return "must be a valid UUID"
	case "email":
		return "must be a valid email address"
	case "oneof":
		return fmt.Sprintf("must be a one of %s", fe.Param())
	default:
		return "is valid"
	}
}
