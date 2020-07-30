//shows how we can nest fields and encode to json
package main

import(
	"fmt"
	"encoding/json"
)

type App struct {
    Id string `json:"id"`
}

type Org struct {
    Name string `json:"name"`
}

type AppWithOrg struct {
    App
    Org
}

func main(){
	
	data := []byte(`
    {
        "id": "k34rAT4",
        "name": "My Awesome Org"
    }
`)

var appWithOrg AppWithOrg
err := json.Unmarshal(data, &appWithOrg)
if err != nil {
	fmt.Println(err)
}
fmt.Println(appWithOrg)

// app := appWithOrg.App
// org := appWithOrg.Org

// // AND/OR

// appId := appWithOrg.Id
// orgName := appWithOrg.Name
}

