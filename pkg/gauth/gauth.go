package gauth

import "github.com/go-ini/ini"

// GetSecret returns a secret and an error
func GetSecret(name, configFile string) (secret string, err error) {
	conf, err := ini.Load(configFile)
	if err != nil {
		return "", err
	}

	secret = conf.Section(name).Key("secret").String()

	return secret, nil
}
