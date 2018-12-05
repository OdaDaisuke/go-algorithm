package readjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type UserJsonModel struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func Start() {
	var usersFromJson []UserJsonModel
	rawData, err := ioutil.ReadFile("./read_json/users.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(rawData, &usersFromJson)
	fmt.Println("START PRINT JSON")
	for _, v := range usersFromJson {
		fmt.Printf("Name: %s, Age: %2d\n", v.Name, v.Age)
	}
}
