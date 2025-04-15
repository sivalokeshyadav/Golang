This file contain all golang tutorial with example of restapi

# Table of Contents
## Core Topics
### Packages and Imports

### Variables and Constants

### Data Types (int, string, bool, etc.)

### Pointers

### Functions

### Structs and Methods

### Interfaces

### Error Handling (error type)

## Advanced Topics
### Concurrency (Goroutines & Channels)

### Context package (for cancellation, timeout)

### Reflection

### Defer, Panic, Recover


# Build a REST API with Golang and MongoDB - Gin-gonic Version

Representational state transfer (REST) is an architectural pattern that guides an Application programming interface(API) design and development. REST APIs have become the standard of communication between the server part of the product and its client to increase performance, scalability, simplicity, modifiability, visibility, portability, and reliability.


Gin-gonic, popularly known as Gin, is an HTTP web framework written in Golang with performance and productivity support. Gin uses a custom version of HttpRouter, a lightweight, high-performance HTTP request router that navigates through API routes faster than most frameworks out there.

MongoDB is a document-based database management program used as an alternative to relational databases. MongoDB supports working with large sets of distributed data with options to store or retrieve information seamlessly.



Prerequisites
The following steps in this post require Golang experience. Experience with MongoDB isn’t a requirement, but it’s nice to have.

We will also be needing the following:

A MongoDB account to host database. Signup is completely free.
Postman or any API testing application of your choice

Let’s code
Getting Started
To get started, we need to navigate to the desired directory and run the command below in our terminal

# mkdir gin-mongo-api && cd gin-mongo-api
This command creates a gin-mongo-api folder and navigates into the project directory.

Next, we need to initialize a Go module to manage project dependencies by running the command below:

# go mod init gin-mongo-api
This command will create a go.mod file for tracking project dependencies.

We proceed to install the required dependencies with:

# go get -u github.com/gin-gonic/gin # go.mongodb.org/mongo-driver/mongo # github.com/joho/godotenv github.com/go-playground/validator/v10
# github.com/gin-gonic/gin is a framework for building web applications.

# go.mongodb.org/mongo-driver/mongo is a driver for connecting to MongoDB.

# github.com/joho/godotenv is a library for managing environment variables.

# github.com/go-playground/validator/v10 is a library for validating structs and fields.

Application Entry Point
With the project dependencies installed, we need to create main.go file in the root directory and add the snippet below:

```

package main

import "github.com/gin-gonic/gin"

func main() {
        router := gin.Default()

        router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "data": "Hello from Gin-gonic & mongoDB",
                })
        })

        router.Run("localhost:6000") 
}

```
The snippet above does the following:

Import the required dependencies.
Initialize a Gin router using the Default configuration. The Default function configures Gin router with default middlewares (logger and recovery).
Use the Get function to route to / path and a handler function that returns a JSON of Hello from Gin-gonic & mongoDB.
Use the Run function to attach the router to an http.Server and starts listening and serving HTTP requests on localhost:6000.
Next, we can test our application by starting the development server by running the command below in our terminal.

# go run main.go
Testing the app

Modularization in Golang
It is essential to have a good folder structure for our project. Good project structure simplifies how we work with dependencies in our application and makes it easier for us and others to read our codebase.
To do this, we need to create configs, controllers, models, responses and routes folder in our project directory.

Updated project folder structure

PS: The go.sum file contains all the dependency checksums, and is managed by the go tools. We don’t have to worry about it.

configs is for modularizing project configuration files

controllers is for modularizing application logics.

models is for modularizing data and database logics.

responses is for modularizing files describing the response we want our API to give. This will become clearer later on.

routes is for modularizing URL pattern and handler information.

Setting up MongoDB
With that done, we need to log in or sign up into our MongoDB account. Click the project dropdown menu and click on the New Project button.

New Project

Enter the golang-api as the project name, click Next, and click Create Project..

enter project name
Create Project

Click on Build a Database



Select Shared as the type of database.

Shared highlighted in red

Click on Create to setup a cluster. This might take sometime to setup.

Creating a cluster

Next, we need to create a user to access the database externally by inputting the Username, Password and then clicking on Create User. We also need to add our IP address to safely connect to the database by clicking on the Add My Current IP Address button. Then click on Finish and Close to save changes.

Create user
Add IP

On saving the changes, we should see a Database Deployments screen, as shown below:

Database Screen

Connecting our application to MongoDB
With the configuration done, we need to connect our application with the database created. To do this, click on the Connect button

Connect to database

Click on Connect your application, change the Driver to Go and the Version as shown below. Then click on the copy icon to copy the connection string.

connect application
Copy connection string

Setup Environment Variable
Next, we must modify the copied connection string with the user's password we created earlier and change the database name. To do this, first, we need to create a .env file in the root directory, and in this file, add the snippet below:

# MONGOURI=mongodb+srv://<YOUR USERNAME HERE>:<YOUR PASSWORD HERE>@cluster0.e5akf.mongodb.net/myFirstDatabese?retryWrites=true&w=majority
Sample of a properly filled connection string below:

# MONGOURI=mongodb+srv://malomz:malomzPassword@cluster0.e5akf.mongodb.net/golangDB?retryWrites=true&w=majority
Updated folder structure with .env file

Load Environment Variable
With that done, we need to create a helper function to load the environment variable using the github.com/joho/godotenv library we installed earlier. To do this, we need to navigate to the configs folder and in this folder, create an env.go file and add the snippet below:


```
package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGOURI")
}
```

The snippet above does the following:

Import the required dependencies.
Create an EnvMongoURI function that checks if the environment variable is correctly loaded and returns the environment variable.
Connecting to MongoDB
To connect to the MongoDB database from our application, first we need to navigate to the configs folder and in this folder, create a setup.go file and add the snippet below:


```
package configs

import (
    "context"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client  {
    client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
    if err != nil {
        log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("golangAPI").Collection(collectionName)
    return collection
}

```
The snippet above does the following:

Import the required dependencies.
Create a ConnectDB function that first configures the client to use the correct URI and check for errors. Secondly, we defined a timeout of 10 seconds we wanted to use when trying to connect. Thirdly, check if there is an error while connecting to the database and cancel the connection if the connecting period exceeds 10 seconds. Finally, we pinged the database to test our connection and returned the client instance.
Create a DB variable instance of the ConnectDB. This will come in handy when creating collections.
Create a GetCollection function to retrieve and create collections on the database.
Next, we need to connect to the database when our application startup. To do this, we need to modify main.go as shown below:


```
package main

import (
    "gin-mongo-api/configs" //add this
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    //run database
    configs.ConnectDB()

    router.Run("localhost:6000")
}

```
Setup API Route Handler and Response Type
Route Handler
With that done, we need to create a user_route.go file inside the routes folder to manage all the user-related routes in our application, as shown below:


```
package routes

import "github.com/gin-gonic/gin"

func UserRoute(router *gin.Engine)  {
    //All routes related to users comes here
}
```

Next, we need to attach the newly created route to the http.Server in main.go by modifying it as shown below:


```
package main

import (
    "gin-mongo-api/configs"
    "gin-mongo-api/routes" //add this
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    //run database
    configs.ConnectDB()

    //routes
    routes.UserRoute(router) //add this

    router.Run("localhost:6000")
}

```
Response Type
Next, we need to create a reusable struct to describe our API’s response. To do this, navigate to the responses folder and in this folder, create a user_response.go file and add the snippet below:


```
package responses

type UserResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}
```

The snippet above creates a UserResponse struct with Status, Message, and Data property to represent the API response type.

PS: json:"status", json:"message", and json:"data" are known as struct tags. Struct tags allow us to attach meta-information to corresponding struct properties. In other words, we use them to reformat the JSON response returned by the API.

Finally, Creating REST API’s
Next, we need a model to represent our application data. To do this, we need to navigate to the models folder, and in this folder, create a user_model.go file and add the snippet below:

```

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Name     string             `json:"name,omitempty" validate:"required"`
    Location string             `json:"location,omitempty" validate:"required"`
    Title    string             `json:"title,omitempty" validate:"required"`
}

```
The snippet above does the following:

Import the required dependencies.
Create a User struct with required properties. We added omitempty and validate:"required" to the struct tag to tell Gin-gonic to ignore empty fields and make the field required, respectively.
Create a User Endpoint
With the model setup, we can now create a function to create a user. To do this, we need to navigate to the controllers folder, and in this folder, create a user_controller.go file and add the snippet below:


```
package controllers

import (
    "context"
    "gin-mongo-api/configs"
    "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        newUser := models.User{
            Id:       primitive.NewObjectID(),
            Name:     user.Name,
            Location: user.Location,
            Title:    user.Title,
        }

        result, err := userCollection.InsertOne(ctx, newUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}
```

The snippet above does the following:

Import the required dependencies.
Create userCollection and validate variables to create a collection and validate models using the github.com/go-playground/validator/v10 library we installed earlier on, respectively.
Create a CreateUser function that returns a Gin-gonic handler. Inside the returned handler, we first defined a timeout of 10 seconds when inserting user into the document, validating both the request body and required field using the validator library. We returned the appropriate message and status code using the UserResponse struct we created earlier. Secondly, we created a newUser variable, inserted it using the userCollection.InsertOne function and check for errors if there are any. Finally, we returned the correct response if the insert was successful.
Next, we need to update user_routes.go with the route API URL and corresponding controller as shown below:


```
package routes

import (
    "gin-mongo-api/controllers" //add this
    "github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    router.POST("/user", controllers.CreateUser()) //add this
}

```
Get a User Endpoint
To get the details of a user, we need to modify user_controller.go as shown below:


```
package controllers

import (
    //other import goes here
    "go.mongodb.org/mongo-driver/bson" //add this
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        //create user code goes here
    }
}

func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        var user models.User
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

        err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
    }
}
```

The snippet above does the following:

Import the required dependencies.
Create a GetAUser function that returns a Gin-gonic handler. Inside the returned handler, we first defined a timeout of 10 seconds when finding a user in the document, a userId variable to get the user’s id from the URL parameter and a user variable. We converted the userId from a string to a primitive.ObjectID type, a BSON type MongoDB uses. Secondly, we searched for the user using the userCollection.FindOne, pass the objId as a filter and use the Decode attribute method to get the corresponding object. Finally, we returned the decoded response.
Next, we need to update user_routes.go with the route API URL and corresponding controller as shown below:

```

package routes

import (
    //import goes here
)

func UserRoute(router *gin.Engine) {
    //other routes goes here
    router.GET("/user/:userId", controllers.GetAUser()) //add this
}

```
PS: We also passed a userId as a parameter to the URL path. The specified parameter must match the one we specified in the controller.

Edit a User Endpoint
To edit a user, we need to modify user_controller.go as shown below:


```
package controllers

import (
    //import goes here
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        //create user code goes here
    }
}

func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        // get a user code goes here
    }
}

func EditAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        var user models.User
        defer cancel()
        objId, _ := primitive.ObjectIDFromHex(userId)

        //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}
        result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //get updated user details
        var updatedUser models.User
        if result.MatchedCount == 1 {
            err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
    }
}
```

The EditAUser function above does the same thing as the CreateUser function. However, we included an update variable to get updated fields and updated the collection using the userCollection.UpdateOne. Lastly, we searched for the updated user’s details and returned the decoded response.

Next, we need to update user_routes.go with the route API URL and corresponding controller as shown below:

```

package routes

import (
    //import goes here
)

func UserRoute(router *gin.Engine) {
    //other routes goes here
    router.PUT("/user/:userId", controllers.EditAUser()) //add this
}
```

Delete a User Endpoint
To delete a user, we need to modify user_controller.go as shown below:


```
package controllers

import (
    //import goes here
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        //create user code goes here
    }
}

func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        // get a user code goes here
    }
}

func EditAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        //edit a user code goes here
    }
}

func DeleteAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

        result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound,
                responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
            )
            return
        }

        c.JSON(http.StatusOK,
            responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
        )
    }
}
```

The DeleteAUser function follows the previous steps by deleting the matched record using the userCollection.DeleteOne. We also checked if an item was successfully deleted and returned the appropriate response.

Next, we need to update user_routes.go with the route API URL and corresponding controller as shown below:


```
package routes

import (
    //import goes here
)

func UserRoute(router *gin.Engine) {
    //other routes goes here
    router.DELETE("/user/:userId", controllers.DeleteAUser()) //add this
}
```

Get List of Users Endpoint
To get the list of users, we need to modify user_controller.go as shown below:


```
package controllers

import (
    //import goes here
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        //create user code goes here
    }
}

func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        // get a user code goes here
    }
}

func EditAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        //edit a user code goes here
    }
}

func DeleteAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        //delete a user code goes here
    }
}

func GetAllUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var users []models.User
        defer cancel()

        results, err := userCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleUser models.User
            if err = results.Decode(&singleUser); err != nil {
                c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }

            users = append(users, singleUser)
        }

        c.JSON(http.StatusOK,
            responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
        )
    }
}
```

The GetAllUsers function follows the previous steps by getting the list of users using the userCollection.Find. We also read the retuned list optimally using the Next attribute method to loop through the returned list of users.

Next, we need to update user_routes.go with the route API URL and corresponding controller as shown below:


```
package routes

import (
    //import goes here
)

func UserRoute(router *gin.Engine) {
    //other routes goes here
    router.GET("/users", controllers.GetAllUsers()) //add this
}
```
