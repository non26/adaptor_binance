package helper

import (
	"errors"
	"strings"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func ValidateSide(side string) error {
	side = strings.ToUpper(side)
	if side != bnconstant.BUY && side != bnconstant.SELL {
		return errors.New("invalid side")
	}
	return nil
}
