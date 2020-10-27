package hashids

import (
	"fmt"
	"github.com/speps/go-hashids"
	"github.com/spf13/viper"
)

func Encode(data []int) (string, error) {
	if viper.GetString("salt.key") == ""{
		return "", fmt.Errorf("salt key is not initiated")
	}
	hdata := hashids.NewData()
	hdata.MinLength = 15
	hdata.Salt = viper.GetString("salt.key")

	hid, _ := hashids.NewWithData(hdata)

	hash, err := hid.Encode(data)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func Decode(data string) (int, error) {
	if viper.GetString("salt.key") == ""{
		return 0, fmt.Errorf("salt key is not initiated")
	}
	var d []int
	hdata := hashids.NewData()
	hdata.MinLength = 15
	hdata.Salt = viper.GetString("salt.key")

	hid, _ := hashids.NewWithData(hdata)
	d, err := hid.DecodeWithError(data)

	if err != nil {
		return 0, err
	}

	return d[0], nil
}
