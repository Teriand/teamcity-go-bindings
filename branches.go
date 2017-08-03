package teamcity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAllBranches(bt BuildTypeID) (Branches, error) {
	branches := Branches{}
	chData := DataFlow{
		Response: make(chan *http.Response, 1),
	}

	url := fmt.Sprint(c.URL, "/app/rest/buildTypes/id:", bt, "/branches")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return branches, err
	}

	chData.Request = req
	c.Flow <- chData

	for res := range chData.Response {
		body, err := processResponse(res)
		if err != nil {
			return branches, err
		}
		if err := json.Unmarshal(body, &branches); err != nil {
			return branches, err
		}
	}

	for _, branch := range branches.Branches {
		branch.Name = strings.Replace(branch.Name, "/refs/heads/", "", -1)
	}
	return branches, nil
}
