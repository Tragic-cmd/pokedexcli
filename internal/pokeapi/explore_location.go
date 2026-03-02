package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// exploreLocation -
func (c *Client) ExploreLocation(area string) (RespArea, error) {
	url := baseURL + "/location-area/" + area + "/"

	if val, ok := c.cache.Get(url); ok {
		areaResp := RespArea{}
		err := json.Unmarshal(val, &areaResp)
		if err != nil {
			return RespArea{}, err
		}

		return areaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespArea{}, err
	}

	areaResp := RespArea{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespArea{}, err
	}

	c.cache.Add(url, dat)
	return areaResp, nil
}
