# Pokemon Wiki

## 😎 **Description**

This is a REST API that can do all the CRUD operations like create, update, delete or display pokemons and their masters. These are JSON objects that are saved in a database hosted on a mongoDB Atlas cluster.


## 👨‍🏫 **Quick setup and test**

Clone the repo with https:
```
git clone https://github.com/jcamargoendava/pokemonwiki.git
```
Change directory into the app folder:
```
cd pokemonwiki
```
Start the local server
```
go run main.go
```

Let's quickly create a pokemon. Open up Postman and send a POST request with the following raw body:
```
{
    "PokemonID":write your new pokemon id here,
    "Name":"write your new pokemon name here",
    "Img":"write your new image name here"
}

example:
{
    "PokemonID":35,
    "Name":"PokemonBeast",
    "Img":"BeastImage.jpg"
}
```
to this address:
```
http://localhost:8081/api/pokemon
```
If everything works well, you should receive back a **200** successful message similar to this:
```
{
    "data": {
        "InsertedID": "615edb33bf2a75ecf0479345"
    }
}
```


## 🗄️ **Project Structure** ##
The project code is divided into 6 main folders:
### **Controllers** ###
This is where you have the main logic of the project.
All requests go through the functions included in this package.
The controllers are the connection point between the models and the rest of the application

### **Database** ###
This package is responsible for all the database settings, mainly the database mongoDB connection URI used to connect to the Atlas cluster

### **Models** ###
This is were the object models that populate the database are defined, field by field. Their exact name and type. For a successful request to occur, you need to send objects exactly matching the models defined here.


```
type Pokemon struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	PokemonID int                `json:"pokemon_id" bson:"pokemon_id"`
	Name      string             `json:"name" bson:"name"`
	Img       string             `json:"img" bson:"img"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
```

```
type Master struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Gender    string             `json:"gender" bson:"gender"`
	Age       string             `json:"age" bson:"age"`
	Pokemons  []string           `json:"pokemons" bson:"pokemons"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
```

### **Repository** ###
Package that is responsible for all the database related logic

### **Routes** ###
Package responsible for setting up all the main endpoints of the application

### **Services** ###
Package responsible with storing most of the application's logic

### **Main.go file** ###
It's the main entry point to the entire application





## 🤓 **API Documentation**


### Endpoints
- `GET` &ensp;&ensp;&ensp;&ensp; *`api/pokemons`*  
    Returns a list of all the pokemon objects inside the mongoDB database

- `GET` &ensp;&ensp;&ensp;&ensp; *`api/pokemon/:name`*  
Returns the pokemon object that matches the path parameter name.  
Example: *api/pokemon/Pikachu100*

- `POST`&ensp;&ensp;&ensp;&ensp;*`api/pokemon`*  
Creates a pokemon object and saves it to the database
Since it's a POST request, it sends a request body composed of all the fields of a pokemon.
```
{
    "PokemonID":4326,
    "Name":"Pikachu",
    "Img":"PikachuPicture.jpg"
}
```
Notice there are only 3 fields in the object above and that's because the database id, the createdAt and modifiedAt fields are populated automatically. And the same goes for the master objects.  

- `PUT` &ensp;&ensp;&ensp;&ensp; *`api/pokemon/:id`*  
Updates a pokemon object that matches the id in the path  
If that id doesn't exist, it will create a new empty object with that id
- `DELETE`&ensp; *`api/pokemon/:id`*  
Deletes a pokemon object that matches the id in the path  
- `GET`&ensp;&ensp;&ensp;&ensp;&ensp;*`api/master/:id`*  
Returns a master object that matches the id in the path  
- `POST`&ensp;&ensp;&ensp;&ensp;*`api/master`*  
Creates a master object and saves it to the database.
Since it's a POST request, it sends a request body composed of all the fields of a master.  

```
{
    "Name":"Player1",
	"Gender":"Male",
	"Age":92,
	"Pokemons": [{},{},{},....,{}],
}
```

- `PUT`&ensp;&ensp;&ensp;&ensp;&ensp;*`api/master/:id`*  
Updates a master object that matches the id in the path
- `DELETE` &ensp;*`api/master/:id`*  
Deletes a master object that matches the id in the path
