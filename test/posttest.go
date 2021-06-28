package test

import (
	"../models"
	"../utils/database"
)

func CreatePosts() {
	posts := []models.Post{
		{Title: "Foot", Content: "contentTest1", PublisherID: 1, Category: "Sport"},
		{Title: "IPhone", Content: "contentTest2", PublisherID: 3, Category: "Tech"},
		{Title: "Samsung", Content: "contentTest3", PublisherID: 3, Category: "Tech"},
		{Title: "Macron", Content: "contentTest3", PublisherID: 3, Category: "Political"},
		{Title: "Le Cygne Noir", Content: "contentTest3", PublisherID: 8, Category: "Read"},
		{Title: "Miracle Morning", Content: "contentTest3", PublisherID: 10, Category: "Read"},
		{Title: "F1", Content: "contentTest3", PublisherID: 3, Category: "Sport"},
		{Title: "Dev", Content: "contentTest3", PublisherID: 1, Category: "Tech"},
		{Title: "CyberSecurity", Content: "contentTest3", PublisherID: 4, Category: "Tech"},
		{Title: "Tennis", Content: "contentTest3", PublisherID: 4, Category: "Sport"},
	}

	for i := 0; i < len(posts); i++ {
		post := models.Post{Title: posts[i].Title, Content: posts[i].Content, PublisherID: posts[i].PublisherID, Category: posts[i].Category}
		database.AddPost(post)
	}
}
