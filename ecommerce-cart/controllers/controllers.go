package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iamvibs/golang-projects/ecommerce-cart/database"
	"github.com/iamvibs/golang-projects/ecommerce-cart/models"
	"github.com/iamvibs/golang-projects/ecommerce-cart/tokens"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()
var userCollection *mongo.Collection = database.UserData(database.Client, "Users")
var productCollection *mongo.Collection = database.ProductData(database.Client, "Products")

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassoword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassoword), []byte(userPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "login or password is incorrect"
		valid = false
	}
	return valid, msg
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
			return
		}

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()

		token, refreshtoken, _ := tokens.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}

		c.JSON(http.StatusCreated, "successfully signed in!")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var founduser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := tokens.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID)

		tokens.UpdateAllTokens(token, refreshToken, founduser.User_ID)
		c.JSON(http.StatusFound, founduser)
	}
}

func ProductViewerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var products models.Product
		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		products.Product_ID = primitive.NewObjectID()
		_, anyerr := productCollection.InsertOne(ctx, products)
		if anyerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not inserted"})
			return
		}
		c.JSON(http.StatusOK, "successfully aded admin")
	}
}

func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productList []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := productCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong")
			return
		}
		err = cursor.All(ctx, &productList)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		defer cursor.Close(ctx)
		if err := cursor.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		c.IndentedJSON(200, productList)
	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var SearchProducts []models.Product
		queryParam := c.Query("name")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusFound, gin.H{"error": "invalid search index"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		searchQueryDB, err := productCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex": queryParam}})
		if err != nil {
			c.IndentedJSON(404, "something went wrong fetching the data")
			return
		}

		defer searchQueryDB.Close(ctx)

		err = searchQueryDB.All(ctx, &SearchProducts)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		if err := searchQueryDB.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid request")
			return
		}

		c.IndentedJSON(200, SearchProducts)
	}
}
