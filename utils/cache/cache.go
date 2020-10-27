package cache

import (
	"fmt"
)

var ()

func CheckAndRetrieve(model string, uid int64, name string, unique string) ([]byte, error) {
	cacheKey := fmt.Sprintf("%s:%s:%v:%s",
		model,
		name,
		uid,
		unique,
	)

	bytes, err := GetCache(cacheKey)

	if err != nil {
		return nil, fmt.Errorf("error while retrieving cache : %v", err.Error())
	}

	return bytes, nil
}

func Query(model string, uid int64, name string, url string, data interface{}, time int) error {
	//hashQuery := md5.New()
	//hashQuery.Write([]byte(query))
	//
	//key := hex.EncodeToString(hashQuery.Sum(nil))

	cacheKey := fmt.Sprintf("%s:%s:%v:%s",
		model,
		name,
		uid,
		url,
	)

	if err := SetCache(cacheKey, data, time); err != nil {
		return err
	}

	return nil
}
