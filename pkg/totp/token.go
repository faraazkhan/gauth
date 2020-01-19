package totp

import (
	"os"
	"path/filepath"
	"time"

	"github.com/faraazkhan/gauth/pkg/gauth"
	"github.com/pquerna/otp/totp"
)

// Token returns a 6 digit token for google authenticator
func Token(name, configFile string) (token string, err error) {
	if configFile == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		configFile = filepath.Join(homeDir, ".gauth")
	}

	secret, err := gauth.GetSecret(name, configFile)
	if err != nil {
		return "", err
	}

	n := time.Now().UTC()

	code, err := totp.GenerateCode(secret, n)

	return code, err
}
