package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Drink struct {
	Name         string `json:"strDrink"`
	Instructions string `json:"strInstructions"`
}

type CocktailBartender struct {
	query  string
	Drinks []Drink
}

func (cb *CocktailBartender) Start() {

	for {
		fmt.Print("What would you like to drink? >")

		fmt.Scanln(&cb.query)
		fmt.Println(cb.query)
		if cb.query != "nothing" {
			cb.FetchDrinks()
		} else {
			fmt.Println("Goodbye!")
			os.Stdout.Close()
		}
	}
}

func (cb *CocktailBartender) FetchDrinks() {
	urlStr, err := url.Parse("https://www.thecocktaildb.com/api/json/v1/1/search.php")
	if err != nil {
		log.Fatal(err)
	}

	q := urlStr.Query()
	q.Add("s", cb.query)
	urlStr.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", urlStr.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	drinks := CocktailBartender{}
	json.NewDecoder(resp.Body).Decode(&drinks)

	cb.Drinks = append(cb.Drinks, drinks.Drinks...)
	//fmt.Println(cb)
	cb.PrintCoctailInstrs()
}

func (cb *CocktailBartender) PrintCoctailInstrs() {
	fmt.Printf("If you want to prepare %s, make:\n", cb.Drinks[0].Name)

	instrSlice := strings.Split(cb.Drinks[0].Instructions, ".")
	//remove leading spaces
	for i := range instrSlice {
		instrSlice[i] = strings.TrimSpace(instrSlice[i])
	}
	count := 1
	for _, in := range instrSlice {
		fmt.Printf("%v. %s.\n", count, in)
		count++
	}
}

func main() {
	b := CocktailBartender{}
	b.Start()
}
