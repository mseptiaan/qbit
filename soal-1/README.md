# Case 1

Andi menjual beraneka ragam buah. Diketahui catatan buah yang dimiliki Andi saat ini adalah sebagai berikut.

```
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
]
```

Dimana data structure dari JSON data diatas adalah sebagai berikut
```
const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal FruitType = "LOCAL"
)

type FruitType string
type Fruits struct {
	ID    int       `json:"fruitId"`
	Name  string    `json:"fruitName"`
	Type  FruitType `json:"fruitType"`
	Stock int       `json:"stock"`
}
```
Buatlah sebuah program sederhana dengan menggunakan bahasa pemrograman Golang untuk menjawab usecase sebagai berikut:

1. Buah apa saja yang dimiliki Andi? (fruitName)
2. Andi memisahkan buahnya menjadi beberapa wadah berdasarkan tipe buah (fruitType). Berapa jumlah wadah yang dibutuhkan? Dan ada buah apa saja di masing-masing wadah?
3. Berapa total stock buah yang ada di masing-masing wadah?
4. Apakah ada komentar terkait kasus di atas?

# Answer
Open terminal and run this command:
```
go mod tidy
go run .
```

Output:
```
Andi memiliki buah : Apel, Kurma, Manggis, Jeruk Bali, Salak 

Berapa jumlah wadah yang dibutuhkan? : 2
Keranjang ke-1 isi buah dengan tipe Import dengan total buah dalam keranjang adalah 100
         Buah Apel dengan stok 10
         Buah Kurma dengan stok 20
         Buah apel dengan stok 50
         Buah KURMA dengan stok 20
Keranjang ke-2 isi buah dengan tipe Local dengan total buah dalam keranjang adalah 260
         Buah Manggis dengan stok 100
         Buah Jeruk Bali dengan stok 10
         Buah Salak dengan stok 150
```