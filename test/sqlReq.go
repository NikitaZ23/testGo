package main

import (
	"database/sql"

  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

func updatePassport(typeP, numberP, newT, newP string){
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
	if err != nil{
		fmt.Println("Не подключились")
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("Select id from test.passport where Type = '%s' and Number = '%s'", typeP, numberP))
	if err != nil{
		fmt.Println(err)
	}

	for res.Next(){
		var id int
		err = res.Scan(&id)
		if err != nil{
			fmt.Println(err)
		}

		res, err = db.Query(fmt.Sprintf("UPDATE test.passport set Type = '%s', Number = '%s'  where id = '%d'", newT, newP, id))
		if err != nil{
			fmt.Println(err)
		}
	}
}

func addPassport(typeP, numberP string)int{
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("Select id from test.passport where Type = '%s' and Number = '%s'", typeP, numberP))
  if err != nil{
    fmt.Println(err)
  }

  for res.Next(){
    var id int
    err = res.Scan(&id)
    if err != nil{
      fmt.Println(err)
    }
    _, err = db.Query(fmt.Sprintf("delete from test.passport where id = '%d'", id))
    if err != nil{
      fmt.Println(err)
    }
  }

  insert, err := db.Query(fmt.Sprintf("Insert into test.passport (Type, Number) VALUES('%s','%s');", typeP, numberP))
  if err != nil{
    fmt.Println(err)
  }
  defer insert.Close()

  res, err = db.Query(fmt.Sprintf("Select id from test.passport where Type = '%s' and Number = '%s'", typeP, numberP))
  if err != nil{
    fmt.Println(err)
  }

  var id int
  for res.Next(){
    err = res.Scan(&id)
    if err != nil{
      fmt.Println(err)
    }
  }

  return id
}

func getCompany(companyName string) int{
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("Select id from test.company where Name = '%s'", companyName))
  if err != nil{
    fmt.Println(err)
  }

  var id int
  pr := false
  for res.Next(){
    pr = true
    err = res.Scan(&id)
    if err != nil{
      fmt.Println(err)
    }
  }

  if !pr{
    insert, err := db.Query(fmt.Sprintf("Insert into test.company (Name) VALUES('%s');", companyName))
    if err != nil{
      fmt.Println(err)
    }
    defer insert.Close()

    res, err := db.Query(fmt.Sprintf("Select id from test.company where Name = '%s'", companyName))
    if err != nil{
      fmt.Println(err)
    }

    for res.Next(){
      err = res.Scan(&id)
      if err != nil{
        fmt.Println(err)
      }
    }
  }
  return id
}

func getDepartment(departmentName string, depPhone string,  idC int) int{
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("select id from department where CompanyId = '%d' and Name = '%s';", idC, departmentName))
  if err != nil{
    fmt.Println(err)
  }

  var id int
  pr := false
  for res.Next(){
    pr = true
    err = res.Scan(&id)
    if err != nil{
      fmt.Println(err)
    }
  }

  if !pr{
    insert, err := db.Query(fmt.Sprintf("Insert into test.department (Name, Phone, CompanyId)"+
    " values('%s', '%s', '%d');", departmentName, depPhone, idC))
    if err != nil{
      fmt.Println(err)
    }
    defer insert.Close()

    res, err := db.Query(fmt.Sprintf("select id from department where CompanyId = '%d' and Name = '%s';", idC, departmentName))
    if err != nil{
      fmt.Println(err)
    }

    for res.Next(){
      err = res.Scan(&id)
      if err != nil{
        fmt.Println(err)
      }
    }
  }
  return id
}

func addEmployee(nameS , surnameS, phoneS string, idC, idP, idD int) int {
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("select id from employees where Name = '%s' and Surname = '%s' and " +
    "Phone = '%s' and Passport = '%d';", nameS, surnameS, phoneS, idP))
  if err != nil{
    fmt.Println(err)
  }

  var id int
  pr := false
  for res.Next(){
    pr = true
    err = res.Scan(&id)
    if err != nil{
      fmt.Println(err)
    }
  }

  if !pr{
    insert, err := db.Query(fmt.Sprintf("Insert into test.employees (Name, Surname, Phone, CompanyId, Passport, Department)"+
    " values('%s', '%s', '%s', '%d', '%d', '%d');", nameS, surnameS, phoneS, idC, idP, idD))
    if err != nil{
      fmt.Println(err)
    }
    defer insert.Close()

    res, err := db.Query(fmt.Sprintf("select id from employees where Name = '%s' and Surname = '%s' and " +
      "Phone = '%s' and CompanyId = '%d' and Passport = '%d' and Department = '%d';", nameS, surnameS, phoneS, idC, idP, idD))
    if err != nil{
      fmt.Println(err)
    }

    for res.Next(){
      err = res.Scan(&id)
      if err != nil{
        fmt.Println(err)
      }
    }
  }

  return id
}

func deleteS(nameS int) bool{

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

	_, err = db.Query(fmt.Sprintf("delete from test.employees where id = '%d'", nameS))
	if err != nil{
		fmt.Println(err)
	}

  res, err := db.Query(fmt.Sprintf("select * from employees where id = '%d';", nameS))
  if err != nil{
    fmt.Println(err)
  }

	pr := true
	for res.Next(){
		pr = false
	}

	return pr
}

func getDepComp(idC int, nameD string) int{
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
	if err != nil{
		fmt.Println("Не подключились")
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("select id from department where Name = '%s' and CompanyId = '%d';", nameD, idC))
	if err != nil{
		fmt.Println(err)
	}

	var id int
	for res.Next(){
		err = res.Scan(&id)

		if err != nil{
			fmt.Println(err)
		}
	}

	return id
}
