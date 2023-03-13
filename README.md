# goMovies

goMovies is a movie management web application that allows you to create, read, update, and delete information about movies and directors. You can also set a rating and watched status for each movie. The application also includes an authorization system to secure access to the data.

## Features

* View a list of all movies and directors
* View detailed information about each movie and director
* Add, edit, and delete movies and directors
* Set a rating and watched status for each movie
* Authorization system to secure access to the data

## Technologies

* Go programming language
* SQLite database
* HTML/CSS templates
* Bootstrap CSS framework

## Installation

1. Clone the repository:

    `git clone https://github.com/1337yeeee/goMovies.git`

2. Navigate to the project directory:

    `cd goMovies`

3. Install the necessary dependencies:

    `go get`

4. Set up environment variables:

    Create a .env file in the project root directory
    Add the following environment variables to the file:
    ```
    PORT=8080
    DB_FILE_NAME=/path/to/database/file.db
    LOGFILE=/path/to/logfile/file.log
    ```

5. Run the application:

    `go run main.go`

    Access the application at `http://localhost:8080`

## Usage

  * Movies page: Displays a list of all movies and directors
  * Movie page: Displays detailed information about a movie, including the director, release year, rating, and watched status. Allows you to set the rating and watched status.  
  * Director page: Displays detailed information about a director, including the movies they have directed.
