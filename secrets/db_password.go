package secrets

import (
	"os"
)

var dBPassword *string

func readDBPassword() {
	pwd, err := os.ReadFile("/run/secrets/db_password")
	if err != nil {
		panic(err)
	}
	dBPassword = new(string)
	*dBPassword = string(pwd)
}

func GetDBPassword() string {
	if dBPassword == nil {
		readDBPassword()
	}
	return *dBPassword
}
