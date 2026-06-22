# ❌ ⭕ CLI Tic-Tac-Toe with Minimax AI
An interactive Tic-Tac-Toe game played entirely in the terminal, featuring a local 2-player mode and a 1-player mode against an AI. The AI has two difficulty levels: Easy (random moves) and Hard (an unbeatable AI powered by the Minimax algorithm).

## 🚀 Features
* **Game Modes**:
  * **Player vs Player**: Local pass-and-play matches.
  * **Player vs AI**: Solo play against the computer.
* **AI Difficulties**:
  * **Easy**: The AI plays random available squares.
  * **Hard**: The AI uses decision-tree Minimax traversal, making it completely unbeatable. The best outcome you can achieve is a draw.
* **Modern CLI Design**: Rendered using clear grid alignment and ANSI colors for player symbols (`X` in green, `O` in red).
* **Robust Input Validation**: Validates user inputs for valid coordinates and prevents players from overwriting active cells.

## 🛠️ Go Concepts Demonstrated
* **Minimax Decision-Tree Traversal**: A recursive backtracking algorithm evaluating all possible game branches to maximize utility for the AI and minimize utility for the human player.
* **Multidimensional Arrays/Slices**: State tracking via a standard `[3][3]rune` array.
* **Terminal ANSI Graphics**: Building visual layouts inside raw terminal streams using styling codes.
* **Console Reads**: Scanning coordinates from user console lines using `bufio.NewReader` and processing them.

## 📖 How to Run
1. Navigate to this directory:
   ```bash
   cd tic-tac-toe
   ```
2. Run the game:
   ```bash
   go run main.go
   ```
3. Input moves using coordinates of `row col` (e.g. `1 3` for Row 1, Column 3):
   ```
        1     2     3
     +-----+-----+-----+
   1 |     |     |  X  |
     +-----+-----+-----+
   2 |     |     |     |
     +-----+-----+-----+
   3 |     |     |     |
     +-----+-----+-----+
   ```