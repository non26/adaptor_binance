package helper

import (
	"errors"
	"strings"
)

func ValidateSymbol(symbol string) error {
	symbol = strings.ToUpper(symbol)
	if strings.HasSuffix(symbol, "USDT") {
		return errors.New("invalid symbol")
	}
	return nil
}
