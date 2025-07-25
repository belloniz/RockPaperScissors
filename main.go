package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestPlayer struct {
	PlayerPlay string `json:"player_play"`
}

type ResponseAPI struct {
	APIPlay    string `json:"api_play"`
	PlayerPlay string `json:"player_play"`
	Winner     string `json:"winner"`
}

var possiblePlays []string

func main() {
	possiblePlays = []string{"pedra", "papel", "tesoura"}

	router := gin.Default()
	router.POST("/play", playRockPaperScissors)

	router.Run("localhost:4000")
}

func playerInputIsValid(playerInput string) bool {
	for _, play := range possiblePlays {
		if play == playerInput {
			return true
		}
	}
	return false
}

func checkWinner(playerInput string, apiInput string) string {
	playerWin := "player"
	apiWin := "api"

	if playerInput == apiInput {
		return "draw"
	}

	switch apiInput {
	case "pedra":
		if playerInput == "tesoura" {
			return apiWin
		}
		return playerWin
	case "papel":
		if playerInput == "pedra" {
			return apiWin
		}
		return playerWin
	case "tesoura":
		if playerInput == "papel" {
			return apiWin
		}
		return playerWin
	default:
		return "error checking winner"
	}

}

func playRockPaperScissors(c *gin.Context) {
	var requestPlayer RequestPlayer

	err := c.BindJSON(&requestPlayer)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	if !playerInputIsValid(requestPlayer.PlayerPlay) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid player input"})
		return
	}

	apiPlay := possiblePlays[rand.Intn(len(possiblePlays))]

	c.IndentedJSON(
		http.StatusOK,
		ResponseAPI{
			APIPlay:    apiPlay,
			PlayerPlay: requestPlayer.PlayerPlay,
			Winner:     checkWinner(requestPlayer.PlayerPlay, apiPlay),
		},
	)
	return
}
