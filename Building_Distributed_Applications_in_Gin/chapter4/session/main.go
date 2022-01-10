package main

import (
	"Building_Distributed_Applications_in_Gin/chapter4/session/handler"
	"context"
	"log"

	//"Building_Distributed_Applications_in_Gin/chapter3/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var MONGODB_URI string
var MONGO_DATABASE string
var recipesHandler *handler.RecipesHandler
var authHandler *handler.AuthHandler

func init() {
	// Recipes := make([]models.Recipe, 0)
	// file, _ := ioutil.ReadFile("recipes.json")
	// _ = json.Unmarshal(file, &Recipes)

	ctx = context.Background()
	MONGO_DATABASE = "demo"
	MONGODB_URI = "mongodb://admin:password@103.81.86.132:27017/demo?authSource=admin"
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(MONGODB_URI))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Conected to connect to MongoDB")
	collection := client.Database(MONGO_DATABASE).Collection("recipes")

	// var listOfRecipes []interface{}
	// for _, recipe := range recipes {
	// 	listOfRecipes = append(listOfRecipes, recipe)
	// }
	//collection := client.Database(MONGO_DATABASE).Collection("recipes")
	// insertManyResult, err := collection.InsertMany(
	// 	ctx, listOfRecipes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Inserted recipes: ",
	// 	len(insertManyResult.InsertedIDs))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "103.81.86.132:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping()
	log.Println(status)
	recipesHandler = handler.NewRecipesHandler(ctx, collection, redisClient)
	authHandler = handler.NewAuthHandler(ctx, collection)
}

func main() {
	router := gin.Default()
	router.POST("/signin", authHandler.SignInHandlerSession)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.GET("/recipes/search", recipesHandler.SearchRecipesHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipesHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipesHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetRecipesHandler)
		router.POST("/signout", authHandler.SignOutHandler)
	}
	router.Run(":8000")
}
