package main

import (
  "fmt"
  "strings"
  "os"
)

func main() {
  player1_wins := 0
  player2_wins := 0
  // counts wins for each player
  for {
    player1_input := getPlayer1Input()
    player2_input := getPlayer2Input()
    // gets input for both players
    winner := decideWinner(player1_input, player2_input)
    // gets the winner
    if (winner == "draw") {
      player1_wins += 1
      player2_wins += 1
      fmt.Println("Draw")
    } else if (winner == "player 1 wins") {
        player1_wins += 1
        fmt.Println("Player 1 wins")
    } else if (winner == "player 2 wins") {
        player2_wins += 1
        fmt.Println("Player 2 wins")
    }
    // adds wins to each player
    var quit string
    // declare variabel to decide to quit
    for {
      fmt.Println("Do you want to quit (y/n)?")
      fmt.Scanf("%s", &quit)
      quit = strings.ToLower(quit)
      // asks if they want to quit
      if (quit != "") {
        // if it's not an empty string
        if (quit == "n") {
          break
        } else if (quit == "y") {
            fmt.Printf("Player 1 wins: %d\n", player1_wins)
            fmt.Printf("Player 2 wins: %d\n", player2_wins)
            os.Exit(0)
        }
      }
    }
  }
}

func getPlayer1Input() string {
  var player1_input string
  // variable to store player 1s input
  for {
    fmt.Println("player 1 enter rock, paper, scissors: ")
    fmt.Scanf("%s", &player1_input)
    player1_input = strings.ToLower(player1_input)
    // gets player 1s chocie

    if (player1_input != "") {
      // if they actually enter something
      if (player1_input == "rock") {
        break
      } else if (player1_input == "paper") {
          break
      } else if (player1_input == "scissors") {
          break
      } else {
          continue
      }
      // verify input
    } else {
        continue
      // if they don't enter anything restyart loop and try again
    }
 }
 // gets input for player 1 and verifies that that input is correct
 fmt.Printf("%s\n", strings.Repeat("\n\t\t*************\t\t\n", 25))
 return player1_input
}

func getPlayer2Input() string {
  var player2_input string
  // variable to store player 1s input
  for {
    fmt.Println("player 2 enter rock, paper, scissors: ")
    fmt.Scanf("%s", &player2_input)
    player2_input = strings.ToLower(player2_input)
    // gets player 1s chocie

    if (player2_input != "") {
      // if not a zero value
      if (player2_input == "rock") {
        break
      } else if (player2_input == "paper") {
          break
      } else if (player2_input == "scissors") {
          break
      } else {
          continue
      }
      // verify input
    } else {
        continue
      // if they don't enter anything get input
    }
 }
 // gets input for player 1 and verifies that that input is correct
 fmt.Printf("%s\n", strings.Repeat("\n\t\t*************\t\t\n", 25))
 return player2_input
}

func decideWinner(player1_input string, player2_input string) string {
  if (player1_input == player2_input) {
    // if it is a draw
    return "draw"
  }

  if (player1_input == "rock") {
    if (player2_input == "paper") {
      return "player 2 wins"
    } else if (player2_input == "scissors") {
        return "player 1 wins"
    }
  } else if (player1_input == "paper") {
      if (player2_input == "scissors") {
        return "player 2 wins"
    } else if (player2_input == "rock") {
        return "player 1 wins"
    }
  } else if (player1_input == "scissors") {
      if (player2_input == "rock") {
        return "player 2 wins"
      } else if (player2_input == "paper") {
        return "player 1 wins"
      }
  }
  return "error"
}
