package validator

import "go-commons/errors"

type Validator interface {
	Validate() []errors.ApiErrorItem
}
