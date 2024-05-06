# Case 2

Diketahui data berikut ini.

```
[
  {
    "commentId": 1,
    "commentContent": "Hai",
    "replies": [
      {
        "commentId": 11,
        "commentContent": "Hai juga",
        "replies": [
          {
            "commentId": 111,
            "commentContent": "Haai juga hai jugaa"
          },
          {
            "commentId": 112,
            "commentContent": "Haai juga hai jugaa"
          }
        ]
      },
      {
        "commentId": 12,
        "commentContent": "Hai juga",
        "replies": [
          {
            "commentId": 121,
            "commentContent": "Haai juga hai jugaa"
          }
        ]
      }
    ]
  },
  {
    "commentId": 2,
    "commentContent": "Halooo"
  }
]
```

Buatlah sebuah program sederhana untuk membaca data diatas dan menghitung berapa banyak total komentar yang ada (termasuk semua balasan komentar). Untuk contoh data di atas, total komentar adalah 7 komentar.