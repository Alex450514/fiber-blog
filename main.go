package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Mock Data
var posts = []Post{
	{ID: "1", Title: "First Post", Content: "Hello Fiber!"},
	{ID: "2", Title: "Second Post", Content: "Fiber is fast 🚀"},
}

func main() {
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber Blog API")
	})

	app.Get("/posts", getPosts)
	app.Get("/posts/:id", getPostByID)
	app.Post("/posts", createPost)

	log.Fatal(app.Listen(":3000"))
}

// Handlers
func getPosts(c *fiber.Ctx) error {
	return c.JSON(posts)
}

func getPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, post := range posts {
		if post.ID == id {
			return c.JSON(post)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
}

func createPost(c *fiber.Ctx) error {
	post := new(Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	posts = append(posts, *post)
	return c.Status(201).JSON(post)
}
