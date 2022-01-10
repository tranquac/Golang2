package main

import (
	"Building_Distributed_Applications_in_Gin/chapter3/handler"
	"Building_Distributed_Applications_in_Gin/chapter3/models"
	"context"
	"encoding/json"
	"io/ioutil"
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

func init1() {
	Recipes := make([]models.Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal(file, &Recipes)

	ctx = context.Background()
	MONGO_DATABASE = "demo"
	MONGODB_URI = "mongodb://admin:password@103.81.86.132:27017/demo?authSource=admin"
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(MONGODB_URI))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Conected to connect to MongoDB")
	collection := client.Database(MONGO_DATABASE).Collection("recipes")

	var listOfRecipes []interface{}
	for _, recipe := range Recipes {
		listOfRecipes = append(listOfRecipes, recipe)
	}
	//collection := client.Database(MONGO_DATABASE).Collection("recipes")
	insertManyResult, err := collection.InsertMany(
		ctx, listOfRecipes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted recipes: ",
		len(insertManyResult.InsertedIDs))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "103.81.86.132:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping()
	log.Println(status)
	recipesHandler = handler.NewRecipesHandler(ctx, collection, redisClient)
}

func main() {
	init1()
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipesHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipesHandler)
	router.GET("/recipes/search", recipesHandler.SearchRecipesHandler)
	router.GET("/recipes/:id", recipesHandler.GetRecipesHandler)
	router.Run(":8000")
}

