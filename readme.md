# Octoscript Backend Developer Intern Technical Test

A RESTful API for managing social media post tasks built with Go, Gin Framework, GORM, PostgreSQL, and Swagger Documentation.

## Prequisites

- Go 1.24 or later
- PostgreSQL 17 or later


## Database Implementation

The project used GORM, a Go ORM for database operations. The database schema includes table for posts. 

The `database/connect.go` file sets up the database connection.

The `model` package defines the structs for the posts entity and GORM uses these structs to map between the Go code and the database.

## REST API Design

The project provides a RESTful API for posts. The API follows standard REST conventions:

### Posts API Routes

- `GET /api/posts` - Get all posts.
- `POST /api/posts` - Add or create a post.
- `GET /api/posts/:id` - Get a post by id.
- `PUT /api/posts/:id` - Edit or update a post by id.
- `DELETE /api/posts/:id` - Delete a post by id.


## Getting Started

1. Clone the repository:

```sh
git clone https://github.com/ahmadammarm/octoscript-intern-test.git
```

2. Navigate to the project directory:

```sh
cd octoscript-intern-test
```

3. Install the project dependencies:

```sh
go mod download
```

4. Configure database connection in database/connect.go:

```sh
source := "host=localhost user=your_username password=your_password dbname=your_dbname port=5432"
```

5. Generate Swagger documentation:

```sh
swag init
```

6. Run the project:

```sh
go run main.go
```

The project will be available at:

`http://localhost:8080`

After running the application, you can access the Swagger documentation at:

`http://localhost:8080/swagger/index.html`
