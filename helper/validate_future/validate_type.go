package helper

import (
	"errors"
	"strings"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func ValidateType(orderType string) error {
	orderType = strings.ToUpper(orderType)
	if orderType != bnconstant.LIMIT && orderType != bnconstant.MARKET {
		return errors.New("invalid order type")
	}
	return nil
}
