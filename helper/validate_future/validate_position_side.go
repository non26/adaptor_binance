package helper

import (
	"errors"
	"strings"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func ValidatePositionSide(positionSide string) error {
	positionSide = strings.ToUpper(positionSide)
	if positionSide != bnconstant.LONG && positionSide != bnconstant.SHORT {
		return errors.New("invalid position side")
	}

	return nil
}
