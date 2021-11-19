package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type Fruits struct {
	Id       string
	Name     string
	Quantity int
	Price    string
}

type Vegetables struct {
	ProductId   string
	ProductName string
	Quantity    int
	Price       string
}

type Grains struct {
	ItemId   string
	ItemName string
	Quantity int
	Price    string
}

var wg sync.WaitGroup
var Fruit []Fruits
var Vegetable []Vegetables
var Grain []Grains

var summarydetails []interface{}

const fruitsUrl = "https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b"
const vegetablesUrl = "https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c"
const grainsUrl = "https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148"

//Challenge 1

func getName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)

	response, err := http.Get(fruitsUrl)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}

	for _, p := range Fruit {
		if p.Name == par["item"] {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	response1, er := http.Get(vegetablesUrl)
	if er != nil {
		panic(err)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	er1 := json.Unmarshal(content1, &Vegetable)
	if er1 != nil {
		panic(er1)
	}

	for _, p := range Vegetable {
		if p.ProductName == par["item"] {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}

	response2, err2 := http.Get(grainsUrl)
	if err2 != nil {
		panic(err)
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)
	}

	for _, p := range Grain {
		if p.ItemName == par["item"] {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")
}

//Challenge 2

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	j, _ := strconv.Atoi(par["quantity"])

	response, err := http.Get(fruitsUrl)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}

	for _, p := range Fruit {
		if p.Quantity >= j {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	response1, er := http.Get(vegetablesUrl)
	if er != nil {
		panic(err)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	er1 := json.Unmarshal(content1, &Vegetable)
	if er1 != nil {
		panic(er1)
	}

	for _, p := range Vegetable {
		if p.Quantity >= j {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}

	response2, err2 := http.Get(grainsUrl)
	if err2 != nil {
		panic(err)
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)
	}

	for _, p := range Grain {
		if p.Quantity >= j {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")
}

//Challenge 3

func getQuantity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	var str string
	j, _ := strconv.Atoi(par["quantity"])
	if !strings.Contains(par["price"], "$") {
		str = "$" + par["price"]
	} else if strings.Index(par["price"], "$") > 1 {
		k := strings.TrimRight(par["price"], "$")
		str = "$" + k
	} else {
		str = par["price"]
	}

	response, err := http.Get(fruitsUrl)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}

	for _, p := range Fruit {
		if p.Quantity >= j && p.Price == str {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	response1, er := http.Get(vegetablesUrl)
	if er != nil {
		panic(err)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	er1 := json.Unmarshal(content1, &Vegetable)
	if er1 != nil {
		panic(er1)
	}

	for _, p := range Vegetable {
		if p.Quantity >= j && p.Price == str {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}

	response2, err2 := http.Get(grainsUrl)
	if err2 != nil {
		panic(err)
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)
	}

	for _, p := range Grain {
		if p.Quantity >= j && p.Price == str {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")
}

//Challenge 4

func getitems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(summarydetails)
}

//Challenge 5

func getFruits() {
	response, err := http.Get(fruitsUrl)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}
	wg.Done()
	defer response.Body.Close()
}

func getVegetables() {
	response1, er := http.Get(vegetablesUrl)

	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err2 := json.Unmarshal(content1, &Vegetable)
	if err2 != nil {
		panic(err2)
	}
	wg.Done()
	defer response1.Body.Close()
}

func getGrains() {
	response2, er1 := http.Get(grainsUrl)

	if er1 != nil {
		panic(er1)
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)

	}
	wg.Done()
	defer response2.Body.Close()
}

func getfastitem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)

	go getFruits()
	go getVegetables()
	go getGrains()
	wg.Add(3)
	for _, p := range Fruit {
		if p.Name == par["item"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	for _, c := range Vegetable {
		if c.ProductName == par["item"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	for _, g := range Grain {
		if g.ItemName == par["item"] {
			json.NewEncoder(w).Encode(g)
			return
		}
	}
	json.NewEncoder(w).Encode("!! Item Not Found")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/show-summary", getitems).Methods("GET")
	r.HandleFunc("/buy-item-qty-price/{quantity}/{price}", getQuantity).Methods("GET")
	r.HandleFunc("/buy-item-name/{item}", getName).Methods("GET")
	r.HandleFunc("/fast-buy-item/{item}", getfastitem).Methods("GET")
	r.HandleFunc("/buy-item-qty/{quantity}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
	wg.Wait()
}
