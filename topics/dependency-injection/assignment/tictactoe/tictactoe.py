import re


class Game:
    def __init__(self):
        self.board = [
            [" ", " ", " "],
            [" ", " ", " "],
            [" ", " ", " "],
        ]
        self.turn = "X"

    def play(self):
        while True:
            # Print the current board state
            print("  1 2 3")
            for i in range(3):
                print(f"{i + 1} " + "|".join(self.board[i]))

            # Accept the next move
            move = input(f"{self.turn}'s move, enter row and column: ")
            row, col = re.split(r"[ \t,]+", move)
            row_index, col_index = int(row) - 1, int(col) - 1

            # Validate the move
            if self.board[row_index][col_index] != " ":
                print("Choose an empty space!")
                continue

            # Update the board state
            self.board[row_index][col_index] = self.turn

            # Check for a winner
            winner = False
            configs = [
                [(0, 0), (0, 1), (0, 2)],
                [(1, 0), (1, 1), (1, 2)],
                [(2, 0), (2, 1), (2, 2)],
                [(0, 0), (1, 0), (2, 0)],
                [(0, 1), (1, 1), (2, 1)],
                [(0, 2), (1, 2), (2, 2)],
                [(0, 0), (1, 1), (2, 2)],
                [(2, 0), (1, 1), (0, 2)],
            ]
            for config in configs:
                config_wins = True
                for row_index, col_index in config:
                    if self.board[row_index][col_index] != self.turn:
                        config_wins = False
                        break
                if config_wins:
                    winner = True
                    break
            if winner:
                print(f"Congratulations {self.turn}, you won!")
                break

            # There was no winner, so go to the next player
            if self.turn == "X":
                self.turn = "O"
            else:
                self.turn = "X"


def main():
    game = Game()
    game.play()


if __name__ == "__main__":
    main()
