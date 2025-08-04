package helper

import (
	"strconv"

	bnutils "github.com/non26/tradepkg/pkg/bn/utils"
)

func GetTimestamp() string {
	timestamp := strconv.FormatInt(bnutils.GetBinanceTimestamp(), 10)
	return timestamp
}
