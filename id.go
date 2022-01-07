package pokeapi

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"strings"
)

type ID string

func (id ID) String() string {
	return string(id)
}

func (id ID) Validate() *errors.RestErr {
	if strings.TrimSpace(id.String()) == "" {
		return errors.BadRequestErrorf(errors.MissingMandatoryParamErrMsg, "id or name")
	}
	return nil
}
