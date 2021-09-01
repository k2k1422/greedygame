package utils

import (
	"auctioner/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func GetMaxValueBidder(bidderList []models.RegisterBidder) (models.ResponseBidder, error) {
	var (
		mu           = &sync.Mutex{}
		bidValueList = make([]models.ResponseBidder, 0)
	)

	var wg sync.WaitGroup
	var i int

	for i = 0; i < len(bidderList); i++ {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			bidValue, err := SendRegisterRequest(url)
			if err == nil {
				mu.Lock()
				bidValueList = append(bidValueList, bidValue)
				mu.Unlock()
			}

		}(bidderList[i].Url)

	}
	wg.Wait()
	if len(bidValueList) == 0 {
		return models.ResponseBidder{}, nil
	} else {
		maxBidder := bidValueList[0]
		for i = 1; i < len(bidValueList); i++ {
			if maxBidder.BidValue < bidValueList[i].BidValue {
				maxBidder = bidValueList[i]
			}
		}
		return maxBidder, nil
	}
}

func SendRegisterRequest(url string) (models.ResponseBidder, error) {
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return models.ResponseBidder{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return models.ResponseBidder{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return models.ResponseBidder{}, err
	}
	var result models.ResponseTemplate
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return models.ResponseBidder{}, err
	}
	// fmt.Print(result["response_code"])
	if result.ResponseCode != "111" {
		return models.ResponseBidder{}, errors.New("failed to register")
	}
	return result.Data, nil
}
