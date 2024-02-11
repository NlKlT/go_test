package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	CompanyId int    `json:"company_id"`
	Passport  struct {
		Type   string `json:"type"`
		Number string `json:"number"`
	} `json:"passport"`
	Department struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"department"`
}

var employees []Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyId := params["company_id"]
	var filteredEmployees []Employee
	for _, employee := range employees {
		if employee.CompanyId == toInt(companyId) {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}
	json.NewEncoder(w).Encode(filteredEmployees)
}

func getEmployeesByDepartment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyId := params["company_id"]
	departmentName := params["department_name"]
	var filteredEmployees []Employee
	for _, employee := range employees {
		if employee.CompanyId == toInt(companyId) && employee.Department.Name == departmentName {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}
	json.NewEncoder(w).Encode(filteredEmployees)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	employee.Id = len(employees) + 1
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i, employee := range employees {
		if employee.Id == toInt(id) {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func updateEmployeeDepartment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var updatedDepartment struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
	json.NewDecoder(r.Body).Decode(&updatedDepartment)
	for i, employee := range employees {
		if employee.Id == toInt(id) {
			if updatedDepartment.Name != "" {
				employee.Department.Name = updatedDepartment.Name
			}
			if updatedDepartment.Phone != "" {
				employee.Department.Phone = updatedDepartment.Phone
			}
			employees[i] = employee
			break
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var updatedEmployee Employee
	json.NewDecoder(r.Body).Decode(&updatedEmployee)
	for i, employee := range employees {
		if employee.Id == toInt(id) {
			if updatedEmployee.Name != "" {
				employee.Name = updatedEmployee.Name
			}
			if updatedEmployee.Surname != "" {
				employee.Surname = updatedEmployee.Surname
			}
			if updatedEmployee.Phone != "" {
				employee.Phone = updatedEmployee.Phone
			}
			if updatedEmployee.CompanyId != 0 {
				employee.CompanyId = updatedEmployee.CompanyId
			}
			if updatedEmployee.Passport.Type != "" {
				employee.Passport.Type = updatedEmployee.Passport.Type
			}
			if updatedEmployee.Passport.Number != "" {
				employee.Passport.Number = updatedEmployee.Passport.Number
			}
			if updatedEmployee.Department.Name != "" {
				employee.Department.Name = updatedEmployee.Department.Name
			}
			if updatedEmployee.Department.Phone != "" {
				employee.Department.Phone = updatedEmployee.Department.Phone
			}
			employees[i] = employee
			break
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees/company/{company_id}", getEmployees).Methods("GET")
	router.HandleFunc("/employees/department/{company_id}/{department_name}", getEmployeesByDepartment).Methods("GET")
	router.HandleFunc("/employees", createEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}/department", updateEmployeeDepartment).Methods("PUT")
	http.ListenAndServe(":8080", router)
}

func toInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return result
}
