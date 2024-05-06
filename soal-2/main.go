package main

import (
	"encoding/json"
	"fmt"
)

var comments = `
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
]`

type Comment struct {
	CommentID      int       `json:"commentId"`
	CommentContent string    `json:"commentContent"`
	Replies        []Comment `json:"replies,omitempty"`
}

func main() {
	var commentArr []Comment
	err := json.Unmarshal([]byte(comments), &commentArr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	totalKomentar := count(commentArr)
	fmt.Printf("Total Komentar adalah : %d", totalKomentar)
}

func count(comments []Comment) int {
	total := len(comments)
	for _, comment := range comments {
		total += count(comment.Replies)
	}
	return total
}
