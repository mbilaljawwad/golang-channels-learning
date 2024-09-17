package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func GetUsersDictionary() (map[string]User, error) {
	jsonFile, err := os.Open("../internal/dictionary/users.json")
	if err != nil {
		fmt.Println("Unable to open file. ", err)
		return nil, err
	}
	defer jsonFile.Close()

	dec := json.NewDecoder(jsonFile)

	var data []User
	err = dec.Decode(&data)
	if err != nil {
		fmt.Println("Unable to read data. ", err)
		return nil, err
	}

	var dict = make(map[string]User)
	for _, d := range data {
		dict[d.Name] = d
	}

	return dict, nil
}
