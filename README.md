**Fiber REST API Example**

This project is a simple RESTful API built with Go using the Fiber web framework and GORM ORM with SQLite as the database. It provides basic CRUD operations for a Todo resource.

**Features**

Create, read, update, and delete Todo items

Uses SQLite for persistent storage

Clean project structure with separation of concerns

**Project Structure**
```
.
├── main.go                # Application entry point and route definitions
├── go.mod                 # Go module definition
├── myDatabase.db          # SQLite database file (created at runtime)
├── dal/                   # Data Access Layer (DAL) for database operations
│   └── todo.go
├── database/              # Database connection logic
│   └── database.go
├── services/              # Service layer for business logic
│   └── todo.go
└── types/                 # Type definitions (DTOs)
    └── todo.go
```

**Getting Started**

**Installation**

1. **Clone the repository**:
```sh
git clone <https://github.com/pehlivanyunuscan/rest-api-fiber>

cd rest-api-fiber
```
2. **Install dependencies**:
```sh
go mod tidy
```
3. **Run the application**:
```sh
go run main.go
```

The server will start at `http://localhost:3000`.

**API Endpoints**

| Method | Endpoint            | Description                |
|--------|---------------------|----------------------------|
| GET    | `/`                 | Welcome message            |
| POST   | `/todos`            | Create a new Todo          |
| GET    | `/todos`            | List all Todos             |
| GET    | `/todos/:todoID`    | Get a Todo by ID           |
| PUT    | `/todos/:todoID`    | Update a Todo by ID        |
| DELETE | `/todos/:todoID`    | Delete a Todo by ID        |

**Example Todo JSON**
```json
{
  "title": "Buy groceries"
}
```
**Example Usage**

1.**Create a Todo**:
```sh
curl -X POST http://localhost:3000/todos \
  
  -H "Content-Type: application/json" \
  
  -d '{"title":"Buy groceries"}'
```
2. **Get all Todos**:
```sh
curl `http://localhost:3000/todos`
```
**Using Postman**

You can also use [Postman](https://www.postman.com/) to interact with the API:

1. **Start your server:**  

   Make sure the API is running at `http://localhost:3000`.

3. **Create a Todo:**

   - Set method to `POST`
   
   - URL: `http://localhost:3000/todos`
   
   -  Body: Select `raw` and `JSON`, then enter:

     ```json
     {
       "title": "Buy groceries"
     }
     ```

5. **Get all Todos:**  
   
   - Set method to `GET`
   
   - URL: `http://localhost:3000/todos`

6. **Get, Update, or Delete a Todo by ID:**  
   
   - Use the endpoints `/todos/:todoID` with the appropriate HTTP method (`GET`, `PUT`, or `DELETE`).
   
   - For `PUT`, provide a JSON body with the updated title.




