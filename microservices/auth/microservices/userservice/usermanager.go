package userservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	url = "user-manager"
)

// UserInfo вызов функции userinfo в user-manager
func GetUserInfo(userName string) (UserInfo, error) {

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://%s/userinfo/%s", url, userName),
		nil,
	)

	cl := &http.Client{}

	resp, err := cl.Do(req)
	if err != nil {
		log.Println(err)
		return UserInfo{}, err
	}
	if resp.StatusCode != 200 {
		log.Println("Status code != 200, but ", resp.StatusCode)
		return UserInfo{}, errors.New("status code != 200, but " + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()

	info := UserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Println(err)
		return UserInfo{}, err
	}

	return info, nil
}
