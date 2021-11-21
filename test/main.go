package main

import(
  "fmt"//print
  "net/http"//inet
  "html/template"//html
  "strconv"//toInt

  "database/sql"

  "github.com/gorilla/mux"
  _ "github.com/go-sql-driver/mysql" //"github.com/go-sql-driver/mysql"
)

type CompanyS struct{
  Id int`json:"id"`
  Name string`json:"Name"`
}

type Employee struct{
  Id int
  NameS string
  SurnameS string
  PhoneS string
  CompanyName string
  PassportT string
  PassportN string
  NameD string
  PhoneD string
}

type CompDep struct{
  NameC string
  NameD string
  Phone string
}

var idSG = ""
var emps = []Employee{}
var prDel = false

func index(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
  if err != nil{
    fmt.Println(err)
  }
  t.ExecuteTemplate(w, "index", nil)//из-за того что внутри шаблонов будет динамическое подключение
}

func company_dep(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("templates/company_dep.html", "templates/header.html", "templates/footer.html")
  if err != nil{
    fmt.Println(err)
  }

  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query("select b.Name, a.Name, a.Phone  from test.department a join test.company b where a.CompanyId=b.id;")
  if err != nil{
    fmt.Println(err)
  }

  CD := []CompDep{}
  for res.Next(){
    var cd CompDep
    err = res.Scan(&cd.NameC, &cd.NameD, &cd.Phone) //перебираем все ряды, вытаскиваем значения и проверям успешно ли вытянули
    if err != nil{
      panic(err)//типа cath
    }

    CD = append(CD, cd)
  }

  t.ExecuteTemplate(w, "company_dep", CD)//из-за того что внутри шаблонов будет динамическое подключение
}

func create(w http.ResponseWriter, r *http.Request)  {
  t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

  if err != nil{
  //  fmt.Fprintf(w, err.Error())
    panic(err)
  }
  newS := ""

  if idSG != ""{
    newS = fmt.Sprintf("Id данного сотрудник id: %s",idSG)
    idSG = ""
  }

  t.ExecuteTemplate(w, "create", newS)//из-за того что внутри шаблонов будет динамическое подключение
}

func save_article(w http.ResponseWriter, r *http.Request){
  nameS := r.FormValue("nameS")
  surnameS := r.FormValue("surnameS")
  phoneS := r.FormValue("phoneS")
  companyId := r.FormValue("companyId")
  passportT := r.FormValue("passportT")
  passportN := r.FormValue("passportN")
  departmentN := r.FormValue("departmentN")
  departmentPh := r.FormValue("departmentPh")

  fmt.Println(nameS + " " + surnameS + " " + phoneS+ " " + companyId + " " + passportT + " " + passportN + " " + departmentN +" "+departmentPh)

  idP := addPassport(passportT, passportN)
  idC := getCompany(companyId)
  idD := getDepartment(departmentN,departmentPh, idC)
  idS := addEmployee(nameS, surnameS, phoneS, idC, idP, idD)

  fmt.Println(fmt.Sprintf("Passport: %d",idP))
  fmt.Println(fmt.Sprintf("CompanyId: %d",idC))
  fmt.Println(fmt.Sprintf("departmentID: %d",idD))
  fmt.Println(fmt.Sprintf("Employee: %d",idS))

  idSG = strconv.Itoa(idS)

  http.Redirect(w, r, "/create", http.StatusSeeOther)
}

func showComp(w http.ResponseWriter, r *http.Request)  {
  t, err := template.ParseFiles("templates/showComp.html", "templates/header.html", "templates/footer.html")

  if err != nil{
    panic(err)
  }

  emp := []Employee{}
  if emps != nil{
    emp = emps
    emps = nil
  }
  t.ExecuteTemplate(w, "showComp", emp)
}

func get_SC(w http.ResponseWriter, r *http.Request){
  nameS := r.FormValue("nameS")
  fmt.Println(nameS)

  idC := getCompany(nameS)

  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("select a.id, a.Name, a.Surname, a.Phone, b.Name, c.Type, c.Number, d.Name, d.Phone "+
    "from test.employees a join test.company b join test.passport c join test.department d " +
    "where a.CompanyId = '%d' and b.id = a.CompanyId and a.passport = c.id and a.Department = d.id", idC))
  if err != nil{
    fmt.Println(err)
  }

  emps = []Employee{}
  for res.Next(){
    eml := Employee{}
    err = res.Scan(&eml.Id, &eml.NameS, &eml.SurnameS, &eml.PhoneS,
       &eml.CompanyName, &eml.PassportT, &eml.PassportN, &eml.NameD, &eml.PhoneD ) //перебираем все ряды, вытаскиваем значения и проверям успешно ли вытянули
    if err != nil{
      panic(err)//типа cath
    }

    emps = append(emps, eml)
  }

  http.Redirect(w, r, "/showComp", http.StatusSeeOther)
}

func delete(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("templates/delete.html", "templates/header.html", "templates/footer.html")

  if err != nil{
    panic(err)
  }

  newS := ""

  if prDel{
    newS = "Сотрудник удален"
    prDel = false
  }

  t.ExecuteTemplate(w, "delete", newS)
}

func del_post(w http.ResponseWriter, r *http.Request){
  nameS := r.FormValue("nameS")
  fmt.Println(nameS)

  id, err := strconv.Atoi(nameS)
  if err != nil{
    panic(err)
  }

  prDel = deleteS(id)

  http.Redirect(w, r, "/delete", http.StatusSeeOther)
}

func getCD(w http.ResponseWriter, r *http.Request)  {
  t, err := template.ParseFiles("templates/getCD.html", "templates/header.html", "templates/footer.html")

  if err != nil{
    panic(err)
  }

  emp := []Employee{}
  if emps != nil{
    emp = emps
    emps = nil
  }
  t.ExecuteTemplate(w, "getCD", emp)
}

func get_CD(w http.ResponseWriter, r *http.Request){
  nameC := r.FormValue("nameC")
  nameD:= r.FormValue("nameD")
  fmt.Println(nameC + " " + nameD)

  idC := getCompany(nameC)
  idD := getDepComp(idC, nameD)

  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("select a.id, a.Name, a.Surname, a.Phone, b.Name, c.Type, c.Number, d.Name, d.Phone "+
    "from test.employees a join test.company b join test.passport c join test.department d " +
    "where a.CompanyId = '%d' and a.Department = '%d' and b.id = a.CompanyId and a.passport = c.id and a.Department = d.id", idC, idD))
  if err != nil{
    fmt.Println(err)
  }

  emps = []Employee{}
  for res.Next(){
    eml := Employee{}
    err = res.Scan(&eml.Id, &eml.NameS, &eml.SurnameS, &eml.PhoneS,
       &eml.CompanyName, &eml.PassportT, &eml.PassportN, &eml.NameD, &eml.PhoneD ) //перебираем все ряды, вытаскиваем значения и проверям успешно ли вытянули
    if err != nil{
      panic(err)//типа cath
    }

    emps = append(emps, eml)
  }

  http.Redirect(w, r, "/getCD", http.StatusSeeOther)
}

func update(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("templates/update.html", "templates/header.html", "templates/footer.html")

  if err != nil{
    panic(err)
  }

  newS := ""

  if prDel{
    newS = "Данные сотрудника изменены"
    prDel = false
  }

  t.ExecuteTemplate(w, "update", newS)
}

func updateS(w http.ResponseWriter, r *http.Request){
  nameC := r.FormValue("nameS")//создаем объект, идет на основе библиотеки, передаем r

  t, err := template.ParseFiles("templates/updateS.html", "templates/header.html", "templates/footer.html")
  if err != nil{
    fmt.Println(err)
  }

  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("select a.id, a.Name, a.Surname, a.Phone, b.Name, c.Type, c.Number, d.Name, d.Phone "+
    "from test.employees a join test.company b join test.passport c join test.department d " +
    "where a.id = '%s' and b.id = a.CompanyId and a.passport = c.id and a.Department = d.id", nameC))
  if err != nil{
    fmt.Println(err)
  }

  eml := Employee{}
  for res.Next(){
    err = res.Scan(&eml.Id, &eml.NameS, &eml.SurnameS, &eml.PhoneS,
       &eml.CompanyName, &eml.PassportT, &eml.PassportN, &eml.NameD, &eml.PhoneD ) //перебираем все ряды, вытаскиваем значения и проверям успешно ли вытянули
    if err != nil{
      panic(err)//типа cath
    }
  }

  t.ExecuteTemplate(w, "updateS", eml)
}

func save_update(w http.ResponseWriter, r *http.Request){
  nameID := r.FormValue("nameID")
  nameS := r.FormValue("nameS")
  surnameS := r.FormValue("surnameS")
  phoneS := r.FormValue("phoneS")
  companyId := r.FormValue("companyId")
  passportT := r.FormValue("passportT")
  passportN := r.FormValue("passportN")
  departmentN := r.FormValue("departmentN")
  departmentPh := r.FormValue("departmentPh")

  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
  if err != nil{
    fmt.Println("Не подключились")
  }
  defer db.Close()
  q := fmt.Sprintf("select a.id, a.Name, a.Surname, a.Phone, b.Name, c.Type, c.Number, d.Name, d.Phone "+
    "from test.employees a join test.company b join test.passport c join test.department d " +
    "where a.id = '%s' and b.id = a.CompanyId and a.passport = c.id and a.Department = d.id", nameID)

  res, err := db.Query(q)
  if err != nil{
    fmt.Println(err)
  }

  eml := Employee{}
  for res.Next(){
    err = res.Scan(&eml.Id, &eml.NameS, &eml.SurnameS, &eml.PhoneS,
       &eml.CompanyName, &eml.PassportT, &eml.PassportN, &eml.NameD, &eml.PhoneD ) //перебираем все ряды, вытаскиваем значения и проверям успешно ли вытянули
    if err != nil{
      panic(err)//типа cath
    }
  }

  str := "UPDATE test.employees SET "

  i := 0
  if nameS != eml.NameS{
    str += fmt.Sprintf("Name = '%s'", nameS)
    i++
  }

  if surnameS != eml.SurnameS{
    if i != 0{
      str += ", "
    }else{
      i++
    }
    str += fmt.Sprintf("Surname = '%s'", surnameS)
  }

  if phoneS != eml.PhoneS{
    if i != 0{
      str += ", "
    }else{
      i++
    }
    str += fmt.Sprintf("Phone = '%s'", phoneS)
  }

  if companyId != eml.CompanyName{
    if i != 0{
      str += ", "
    }else{
      i++
    }

    idC := getCompany(companyId)
    str += fmt.Sprintf("CompanyId = '%d'", idC)
  }

  if (passportT != eml.PassportT)||(passportN != eml.PassportN){
    updatePassport(eml.PassportT, eml.PassportN, passportT, passportN)
  }

  if (departmentN != eml.NameD)||(departmentPh != eml.PhoneD){
    if i != 0{
      str += ", "
    }else{
      i++
    }

    idD := getDepartment(departmentN, departmentPh, getCompany(companyId))
    str += fmt.Sprintf("Department = '%d'", idD)
  }

  str += fmt.Sprintf(" WHERE id = '%d';", eml.Id)

  if i > 0{
    _, err := db.Query(str)
    if err != nil{
      fmt.Println(err)
    }
  }

  fmt.Println(str)

  //fmt.Println(nameS + " " + surnameS + " " + phoneS+ " " + companyId + " " + passportT + " " + passportN + " " + departmentN +" "+departmentPh)
  http.Redirect(w, r, "/update", http.StatusSeeOther)
}

func handleFunc()  {
  rtr := mux.NewRouter()

  rtr.HandleFunc("/", index).Methods("GET")//1. адрес 2. вызов метода 3. по ссылке перейдешь, без отправки инфы
  rtr.HandleFunc("/company_dep", company_dep).Methods("GET")

  rtr.HandleFunc("/create", create).Methods("GET")//способ перехода, по ссылке перейдешь, без отправки инфы
  rtr.HandleFunc("/save_article", save_article).Methods("POST")//по ссылке не попадешь

  rtr.HandleFunc("/showComp", showComp).Methods("GET")
  rtr.HandleFunc("/get_SC", get_SC).Methods("POST")

  rtr.HandleFunc("/delete", delete).Methods("GET")
  rtr.HandleFunc("/del_post", del_post).Methods("POST")

  rtr.HandleFunc("/getCD", getCD).Methods("GET")
  rtr.HandleFunc("/get_CD", get_CD).Methods("POST")

  rtr.HandleFunc("/update", update).Methods("GET")
  rtr.HandleFunc("/updateS", updateS).Methods("GET")
  rtr.HandleFunc("/save_update", save_update).Methods("POST")

  http.Handle("/", rtr)//обработка всех url адресов через rtr библиотеку горила
  http.ListenAndServe(":8080", nil)
}

func main()  {
  createDB()
  handleFunc()
}
