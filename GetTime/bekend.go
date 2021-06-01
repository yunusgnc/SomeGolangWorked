package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct {
	Name     string `json:"name"`
	Rating   string `json:"rating"`
	Overview string `json:"overview"`
	Imageurl string `json:"imageURL"`
	ID       int    `json:"id"`
}

func GetAllProduct() ([]Product, error) {
	resp, err := http.Get("http://localhost:3000/product")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, err
}
func AddItem() ([]Product, error) {
	product := Product{Name: "Karayip Korsanlar 2",
		Rating:   "9.9",
		Overview: "Set in the 22nd century, The Matrix tells the story of a computer hacker who joins a group of underground insurgents fighting the vast and powerful computers who now rule the earth.",
		Imageurl: "https://i.pinimg.com/736x/a9/60/9e/a9609e46787f8a69b12037ff92c760e2.jpg"}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post("http://localhost:3000/movies", "application/json;charest=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var prdct []Product
	json.Unmarshal(bodyBytes, &prdct)
	return prdct, err
}
func DeleteItem() (Product, error, string) {
	fmt.Println("4. Performing Http Delete...")
	todo := Product{}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:3000/movies/56", bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)

	return todo, err, bodyString
}
func main() {

	DeleteItem()

	get, _ := GetAllProduct()
	fmt.Println(get)
}
