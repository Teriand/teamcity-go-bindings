package teamcity

import (
	"net/http"

	"github.com/sethgrid/pester"
)

type DataFlow struct {
	Request  *http.Request
	Response chan *http.Response
}

type Client struct {
	HTTPClient *pester.Client
	URL        string
	Username   string
	Password   string
	Flow       chan DataFlow
	semaphore  chan bool
}

type BuildDetails struct {
	Count    int `json:"count"`
	Property []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"property"`
}

type BuildID int
type AgentID int
type PoolID int

type BuildConfigurations struct {
	Count      int         `json:"count"`
	Href       string      `json:"href"`
	BuildTypes []BuildType `json:"buildType"`
}

type BuildType struct {
	ID          BuildTypeID `json:"id"`
	Name        string      `json:"name"`
	ProjectName string      `json:"projectName"`
	ProjectID   string      `json:"projectId"`
	Href        string      `json:"href"`
	WebURL      string      `json:"webUrl"`
	Description string      `json:"description,omitempty"`
	Paused      bool        `json:"paused,omitempty"`
}
type Agent struct {
    ID  AgentID  `json:"id"`
    Name string `json:"name"`
    Pool Pool `json:"Pool"`
}
type Pool struct {
    ID  PoolID  `json:"id"`
    Name string `json:"name"`
}

type BuildTypeID string

type Branch struct {
	Name    string `json:"name"`
	Default bool   `json:"default,omitempty"`
}

type Branches struct {
	Count    int      `json:"count"`
	Branches []Branch `json:"branch"`
}

type Build struct {
	ID          BuildID     `json:"id"`
	BuildTypeID BuildTypeID `json:"buildTypeId"`
	Number      string      `json:"number"`
	Status      string      `json:"status"`
	State       string      `json:"state"`
	Running     string      `json:"running"`
	BranchName  string      `json:"branchName"`
	Href        string      `json:"href"`
	WebURL      string      `json:"webUrl"`
	QueuedDate  string      `json:"queuedDate"`
	StartDate   string      `json:"startDate"`
	FinishDate  string      `json:"finishDate"`
	Agent       Agent       `json:"agent"`
}

type Builds struct {
	Count    int     `json:"count"`
	Href     string  `json:"href"`
	NextHref string  `json:"nextHref"`
	PrevHref string  `json:"prevHref"`
	Builds   []Build `json:"build"`
}

type BuildLocator struct {
	BuildType BuildTypeID `yaml:"build_type"`
	Branch    string      `yaml:"branch"`
	Status    string      `yaml:"status"`
	Running   string      `yaml:"running"`
	SinceDate  string      `yaml:"sincedate"`
	QueuedDate  string      `yaml:"queueddate"`
	FinishDate  string      `yaml:"finishdate"`
	StartDate   string      `json:"startdate"`
	Canceled  string      `yaml:"canceled"`
	State  string      `yaml:"state"`
	Count     int

}

type BuildStatistics struct {
	Count    int `json:"count"`
	Property []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"property"`
	UsedFilter BuildLocator
}
