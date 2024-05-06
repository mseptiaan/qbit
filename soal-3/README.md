# Case 3

Buatlah aplikasi ecommerce sederhana dengan ketentuan sebagai berikut:
1. Fitur
    1. Home / Landing Page
    2. Authentication (email & password)
    3. Product
    4. Add to cart & checkout
2. Tech stack
    1. Backend menggunakan Golang (framework bebas)
3. Buat struktur database (db scheme) dari kebutuhan ini.

# Answers

Saya membuat program sederhana dari ecommerce sederhana
1. Fitur
   1. Authentication & Register
   2. See Product & Search Product
   3. Add to cart & checkout
2. Tech stack
   1. Backend menggunakan Golang dengan dasar API menggunakan gRPC (Google Remote Procedure Call) dimana gRPC lebih unggul dari Rest API pada umumnya
   2. Database ORM (Object Relation Mapping) saya menggunakan GORM
   3. Struktur folder sudah memisahkan antara Model, Repository dan Service untuk mempermudah pengembangan aplikasi.
   4. Authentication sudah menggunakan JWT dengan
   5. Sudah ter-Dockerize untuk memudakan deploy kedalam server
3.  struktur database ![db diagram](/db-diagram.png)

## [Postman Collection](https://www.postman.com/oa-oe-oa-oe/workspace/qbit-septian/collection/21858721-dce892d8-ce2f-4ecb-8c1f-c5de37d3a45b)