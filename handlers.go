package teamcity

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/fatih/structs"
)

func processResponse(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode == 404 {
		return []byte{}, errors.New(string(body))
	}
	if err != nil {
		return []byte{}, errors.New("Failed to read body")
	}
	return body, nil
}

func (c *Client) processDataFlow() {
	for data := range c.Flow {
		c.semaphore <- true
		go func(d DataFlow) {
			defer func() { <-c.semaphore }()
			d.Request.Header.Add("Accept", "application/json")
			d.Request.Header.Add("Connection", "close")
			d.Request.Close = true
			res, err := c.HTTPClient.Do(d.Request)
			if err != nil {
				log.Println("processDataFlow %s",err)
				close(d.Response)
				return
			}
			if res.StatusCode == 401 {
				d.Request.SetBasicAuth(c.Username, c.Password)
				res, err = c.HTTPClient.Do(d.Request)
				if err != nil {
					log.Println("processDataFlow %s",err)
					close(d.Response)
					return
				}
			}
			d.Response <- res
			close(d.Response)
		}(data)
	}
}

func convertLocatorToString(str BuildLocator) string {
	resMap := map[string]string{}
	for k, v := range structs.Map(str) {
		if s, ok := v.(string); ok && s != "" && s != "<default>" {
			fieldName := []rune(k)
			fieldName[0] = unicode.ToLower(fieldName[0])
			resMap[string(fieldName)] = fmt.Sprint("(", v, ")")
		} else if i, ok := v.(int); ok && i != 0 {
			fieldName := []rune(k)
			fieldName[0] = unicode.ToLower(fieldName[0])
			resMap[string(fieldName)] = fmt.Sprint("(", i, ")")
		} else if bt, ok := v.(BuildTypeID); ok && string(bt) != "" {
			fieldName := []rune(k)
			fieldName[0] = unicode.ToLower(fieldName[0])
			resMap[string(fieldName)] = fmt.Sprint("(", bt, ")")
		}
	}

	resString := ""
	for k, v := range resMap {
		resString += fmt.Sprint(k, ":", v, ",")
	}

	return strings.TrimRight(resString, ",")
}
