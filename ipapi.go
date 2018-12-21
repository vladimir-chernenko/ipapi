package ipapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type IpDetails struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Capital     string  `json:"capital"`
	Currency    string  `json:"currency"`
	PhonePrefix string  `json:"phonePrefix"`
}

func GetIpDetails(ip string) (IpDetails, error) {
	var ipd IpDetails

	baseURL := "http://ip-api.com/json/"
	c := http.Client{Timeout: time.Second * 2}

	r, err := c.Get(baseURL + ip)
	if err != nil {
		return ipd, err
	}

	err = json.NewDecoder(r.Body).Decode(&ipd)

	if err == nil {
		ipd.Capital = capitals[ipd.CountryCode]
		ipd.Currency = currencies[ipd.CountryCode]
		ipd.PhonePrefix = phonePrefixes[ipd.CountryCode]
	}

	return ipd, err
}
