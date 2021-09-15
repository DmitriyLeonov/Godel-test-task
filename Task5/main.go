package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Ad struct{
	Id int
	Name string
	Description string
	Photo_links []string
	Price float32
}

func (a *Ad) getAds() {

}

func (a *Ad) getAd(id int){

}

func (a Ad) createAd(){

}

func ads_page(w http.ResponseWriter, r *http.Request){
	ad1 := Ad{0, "asdasd", "asdfsda", []string{"1","2","3"}, 2.20}
	tmpl, err := template.ParseFiles("templates/ads.html")
	if err != nil{
		tmpl.Execute(w, err)
	}
	tmpl.Execute(w, ad1)
}

func create_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "create ad")
}

func handleRequest(){
	http.HandleFunc("/",ads_page)
	http.HandleFunc("/create",create_page)
	http.ListenAndServe(":8080", nil)
}

func main(){
	handleRequest()
}