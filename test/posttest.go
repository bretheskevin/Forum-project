package test

import (
	"../models"
	"../utils/database"
)

func CreatePosts() {
	db := database.Connect()

	posts := []models.Post{
		{Title: "test1", Content: "contentTest1", PublisherID: 1, Category: "categoryTest1"},
		{Title: "test2", Content: "contentTest2", PublisherID: 1, Category: "categoryTest2"},
		{Title: "test2", Content: "contentTest3", PublisherID: 1, Category: "categoryTest3"},
	}

	for i := 0; i < len(posts); i++ {
		post := models.Post{Title: posts[i].Title, Content: posts[i].Content, PublisherID: posts[i].PublisherID, Category: posts[i].Category}
		database.AddPost(db, post)
	}
}
