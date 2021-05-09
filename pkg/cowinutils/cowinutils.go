package cowinutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"t2findmyvaccinebot/pkg/common"
)

func GetStates() (common.StateList, error) {

	var stateList common.StateList
	url := "https://cdn-api.co-vin.in/api/v2/admin/location/states"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return stateList, err

	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return stateList, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	errUnM := json.Unmarshal(body, &stateList)
	if errUnM != nil {
		log.Fatalln(errUnM)
		return stateList, errUnM
	}

	return stateList, nil
}

func GetDistricts(stateID string) (common.DistrictList, error) {

	var distList common.DistrictList

	url := "https://cdn-api.co-vin.in/api/v2/admin/location/districts/" + stateID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return distList, err

	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return distList, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
		return distList, err
	}

	errUnM := json.Unmarshal(body, &distList)
	if errUnM != nil {
		log.Fatalln(errUnM)
		return distList, errUnM
	}

	return distList, nil
}

func GetSessionByDist(ID, date string) (common.SessionList, error) {

	var sessionList common.SessionList

	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByDistrict?district_id=" + ID + "&date=" + date
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return sessionList, err

	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return sessionList, err
	}
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
			return sessionList, err
		}
		errUnM := json.Unmarshal(body, &sessionList)
		if errUnM != nil {
			log.Fatalln(errUnM)
			return sessionList, errUnM
		}
	}

	return sessionList, nil

}

func GetSessionByPin(pinID, date string) (common.SessionList, error) {

	var sessionList common.SessionList
	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin?pincode=" + pinID + "&date=" + date
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return sessionList, err

	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return sessionList, err
	}
	if resp.StatusCode == http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
			return sessionList, err
		}

		errUnM := json.Unmarshal(body, &sessionList)
		if errUnM != nil {
			log.Fatalln(errUnM)
			return sessionList, errUnM
		}
	}
	return sessionList, nil

}
