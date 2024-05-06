package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"reflect"
	"strings"
)

var fruits = `
[
  {
    "fruitId": 1,
    "fruitName": "Apel",
    "fruitType": "IMPORT",
    "stock": 10
  },
  {
    "fruitId": 2,
    "fruitName": "Kurma",
    "fruitType": "IMPORT",
    "stock": 20
  },
  {
    "fruitId": 3,
    "fruitName": "apel",
    "fruitType": "IMPORT",
    "stock": 50
  },
  {
    "fruitId": 4,
    "fruitName": "Manggis",
    "fruitType": "LOCAL",
    "stock": 100
  },
  {
    "fruitId": 5,
    "fruitName": "Jeruk Bali",
    "fruitType": "LOCAL",
    "stock": 10
  },
  {
    "fruitId": 5,
    "fruitName": "KURMA",
    "fruitType": "IMPORT",
    "stock": 20
  },
  {
    "fruitId": 5,
    "fruitName": "Salak",
    "fruitType": "LOCAL",
    "stock": 150
  }
]`

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

type FruitType string
type Fruits struct {
	ID    int       `json:"fruitId"`
	Name  string    `json:"fruitName"`
	Type  FruitType `json:"fruitType,omitempty"`
	Stock int       `json:"stock"`
}

type FruitArr struct {
	Fruits []Fruits
}

type FruitSum struct {
	FruitType  string   `json:"fruitType"`
	TotalStock int      `json:"totalStock"`
	Fruits     []Fruits `json:"fruits"`
}

func (f *FruitArr) OwnFruits() []string {
	return RemoveDuplicateStr(f.Fruits, "Name")
}

func (f *FruitArr) SeparateFruitType() []FruitSum {
	var fruitSum []FruitSum
	fruit := RemoveDuplicateStr(f.Fruits, "Type")

	for _, s := range fruit {
		sum := 0
		var fruitData []Fruits
		for _, t := range f.Fruits {
			if s == cases.Title(language.English).String(strings.ToLower(string(t.Type))) {
				sum += t.Stock
				fruitData = append(fruitData, t)
			}
		}
		fruitSum = append(fruitSum, FruitSum{
			FruitType:  s,
			TotalStock: sum,
			Fruits:     fruitData,
		})
	}

	return fruitSum
}

func RemoveDuplicateStr(arr []Fruits, key string) []string {
	encountered := map[interface{}]struct{}{}
	result := make([]string, 0)

	for i := 0; i < reflect.ValueOf(arr).Len(); i++ {
		fieldValue := reflect.ValueOf(arr).Index(i).FieldByName(key).Interface()
		value := cases.Title(language.English).String(strings.ToLower(fmt.Sprintf("%s", fieldValue)))

		if _, ok := encountered[value]; !ok {
			encountered[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}

func main() {
	var result []Fruits
	// Unmarshal the JSON into the map
	err := json.Unmarshal([]byte(fruits), &result)
	if err != nil {
		log.Fatal(err)
	}

	var fruit FruitArr
	fruit.Fruits = result

	// Buah apa saja yang dimiliki Andi? (fruitName)
	fmt.Printf("Andi memiliki buah : %s \n\n", strings.Join(fruit.OwnFruits(), ", "))

	// Andi memisahkan buahnya menjadi beberapa wadah berdasarkan tipe buah (fruitType). Berapa jumlah wadah yang dibutuhkan? Dan ada buah apa saja di masing-masing wadah? dan Berapa total stock buah yang ada di masing-masing wadah?
	fmt.Printf("Berapa jumlah wadah yang dibutuhkan? : %d\n", len(fruit.SeparateFruitType()))
	for i, f := range fruit.SeparateFruitType() {
		fmt.Printf("Keranjang ke-%d isi buah dengan tipe %s dengan total buah dalam keranjang adalah %d", i+1, f.FruitType, f.TotalStock)
		for _, fruitData := range f.Fruits {
			fmt.Printf("\n\t Buah %s dengan stok %d", fruitData.Name, fruitData.Stock)
		}
		fmt.Printf("\n")
	}
}
