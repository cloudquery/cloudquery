package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"golang.org/x/exp/maps"
)

type Service struct {
	Name string `json:"name"`
}

func Discover() ([]string, error) {
	resp, err := http.Get("https://discovery.googleapis.com/discovery/v1/apis")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var services struct {
		Items []Service `json:"items"`
	}
	err = json.Unmarshal(body, &services)
	if err != nil {
		return nil, err
	}

	servicesMap := make(map[string]bool)
	for _, service := range services.Items {
		servicesMap[service.Name+".googleapis.com"] = true
	}

	serviceList := maps.Keys(servicesMap)
	sort.Strings(serviceList)
	return serviceList, nil
}
