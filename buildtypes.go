package teamcity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAllBuildConfigurations() (BuildConfigurations, error) {
	buildConfigs := BuildConfigurations{}
	chData := DataFlow{
		Response: make(chan *http.Response, 1),
	}

	url := fmt.Sprint(c.URL, "/app/rest/buildTypes")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return buildConfigs, err
	}

	chData.Request = req
	c.Flow <- chData

	for res := range chData.Response {
		body, err := processResponse(res)
		defer res.Body.Close()
		if err != nil {
			return buildConfigs, err
		}
		if err := json.Unmarshal(body, &buildConfigs); err != nil {
			return buildConfigs, err
		}
	}
	return buildConfigs, nil
}
