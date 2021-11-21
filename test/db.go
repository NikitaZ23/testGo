package main

import (
	"database/sql"

  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

func createDB()  {
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

	_,err = db.Exec("CREATE DATABASE test")
   if err != nil {
       fmt.Println(err)
   }else{
     fmt.Println("test created")
   }

   _,err = db.Exec("USE test")
   if err != nil {
       fmt.Println(err)
   }else{
     fmt.Println("USE test")
   }

  _, err = db.Query("CREATE TABLE `test`.`Company` ( "+
    "`id` INT NOT NULL AUTO_INCREMENT, "+
    "`Name` VARCHAR(45) NOT NULL,"+
    "PRIMARY KEY (`id`));")
  if err != nil{
    fmt.Println(err)
  }else{
    fmt.Println("Table Company created")
  }

  _, err = db.Query("CREATE TABLE `test`.`Department` ( "+
      "`id` INT NOT NULL AUTO_INCREMENT,"+
      "`Name` VARCHAR(45) NOT NULL,"+
      "`Phone` VARCHAR(45) NOT NULL,"+
      "`CompanyId` INT NOT NULL,"+
      "PRIMARY KEY (`id`),"+
      "KEY `dep_comp_1` (`CompanyId`),"+
      "CONSTRAINT `dep_comp_1` FOREIGN KEY (`CompanyId`)"+
      " REFERENCES `test`.`Company` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE);")
  if err != nil{
    fmt.Println(err)
  }else{
    fmt.Println("Table department created")
  }

  _, err = db.Query("CREATE TABLE `test`.`Passport` ("+
    "`id` INT NOT NULL AUTO_INCREMENT,"+
    "`Type` VARCHAR(45) NOT NULL,"+
    "`Number` VARCHAR(45) NOT NULL,"+
    "PRIMARY KEY (`id`));")
  if err != nil{
    fmt.Println(err)
  }else{
    fmt.Println("Table Passport created")
  }

  _, err = db.Query("CREATE TABLE `test`.`employees` ("+
  "`id` INT NOT NULL AUTO_INCREMENT,"+
  "`Name` VARCHAR(45) NOT NULL,"+
  "`Surname` VARCHAR(45) NOT NULL,"+
  "`Phone` VARCHAR(45) NOT NULL,"+
  "`CompanyId` INT NOT NULL,"+
  "`Passport` INT NOT NULL,"+
  "`Department` INT NOT NULL,"+
  "PRIMARY KEY (`id`),"+
  "INDEX `Comp_id_idx` (`CompanyId` ASC) VISIBLE,"+
  "INDEX `Passport_id_idx` (`Passport` ASC) VISIBLE,"+
  "INDEX `Department_id_idx` (`Department` ASC) VISIBLE,"+
  "KEY `Comp_id` (`CompanyId`),"+
  "CONSTRAINT `Comp_id` FOREIGN KEY (`CompanyId`)"+
  " REFERENCES `test`.`Company` (`id`)"+
  "ON DELETE NO ACTION "+
  "ON UPDATE CASCADE,"+
  "KEY `Passport_id` (`Passport`),"+
  "CONSTRAINT `Passport_id` FOREIGN KEY (`Passport`)"+
  " REFERENCES `test`.`Passport` (`id`)"+
  "ON DELETE NO ACTION "+
  "ON UPDATE CASCADE,"+
  "KEY `Department_id` (`Department`),"+
  "CONSTRAINT `Department_id` FOREIGN KEY (`Department`)"+
  " REFERENCES `test`.`Department` (`id`)"+
  "ON DELETE NO ACTION "+
  "ON UPDATE CASCADE);")

  if err != nil{
    fmt.Println(err)
  }else{
    fmt.Println("Table Employees created")
  }

  fmt.Println("Create tables")
  addRecords()
}

func addRecords()  {
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  provZapComp("Apple")
  provZapComp("Samsung")
  provZapComp("Nokia")

  provZapDep("Iphone10", "+1232567890", "Apple")
  provZapDep("Iphone11", "+1231567890", "Apple")
  provZapDep("Iphone12", "+1233567890", "Apple")

  provZapDep("SamsungA", "+2230567890", "Samsung")
  provZapDep("SamsungS", "+22345678890", "Samsung")
  provZapDep("SamsungG", "+2234569890", "Samsung")

  provZapDep("Nokia6", "+3234532890", "Nokia")
  provZapDep("Nokia9", "+3234547890", "Nokia")

  fmt.Println("Add records")
}

func provZapComp(name string)  {
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("Select * from test.company where Name = '%s'", name))
  if err != nil{
    fmt.Println(err)
  }

  pr := false
  for res.Next(){
    pr = true
  }

  if !pr{
    insert, err := db.Query(fmt.Sprintf("Insert into test.company (Name) values('%s')", name))
    if err != nil{
      fmt.Println(err)
    }
    defer insert.Close()
  }
}

func provZapDep(name string, phone string, name_comp string)  {
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("Select * from test.department where Name = '%s'", name))
  if err != nil{
    fmt.Println(err)
  }

  pr := false
  for res.Next(){
    pr = true
  }

  if !pr{
    insert, err := db.Query(fmt.Sprintf("Insert into test.department (Name, Phone, CompanyId)"+
    " values('%s', '%s', (Select id from test.company where Name = '%s'));", name, phone, name_comp))
    if err != nil{
      fmt.Println(err)
    }
    defer insert.Close()
  }
}
