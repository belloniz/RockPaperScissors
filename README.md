# Rock Paper Scissors API

This application is a simulation of the Rock, Paper, Scissors game using a simple HTTP API. The user sends a request with their play, and the server randomly chooses its own play and returns the result of the match.

It includes the following features:

- Accept a player move via POST request;
- Randomly generate a move for the API;
- Return the API move, the player's move, and the winner;
- Validate the input and return an error message for invalid moves.

## Technologies

This app was developed using:

- Go version 1.21+
- Gin Web Framework

## How to run this application

You need to have [Go installed](https://golang.org/doc/install).

1. Clone this repository by running the following in your terminal:
<pre>git clone git@github.com:your-user/rockpaperscissors.git</pre>

2. Access the app folder:
<pre>cd rockpaperscissors</pre>

3. Run the application:
<pre>go run main.go</pre>

Your terminal should display something like this:
<pre>[GIN-debug] Listening and serving HTTP on localhost:4000</pre>

## How to use the API
Since the application is running locally, we will use http://localhost:4000 as the base url. Using an API client like Postman, Insomnia, etc, you can use the following route as an example:

### POST /cart/
This is an example of the parameters to be sent in the request's body:
```
{
    "player_play": "papel"
}
```

If the request is sucessfull, it will return a status code of `200` and the game result.

<img width="306" height="414" alt="image" src="https://github.com/user-attachments/assets/b2f57a35-1fdb-4a03-a267-0a37d7df944a" />

If the player input is invalid, it will return an error.

<img width="324" height="378" alt="image" src="https://github.com/user-attachments/assets/a9e79a3a-6c15-4c9f-84a0-c824ac8e6e75" />

