# Lab Overview:

Build an API server that allows users to search for movies, view movie details, and manage a list of favorite movies.
The server will use the external movie API at https://distribution-uat.dev.muangthai.co.th/mtl-node-red/golang-course/movie-api/list for fetching movie data
and PostgreSQL for storing user favorites.

----------------

## Service Requirements:

### Endpoints:

#### Search Movies:

- **GET /movies**
    - Fetches a list of movies matching the search query from the external API.
    - Returns a JSON array of movies with basic information (title, year, ID).

#### Get Movie Details:

- **GET /movies/{id}**
    - Retrieves detailed information about a specific movie using its ID from the external API.
    - Returns a JSON object with detailed movie info (title, year, plot, cast, etc.).

#### Get Favorite Movies:

- **GET /favorites**
    - Retrieves all favorite movies for a specific user from the database.
    - Returns a JSON array of the user’s favorite movies.

#### Add Favorite Movie: (TOBE IMPLEMENTED)

- **POST /favorites**
    - Adds a movie to the user’s list of favorites in the PostgreSQL database.
    - Expects a JSON body with the movie ID.
    - Returns a success message with the added movie’s details.

#### Remove Favorite Movie: (TOBE IMPLEMENTED)

- **DELETE /favorites/{movie_id}**
    - Removes a movie from the user’s list of favorites.
    - Returns a success message confirming deletion.

----------------

## Database Schema:

- **favorite_movies Table:**
    - movie_id (VARCHAR, Primary Key)
    - title (VARCHAR)
    - year (INTEGER)
    - rating (FLOAT)
    - created_at (TIMESTAMP)

----------------

## LAB Requirements:

### Set up PostgreSQL Database using Docker Command:

```bash
docker run -d \
  --name postgres_db \
  --restart always \
  -e POSTGRES_USER=admin_user \
  -e POSTGRES_PASSWORD=admin_password \
  -e POSTGRES_DB=database \
  -p 5432:5432 \
  -v postgres_data:/var/lib/postgresql/data \
  postgres:latest
```

### Implement POST /favorites:

1. Add POST /favorites endpoint to API Server.
2. Implement handler that process the following json body from request body:

```json
{
  "movie_id": "tt33060158"
}
```

3. Implement add favorite service function in interface and implementation.

- Service must validate existence of given movie_id and fetch movie details from external API. (Using GetMovieDetail
  from MovieAPIConnector)
- In case of movie not found, return error.
- In case of movie found, passing movie detail to favorite repository to save in database.

4. Implement add favorite repository function in interface and implementation.

- This function will receive movie details and save the following detail in database.

5. Validate by adding a favorite movie using POST /favorites endpoint.

- Check if the movie is added to the database.
- Check if the movie is not added to the database if the movie_id is invalid.
- Check if the added movie appear in the GET /favorites endpoint.

6. Try to add duplicate favorite movie and check if the service return appropriate error.

### Implement external config

1. Currently, an external API URL, database credentials, are hardcoded in the code.
   Implement a way to read the external API URL from both ENV variable and .env config file.

2. Implement config package to read config from ENV variable and .env file.

3. Try using wire dependency injection to inject config into required app, service, repository.

4. Validate result by update config in .env file and check if the service still works.

### Try running service in Docker container

1. Implement Dockerfile to build and run service.
2. Run service with port mapping 8080:8080
3. Try establish database connection to postgres_db container.
4. Try to add favorite movie using POST /favorites endpoint.

## Extra 1:
### Using docker compose to run both service and database

## Extra 2: 
### Implement DELETE /favorites/{movie_id}:

## Extra 3:
### Implement service unit test using mock