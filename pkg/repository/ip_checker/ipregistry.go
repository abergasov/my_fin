package ip_checker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type ipRegistryResponse struct {
	IP       string      `json:"ip"`
	Type     string      `json:"type"`
	Hostname interface{} `json:"hostname"`
	Carrier  struct {
		Name interface{} `json:"name"`
		Mcc  interface{} `json:"mcc"`
		Mnc  interface{} `json:"mnc"`
	} `json:"carrier"`
	Connection struct {
		Asn          int    `json:"asn"`
		Domain       string `json:"domain"`
		Organization string `json:"organization"`
		Route        string `json:"route"`
		Type         string `json:"type"`
	} `json:"connection"`
	Currency struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		NameNative   string `json:"name_native"`
		Plural       string `json:"plural"`
		PluralNative string `json:"plural_native"`
		Symbol       string `json:"symbol"`
		SymbolNative string `json:"symbol_native"`
		Format       struct {
			Negative struct {
				Prefix string `json:"prefix"`
				Suffix string `json:"suffix"`
			} `json:"negative"`
			Positive struct {
				Prefix string `json:"prefix"`
				Suffix string `json:"suffix"`
			} `json:"positive"`
		} `json:"format"`
	} `json:"currency"`
	Location struct {
		Continent struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"continent"`
		Country struct {
			Area              int      `json:"area"`
			Borders           []string `json:"borders"`
			CallingCode       string   `json:"calling_code"`
			Capital           string   `json:"capital"`
			Code              string   `json:"code"`
			Name              string   `json:"name"`
			Population        int      `json:"population"`
			PopulationDensity float64  `json:"population_density"`
			Flag              struct {
				Emoji        string `json:"emoji"`
				EmojiUnicode string `json:"emoji_unicode"`
				Emojitwo     string `json:"emojitwo"`
				Noto         string `json:"noto"`
				Twemoji      string `json:"twemoji"`
				Wikimedia    string `json:"wikimedia"`
			} `json:"flag"`
			Languages []struct {
				Code   string `json:"code"`
				Name   string `json:"name"`
				Native string `json:"native"`
			} `json:"languages"`
			Tld string `json:"tld"`
		} `json:"country"`
		Region struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"region"`
		City      string  `json:"city"`
		Postal    string  `json:"postal"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Language  struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"language"`
		InEu bool `json:"in_eu"`
	} `json:"location"`
	Security struct {
		IsBogon         bool `json:"is_bogon"`
		IsCloudProvider bool `json:"is_cloud_provider"`
		IsTor           bool `json:"is_tor"`
		IsTorExit       bool `json:"is_tor_exit"`
		IsProxy         bool `json:"is_proxy"`
		IsAnonymous     bool `json:"is_anonymous"`
		IsAbuser        bool `json:"is_abuser"`
		IsAttacker      bool `json:"is_attacker"`
		IsThreat        bool `json:"is_threat"`
	} `json:"security"`
	TimeZone struct {
		ID               string    `json:"id"`
		Abbreviation     string    `json:"abbreviation"`
		CurrentTime      time.Time `json:"current_time"`
		Name             string    `json:"name"`
		Offset           int       `json:"offset"`
		InDaylightSaving bool      `json:"in_daylight_saving"`
	} `json:"time_zone"`
}

type IPRegistry struct {
	apiKey string
	ips    map[string]ipRegistryResponse
	ipsMU  sync.RWMutex
}

func InitIPRegistry(api string) *IPRegistry {
	return &IPRegistry{
		apiKey: api,
		ips:    make(map[string]ipRegistryResponse, 0),
	}
}

func (p *IPRegistry) GetMe(ip string) (country, city string) {
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
	return val.Location.Country.Name, val.Location.City
}

func (p *IPRegistry) getData(ip string) (ipRegistryResponse, error) {
	var resp ipRegistryResponse
	req, _ := http.NewRequest("GET", "https://api.ipregistry.co/"+ip+"?key="+p.apiKey, nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		println("err ask ip", err.Error())
		return resp, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		println("invalid code", res.StatusCode, string(body))
		return resp, errors.New("invalid code")
	}

	_ = json.Unmarshal(body, &resp)
	return resp, nil
}
