package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Ad struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Photo_links [3]string `json:"photo_links"`
	Price float32 `json:"price"`
	Date time.Time `json:"date"`
}

func (a *Ad) getAds() {

}

func (a *Ad) getAd(id int){
	
}

func (a Ad) createAd(ad Ad){
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3307)/ads")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	insert ,err := db.Query("INSERT INTO `ads`")
}

func ads_page(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "index", nil)
}

func create_ad(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "create", nil)
}

func save_ad(w http.ResponseWriter, r *http.Request){
	ad := Ad{
		Name:        r.FormValue("name"),
		Description:  r.FormValue("description"),
		Photo_links: [3]string{ r.FormValue("main_photo"),
		 r.FormValue("photo1"),
		 r.FormValue("photo2")},
		Price:       0.0,
		Date:        time.Now(),
	}
	fval,err := strconv.ParseFloat(r.FormValue("price"),32)
	if err != nil{
		fmt.Println(err)
	}
	ad.Price = float32(fval)
	ad.createAd(ad);
}


func handleRequest(){
	http.HandleFunc("/",ads_page)
	http.HandleFunc("/create",create_ad)
	http.HandleFunc("/save_ad",save_ad)
	http.ListenAndServe(":8080", nil)
}

func main(){
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3307)/ads")
	if err != nil{
		panic(err)
	}
	defer db.Close()
	fmt.Println("connected")
	handleRequest()
}