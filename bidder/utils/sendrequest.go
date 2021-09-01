package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendRegisterRequest(url string, host string, port string) error {
	method := "POST"
	send_url := fmt.Sprintf(`http://%s:%s/bid`, host, port)
	payload := strings.NewReader(fmt.Sprintf(`
	{"url": "%s"}
  	`, send_url))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Print(result["response_code"])
	if result["response_code"] != "111" {
		return errors.New("failed to register")
	}
	return nil
}
