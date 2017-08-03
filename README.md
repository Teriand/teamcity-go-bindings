# TeamCity API bindings
Teamcity v10 API client for golang. Only basic methods have been implemented so far. PRs are welcome.

## Installation
```
$ go get -u github.com/Guidewire/teamcity-go-bindings
```

## Usage example
```go
package main

import (
	"fmt"
	"log"

	tc "github.com/guidewire/teamcity-go-bindings"
)

func main() {
	// put your login and password here
	client := tc.New("https://teamcity.com", "login", "password", 10)

	// Get all available build types
	bt, err := client.GetAllBuildConfigurations()
	if err != nil {
		log.Println(err)
	}

	// Get branches list for the first 3 build types found
	for i := 0; i < 3; i++ {
		b, err := client.GetAllBranches(bt.BuildTypes[i].ID)
		if err != nil {
			log.Println(err)
		}
		for v := range b.Branches {
			fmt.Printf("Build type: %s, branch: %s\n", bt.BuildTypes[i].ID, b.Branches[v].Name)
		}
	}
}
```
