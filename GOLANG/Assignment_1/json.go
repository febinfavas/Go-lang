package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var studentName = [5]string{"A", "B", "C", "D", "E"}
	var studentClass = [5]string{"I", "II", "III", "IV", "V"}
	var data = map[string]interface{}{"studentName": studentName, "studnetClass": studentClass}
	jsonData, _ := json.Marshal(map[string]interface{}{"Data": data, "status": "200"})
	fmt.Println(string(jsonData))
}
