package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kevin-dev71/go-react-crud/models"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	app := fiber.New()
	app.Use(cors.New())
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb"))

	if err != nil {
		panic(err)
	}

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User

		coll := client.Database("gomongodb").Collection("users")
		results, err := coll.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"data": users,
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user)

		coll := client.Database("gomongodb").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{
			{
				Key: "name", Value: user.Name,
			},
		})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})
	app.Listen(":" + port)
	fmt.Println("Server running in port " + port)
}
