package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
}
type Employees struct {
	Employee []Employee `json:"emplyees"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("employee.json")
	// Handle error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened employee.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//initialize struct
	var employees Employee
	// jsonFile's content into 'employees' which we defined above
	err = json.Unmarshal(byteValue, &employees)
	if err != nil {
		fmt.Println(err)
	}
	//Create CSV File
	csvFile, err := os.Create("./employee.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, emp := range employees {
		var row []string
		row = append(row, emp.ID)
		row = append(row, emp.EmployeeName)
		row = append(row, emp.EmployeeSalary)
		row = append(row, emp.EmployeeAge)
		writer.Write(row)
	}
	// remember to flush!
	writer.Flush()
}
