package main

import (
	"encoding/json"
	"fmt"
)

func main() {


	json_context := `
	{
	 "company": "网易",
	 "isok": true,
	 "price": 360.36,
	 "subject": [
	  "go",
	  "python",
	  "java"
	 ]
	}

	`

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(json_context),&m)
	if err != nil{
		return
	}
	fmt.Println(m)
	fmt.Println(m["company"])

}
