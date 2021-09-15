package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"

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
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3307)/ads")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `ads`")
	if err != nil{
		panic(err)
	}
	for res.Next(){
		var ad Ad
		err = res.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.Date, &ad.Photo_links[0], &ad.Photo_links[1],
			&ad.Photo_links[2])
		if err != nil{
			panic(err)
		}

		adv = append(adv, ad)
	}
}

func (a *Ad) getAd(id int){
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3307)/ads")
	if err != nil{
		panic(err)
	}
	defer db.Close()	
}

var adv = []Ad{}
var showAd = Ad{}

func (a Ad) createAd(ad Ad, w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3307)/ads")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	insert ,err := db.Query(
	 fmt.Sprintf("INSERT INTO `ads` (`name`,`description`,`price`,`main_photo`,`photo1`,`photo2`) VALUES('%s','%s','%f','%s','%s','%s')",
	 ad.Name, ad.Description, ad.Price, ad.Photo_links[0], ad.Photo_links[1], ad.Photo_links[2]))
	if err != nil{
		panic(err)
	}
	defer insert.Close()
	http.Redirect(w,r,"/",302)
}

func ads_page(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		panic(err)
	}
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3307)/ads?parseTime=true")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `ads`")
	if err != nil{
		panic(err)
	}

	adv = []Ad{}
	for res.Next(){
		var ad Ad
		err = res.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.Date, &ad.Photo_links[0], &ad.Photo_links[1],
			&ad.Photo_links[2])
		if err != nil{
			panic(err)
		}

		adv = append(adv, ad)
	}
	tmpl.ExecuteTemplate(w, "index", adv)
}

func ad_page(w http.ResponseWriter, r *http.Request ){
	tmpl, err := template.ParseFiles("templates/ad.html", "templates/header.html", "templates/footer.html")
	vars := mux.Vars(r)
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3307)/ads?parseTime=true")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `ads` WHERE `id`= '%s'", vars["id"]))
	if err != nil{
		panic(err)
	}
	showAd = Ad{}
	for res.Next(){
		var ad Ad
		err = res.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.Date, &ad.Photo_links[0], &ad.Photo_links[1],
			&ad.Photo_links[2])
		if err != nil{
			panic(err)
		}

		showAd = ad
	}
	tmpl.ExecuteTemplate(w, "ad", showAd)
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
	}
	fval,err := strconv.ParseFloat(r.FormValue("price"),32)
	if err != nil{
		fmt.Println(err)
	}
	ad.Price = float32(fval)
	ad.createAd(ad, w, r);
}


func handleRequest(){
	rtr := mux.NewRouter()
	rtr.HandleFunc("/",ads_page).Methods("GET")
	rtr.HandleFunc("/create",create_ad).Methods("GET")
	rtr.HandleFunc("/save_ad",save_ad).Methods("POST")
	rtr.HandleFunc("/{id:[0-9]+}", ad_page).Methods("GET")

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}

func main(){
	
	handleRequest()
}