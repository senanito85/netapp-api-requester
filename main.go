package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Endpoint string `json:"endpoint"`
}

const configFile = "config.json"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	rawCfg, err := ioutil.ReadFile(
		path.Join(wd, configFile),
	)
	if err != nil {
		log.Fatalln(err)
	}

	c := &config{}
	err = json.Unmarshal(rawCfg, c)
	if err != nil {
		log.Fatalln(err)
	}

	if len(c.Username) < 1 {
		log.Fatalln("username must not be empty")
	}

	if len(c.Password) < 1 {
		log.Fatalln("password must not be empty")
	}

	if len(c.Endpoint) < 1 {
		log.Fatalln("endpoint must not be empty")
	}

	req, err := http.NewRequest("GET", c.Endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.SetBasicAuth(c.Username, c.Password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	parsedResp, err := parseJSON(body)
	if err != nil {
		log.Fatalln(err)
	}

	output := convert(parsedResp)
	write(output)
}
