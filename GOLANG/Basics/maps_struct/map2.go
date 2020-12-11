package main

import "fmt"

func main() {

	Employee := map[string]map[string]string{
		"Febin": map[string]string{
			"Position":   "Trainee",
			"EmployeeId": "20",
		},
		"Favas": map[string]string{
			"Position":   "Developer",
			"EmployeeId": "15",
		},
	}

	if temp, name := Employee["Febin"]; name {
		fmt.Println(temp["Position"], temp["EmployeeId"])
	}
}
