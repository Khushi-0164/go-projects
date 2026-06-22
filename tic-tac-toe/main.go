package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ANSI colors for clean visuals
const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGray   = "\033[90m"
)

type Board [3][3]rune

const (
	empty = ' '
	runeX = 'X'
	runeO = 'O'
)

type Move struct {
	row, col int
}

func main() {
	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)

	printHeader()

	for {
		playGame(reader)

		fmt.Printf("\n%sDo you want to play again? (y/n): %s", colorYellow, colorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			fmt.Printf("\n%sThanks for playing! Goodbye! 👋%s\n\n", colorCyan, colorReset)
			break
		}
	}
}
func printHeader() {
	fmt.Printf("%s===========================================%s\n", colorCyan, colorReset)
	fmt.Printf("%s%s         ❌ ⭕  CLI TIC-TAC-TOE  ⭕ ❌         %s\n", colorCyan, colorBold, colorReset)
	fmt.Printf("%s===========================================%s\n", colorCyan, colorReset)
}
func playGame(reader *bufio.Reader) {
	// Choose Game Mode
	fmt.Printf("\n%sChoose Game Mode:%s\n", colorBold, colorReset)
	fmt.Printf("1. Player vs Player\n")
	fmt.Printf("2. Player vs AI\n")
	mode := getChoice(reader, 1, 2, "Select mode (1-2): ")
	var aiDifficulty int
	var playerRune rune
	var aiRune rune
	if mode == 2 {
		// Choose Difficulty
		fmt.Printf("\n%sChoose AI Difficulty:%s\n", colorBold, colorReset)
		fmt.Printf("1. Easy (Random Moves)\n")
		fmt.Printf("2. Hard (Unbeatable Minimax)\n")
		aiDifficulty = getChoice(reader, 1, 2, "Select difficulty (1-2): ")
		// Choose Symbol
		fmt.Printf("\n%sChoose your Symbol:%s\n", colorBold, colorReset)
		fmt.Printf("1. Play as X (Goes First)\n")
		fmt.Printf("2. Play as O (Goes Second)\n")
		symbolChoice := getChoice(reader, 1, 2, "Select symbol (1-2): ")
		if symbolChoice == 1 {
			playerRune = runeX
			aiRune = runeO
		} else {
			playerRune = runeO
			aiRune = runeX
		}
	}
	// Initialize board
	var board Board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = empty
		}
	}
	currentTurn := runeX // X always starts
	fmt.Printf("\n%sGame started! %s%s\n", colorGreen, getTurnLabel(currentTurn), colorReset)
	for {
		printBoard(board)
		// Check if game is over before making a move
		if checkWin(board, runeX) {
			printWinner(runeX, mode, playerRune)
			break
		}
		if checkWin(board, runeO) {
			printWinner(runeO, mode, playerRune)
			break
		}
		if !isMovesLeft(board) {
			fmt.Printf("\n%s🎉 It's a Draw! Excellent game.%s\n", colorYellow, colorReset)
			break
		}
		// Handle moves
		if mode == 1 {
			// Player vs Player
			fmt.Printf("\n%s's turn. ", getTurnLabel(currentTurn))
			row, col := getPlayerMove(reader, board)
			board[row][col] = currentTurn
			currentTurn = nextTurn(currentTurn)
		} else {
			// Player vs AI
			if currentTurn == playerRune {
				fmt.Printf("\n%sYour turn (%s). ", getTurnLabel(playerRune), string(playerRune))
				row, col := getPlayerMove(reader, board)
				board[row][col] = playerRune
				currentTurn = aiRune
			} else {
				fmt.Printf("\n🤖 AI (%s) is thinking...\n", string(aiRune))
				time.Sleep(600 * time.Millisecond) // Short delay to feel organic
				var move Move
				if aiDifficulty == 1 {
					move = getRandomMove(board)
				} else {
					move = findBestMove(board, aiRune, playerRune)
				}
				board[move.row][move.col] = aiRune
				fmt.Printf("🤖 AI played at Row %d, Col %d.\n", move.row+1, move.col+1)
				currentTurn = playerRune
			}
		}
	}

	// Show final board
	printBoard(board)
}
func getChoice(reader *bufio.Reader, min, max int, prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err == nil && val >= min && val <= max {
			return val
		}
		fmt.Printf("%sInvalid choice. Please enter a number between %d and %d.%s\n", colorRed, min, max, colorReset)
	}
}
func printBoard(b Board) {
	fmt.Println()
	fmt.Printf("     %s1%s     %s2%s     %s3%s\n", colorGray, colorReset, colorGray, colorReset, colorGray, colorReset)
	fmt.Printf("  %s+-----+-----+-----+%s\n", colorCyan, colorReset)
	for i := 0; i < 3; i++ {
		fmt.Printf("%s%d%s ", colorGray, i+1, colorReset)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s|%s  %s  ", colorCyan, colorReset, getRuneString(b[i][j]))
		}
		fmt.Printf("%s|%s\n", colorCyan, colorReset)
		fmt.Printf("  %s+-----+-----+-----+%s\n", colorCyan, colorReset)
	}
}
func getRuneString(r rune) string {
	switch r {
	case runeX:
		return colorGreen + colorBold + "X" + colorReset
	case runeO:
		return colorRed + colorBold + "O" + colorReset
	default:
		return " "
	}
}
func getTurnLabel(r rune) string {
	if r == runeX {
		return colorGreen + colorBold + "Player X" + colorReset
	}
	return colorRed + colorBold + "Player O" + colorReset
}
func nextTurn(current rune) rune {
	if current == runeX {
		return runeO
	}
	return runeX
}
func getPlayerMove(reader *bufio.Reader, b Board) (int, int) {
	for {
		fmt.Print("Enter move (row col, e.g. '1 3'): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Printf("%sInvalid format. Enter row and column separated by a space.%s\n", colorRed, colorReset)
			continue
		}
		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Printf("%sInvalid coordinates. Row and Col must be between 1 and 3.%s\n", colorRed, colorReset)
			continue
		}
		// Convert to 0-indexed
		r, c := row-1, col-1
		if b[r][c] != empty {
			fmt.Printf("%sCell already occupied! Choose an empty spot.%s\n", colorRed, colorReset)
			continue
		}
		return r, c
	}
}
func checkWin(b Board, r rune) bool {
	// Rows
	for i := 0; i < 3; i++ {
		if b[i][0] == r && b[i][1] == r && b[i][2] == r {
			return true
		}
	}
	// Columns
	for i := 0; i < 3; i++ {
		if b[0][i] == r && b[1][i] == r && b[2][i] == r {
			return true
		}
	}
	// Diagonals
	if b[0][0] == r && b[1][1] == r && b[2][2] == r {
		return true
	}
	if b[0][2] == r && b[1][1] == r && b[2][0] == r {
		return true
	}
	return false
}
func isMovesLeft(b Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == empty {
				return true
			}
		}
	}
	return false
}
func printWinner(winner rune, mode int, playerRune rune) {
	if mode == 1 {
		fmt.Printf("\n🏆 %s wins the game! Congratulations! 🎉\n", getTurnLabel(winner))
	} else {
		if winner == playerRune {
			fmt.Printf("\n🏆 %sYou Win! You beat the AI! 🎉%s\n", colorGreen, colorReset)
		} else {
			fmt.Printf("\n💀 %sAI Wins! Better luck next time.%s\n", colorRed, colorReset)
		}
	}
}

// AI - Easy Mode: Random selection
func getRandomMove(b Board) Move {
	var emptyCells []Move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == empty {
				emptyCells = append(emptyCells, Move{row: i, col: j})
			}
		}
	}
	return emptyCells[rand.Intn(len(emptyCells))]
}

// AI - Hard Mode: Minimax Algorithm
func evaluate(b Board, ai, human rune) int {
	// Check rows
	for row := 0; row < 3; row++ {
		if b[row][0] == b[row][1] && b[row][1] == b[row][2] {
			if b[row][0] == ai {
				return +10
			} else if b[row][0] == human {
				return -10
			}
		}
	}
	// Check columns
	for col := 0; col < 3; col++ {
		if b[0][col] == b[1][col] && b[1][col] == b[2][col] {
			if b[0][col] == ai {
				return +10
			} else if b[0][col] == human {
				return -10
			}
		}
	}
	// Check diagonals
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		if b[0][0] == ai {
			return +10
		} else if b[0][0] == human {
			return -10
		}
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		if b[0][2] == ai {
			return +10
		} else if b[0][2] == human {
			return -10
		}
	}
	return 0
}
func minimax(b Board, depth int, isMax bool, ai, human rune) int {
	score := evaluate(b, ai, human)

	// If AI (Maximizer) won
	if score == 10 {
		return score - depth
	}
	// If Human (Minimizer) won
	if score == -10 {
		return score + depth
	}
	// If no more moves, it's a draw
	if !isMovesLeft(b) {
		return 0
	}
	if isMax {
		best := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if b[i][j] == empty {
					b[i][j] = ai
					val := minimax(b, depth+1, false, ai, human)
					if val > best {
						best = val
					}
					b[i][j] = empty
				}
			}
		}
		return best
	} else {
		best := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if b[i][j] == empty {
					b[i][j] = human
					val := minimax(b, depth+1, true, ai, human)
					if val < best {
						best = val
					}
					b[i][j] = empty
				}
			}
		}
		return best
	}
}
func findBestMove(b Board, ai, human rune) Move {
	bestVal := -1000
	bestMove := Move{row: -1, col: -1}
	// Optimize search: if board is empty, pick the center immediately
	emptyCount := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == empty {
				emptyCount++
			}
		}
	}
	if emptyCount == 9 {
		return Move{row: 1, col: 1} // Center is mathematically the best first move
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == empty {
				b[i][j] = ai
				moveVal := minimax(b, 0, false, ai, human)
				b[i][j] = empty
				if moveVal > bestVal {
					bestMove.row = i
					bestMove.col = j
					bestVal = moveVal
				}
			}
		}
	}
	return bestMove
}
