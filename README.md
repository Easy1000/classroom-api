# Classroom API

Classroom API is a simple Go-based API for managing students, classes, enrollments, and comments. The project uses MySQL as its database and is containerized using Docker and Docker Compose.

## Features

- Manage Students
- Manage Classes
- Manage Enrollments
- Manage Comments
- Seed initial data, including fetching comments from an external API

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed
- [Docker Compose](https://docs.docker.com/compose/install/) installed

## Project Structure

```
.
├── docker-compose.yml   # Docker Compose file for services
├── Dockerfile           # Dockerfile for building the API service
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksum
├── main.go              # Entry point for the API
├── models.go            # Model definitions for database
├── db.go                # Database connection and initialization
├── seed.go              # Data seeding logic
```

## Getting Started

### Clone the Repository

```bash
git clone <repository-url>
cd classroom-api
```

### Build and Start the Services

Run the following command to start the MySQL and API services:

```bash
docker-compose up --build
```

This will:

- Start a MySQL container with a `classroom_db` database.
- Build and run the API container.

The API will be available at `http://localhost:8080`.

### Environment Variables

The following environment variables are used:

| Variable | Description                | Default Value                                                                         |
| -------- | -------------------------- | ------------------------------------------------------------------------------------- |
| `DB_DSN` | Database connection string | `root:password@tcp(mysql:3306)/classroom_db?charset=utf8mb4&parseTime=True&loc=Local` |

## Endpoints

### Students

- **GET** `/students` - Retrieve all students

### Classes

- **GET** `/classes` - Retrieve all classes

### Enrollments

- **GET** `/enrollments` - Retrieve all enrollments

### Comments

- **POST** `/comments` - Create a new comment
- **GET** `/comments` - Retrieve all comments
- **GET** `/comments/:id` - Retrieve a specific comment by ID
- **DELETE** `/comments/:id` - Delete a comment by ID

## Database Seeding

The application seeds initial data for classes, students, enrollments, and comments:

- Initial data for classes and students is defined in `seed.go`.
- Comments are fetched from [JSONPlaceholder](https://jsonplaceholder.typicode.com/comments).

## Health Checks

- MySQL container has a health check to ensure it's running before starting the API service.

## Development

### Running Locally

You can run the API locally without Docker by setting up the environment:

1. Install Go (1.23.5 or higher).
2. Install MySQL and set up the `classroom_db` database.
3. Set the `DB_DSN` environment variable in your shell:

   ```bash
   export DB_DSN="root:password@tcp(127.0.0.1:3306)/classroom_db?charset=utf8mb4&parseTime=True&loc=Local"
   ```

4. Run the application:

   ```bash
   go run .
   ```

### Testing the API

Use tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to test the endpoints:

Example:

```bash
curl http://localhost:8080/students

```

## API Documentation

You can find the complete API documentation here:

👉 [Published API Documentation on Postman](https://documenter.getpostman.com/view/22323395/2sAYQanrab)
