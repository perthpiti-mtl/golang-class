# Lab Overview:

Build an API server that allows users to search for movies, view movie details, and manage a list of favorite movies. The server will use the external movie API at https://imdb.iamidiotareyoutoo.com/docs/index.html for fetching movie data and PostgreSQL for storing user favorites.

----------------
## Service Requirements:

### Endpoints:
#### Search Movies:
- **GET /movies/search?query={query}**
  - Fetches a list of movies matching the search query from the external API. 
  - Returns a JSON array of movies with basic information (title, year, ID).
  

#### Get Movie Details:
- **GET /movies/{id}**
  - Retrieves detailed information about a specific movie using its ID from the external API.
  - Returns a JSON object with detailed movie info (title, year, plot, cast, etc.).

#### Add Favorite Movie:
- **POST /favorites**
  - Adds a movie to the user’s list of favorites in the PostgreSQL database.
  - Expects a JSON body with the movie ID.
  - Returns a success message with the added movie’s details.

#### Get Favorite Movies:
- **GET /favorites**
  - Retrieves all favorite movies for a specific user from the database.
  - Returns a JSON array of the user’s favorite movies.

#### Remove Favorite Movie:
- **DELETE /favorites/{movie_id}**
  - Removes a movie from the user’s list of favorites.
  - Returns a success message confirming deletion.
----------------
    
## Database Schema:
- **favorite_movies Table:**
  - movie_id (VARCHAR, Primary Key)
  - title (VARCHAR)
  - year (INTEGER)
  - image (VARCHAR)
  - created_at (TIMESTAMP)

----------------
## Functionality Requirements:

- **External API Integration:**
  - Use the external movie API to fetch movie data.
  - Handle API authentication if required.
  - Implement error handling for API failures (e.g., retries, fallback messages).
- **Database Interaction:**
  - Connect to a PostgreSQL database.
  - Perform CRUD operations on the favorites table.
- **Data Models:**
  - Define Go structs for movie data and user favorites.
  - Use JSON tags for proper marshaling/unmarshaling.
- **Middleware and Error Handling:**
  - Implement logging middleware to log requests and responses.
  - Use recovery middleware to handle panics and return appropriate HTTP responses.
  - Validate incoming request data and parameters.
  - Return meaningful HTTP status codes (e.g., 200 OK, 201 Created, 400 Bad Request, 404 Not Found).
- **Configuration Management:**
  - Use environment variables or a config file to manage sensitive data (e.g., database credentials, API keys).
  - Implement proper configuration loading in the application.