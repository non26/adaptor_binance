package helper

import (
	serviceconfig "adaptor/config"
	"errors"
)

func SelectAccount(accountId string, secrets *serviceconfig.Secrets) (apikey string, secretkey string, err error) {
	switch accountId {
	case "1":
		apikey = secrets.Account1.ApiKey
		secretkey = secrets.Account1.ApiSecret
	case "2":
		apikey = secrets.Account2.ApiKey
		secretkey = secrets.Account2.ApiSecret
	default:
		return "", "", errors.New("account not found")
	}
	return apikey, secretkey, nil
}
