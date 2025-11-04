package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Card struct {
	cardType  string
	cardValue int
}

var typesOfCards = []Card{
	{"2", 2},
	{"3", 3},
	{"4", 4},
	{"5", 5},
	{"6", 6},
	{"7", 7},
	{"8", 8},
	{"9", 9},
	{"10", 10},
	{"J", 10},
	{"Q", 10},
	{"K", 10},
	{"A", 11},
}

func dealCard(cardHand []Card) []Card {
	newCard := typesOfCards[rand.Intn(len(typesOfCards))]
	cardHand = append(cardHand, newCard)
	time.Sleep(time.Second)
	return cardHand
}

func checkAce(cardHand []Card) {
	var aceCounter int
	for i := range len(cardHand) {
		if cardHand[i].cardType == "A" {
			cardHand[i].cardValue = 1
			aceCounter++
		}
	}
	if aceCounter > 1 {
		fmt.Println(calcScore(cardHand))
	}
}

func calcScore(cardHand []Card) int {
	var score int
	for i := range len(cardHand) {
		score += cardHand[i].cardValue
	}
	if score > 21 {
		checkAce(cardHand)
	}
	return score
}

func dealerTurn(cardHand []Card) int {
	var dealerScore int
	cardHand = dealCard(cardHand)
	dealerScore = calcScore(cardHand)
	fmt.Println("Dealer Hand:", cardHand, "(", dealerScore, ")")
	for dealerScore < 17 {
		cardHand = dealCard(cardHand)
		dealerScore = calcScore(cardHand)
		fmt.Println("Dealer hand: ", cardHand, "(", dealerScore, ")")
		if dealerScore > 21 {
			checkAce(cardHand)
		}
	}
	return dealerScore
}

func main() {
	fmt.Println("BlackJack is beginning...")
	gameOn := true
	var dealerHand []Card
	var playerHand []Card
	playerScore := 0
	dealerScore := 0
	var playerMove string

	playerHand = dealCard(playerHand)
	dealerHand = dealCard(dealerHand)
	playerHand = dealCard(playerHand)
	playerScore = calcScore(playerHand)

	fmt.Println("Dealer Showing:", dealerHand[0].cardType, "\nPlayer Hand:", playerHand, "(", playerScore, ")")
	for gameOn {
		fmt.Println("Hit or Stay?")
		fmt.Scan(&playerMove)
		if playerMove == "Hit" {
			playerHand = dealCard(playerHand)
			playerScore = calcScore(playerHand)
			if playerScore > 21 {
				checkAce(playerHand)
				playerScore = calcScore(playerHand)
				if playerScore > 21 {
					fmt.Printf("\n\nPlayer loses (%d). Dealer wins (%d).\n\n", playerScore, dealerScore)
					os.Exit(1)
				}
			}
			fmt.Println("Dealer Showing:", dealerHand[0].cardType, "\nPlayer Hand:", playerHand, "(", playerScore, ")")
		} else if playerMove == "Stay" {
			fmt.Println("You are staying. Dealer's turn...")
			dealerScore = dealerTurn(dealerHand)
			if playerScore > dealerScore {
				fmt.Printf("\n\nPlayer wins (%d)! Dealer loses (%d).\n\n", playerScore, dealerScore)
				gameOn = false
			} else if playerScore == dealerScore {
				fmt.Printf("\n\nTie! Player: (%d) Dealer: (%d).\n\n", playerScore, dealerScore)
				gameOn = false
			} else {
				fmt.Printf("\n\nPlayer loses (%d). Dealer wins (%d).\n\n", playerScore, dealerScore)
				gameOn = false
			}
		} else {
			fmt.Println("Please enter a valid decision (Hit or Stay)...")
		}
	}
}
