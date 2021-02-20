package ip_checker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

type ipResponse struct {
	IP                 string  `json:"ip"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	CountryName        string  `json:"country_name"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           string  `json:"latitude"`
	Longitude          string  `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  float64 `json:"country_population"`
	Message            string  `json:"message"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

type PositionChecker struct {
	ips   map[string]ipResponse
	ipsMU sync.RWMutex
}

func NewPosition() *PositionChecker {
	return &PositionChecker{
		ips: make(map[string]ipResponse, 0),
	}
}

func (p *PositionChecker) GetMe(ip string) (country, city string) {
	p.ipsMU.RLock()
	val, ok := p.ips[ip]
	p.ipsMU.RUnlock()
	if !ok {
		var err error
		val, err = p.getData(ip)
		if err != nil {
			return "undefined", "undefined"
		}
		p.ipsMU.Lock()
		p.ips[ip] = val
		p.ipsMU.Unlock()
	}
	return val.Country, val.City
}

func (p *PositionChecker) getData(ip string) (ipResponse, error) {
	var resp ipResponse
	println("https://ipapi.co/" + ip + "/json/")
	req, _ := http.NewRequest("GET", "https://ipapi.co/"+ip+"/json/", nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		println("err ask ip", err.Error())
		return ipResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return resp, errors.New("invalid code")
	}

	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &resp)
	return resp, nil
}
