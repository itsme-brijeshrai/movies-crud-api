# Movie API Service

## Overview

This project is a simple movie management service written in Go. It allows users to perform CRUD operations on a movie database. The project uses the [Gorilla Mux](https://github.com/gorilla/mux) library for routing.

## Endpoints

The API service has the following endpoints:

- `GET /movies`: Fetches all the movies in the database.
- `GET /movies/{id}`: Fetches a specific movie based on its ID.
- `POST /movies/create`: Creates a new movie in the database.
- `PUT /movies/{id}`: Updates the information of a specific movie based on its ID.
- `DELETE /movies/{id}`: Deletes a specific movie based on its ID.

## How to Run

1. Make sure you have Go installed on your system. You can download it from the [official website](https://golang.org/dl/).
2. Install the Gorilla Mux library by running `go get -u github.com/gorilla/mux`.
3. Download or clone this repository to your local machine.
4. Navigate to the project directory and run `go run main.go`.
5. The server will start on port 8000. You can access the API endpoints via `http://localhost:8000`.

## How to Use

You can use a tool like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to interact with the API.

Example usage:

- Fetch all movies: `GET http://localhost:8000/movies`
- Fetch a specific movie: `GET http://localhost:8000/movies/1`
- Create a new movie: `POST http://localhost:8000/movies/create`
- Update a movie: `PUT http://localhost:8000/movies/1`
- Delete a movie: `DELETE http://localhost:8000/movies/1`

## Data Model

The data model for the Movie entity is defined in a struct named `Movie`. Each movie has the following fields:

- `ID`: A unique identifier for the movie.
- `Title`: The title of the movie.
- `Director`: The director of the movie.
- `Rating`: The rating of the movie.

The movie data is stored in a slice named `movies`.

## Note

Please note that this project uses an in-memory database to store the movies. The data will be lost when the server is restarted.

## Future Work

Here are some ideas for future improvements to the project:

- Add validation to the movie creation and update endpoints.
- Implement pagination for the `GET /movies` endpoint.
- Use a persistent database like PostgreSQL or MySQL to store the movies.
- Add authentication and authorization to the API.
