with open("input.txt", "r") as f:
    matrix = [[val for val in line] for line in f.read().split("\n")]

MATRIX_X_SIZE = len(matrix)
MATRIX_Y_SIZE = len(matrix[0])
SEARCH_WORD = "XMAS"


def in_matrix(row, col):
    if row >= 0 and row < MATRIX_X_SIZE and col >= 0 and col < MATRIX_Y_SIZE:
        return True
    return False


def get_neighbour(row, col):
    return ((i, j) for i in (-1, 0, 1) for j in (-1, 0, 1))


def discover(word, row, col, direction):
    word = "".join(word + matrix[row][col])
    if SEARCH_WORD.startswith(word):
        if len(word) == len(SEARCH_WORD):
            return True
        else:
            if in_matrix(row + direction[0], col + direction[1]):
                return discover(word, row + direction[0], col + direction[1], direction)
    return False


number_of_finds = 0
for row in range(MATRIX_X_SIZE):
    for col in range(MATRIX_Y_SIZE):
        for direction in get_neighbour(row, col):
            if in_matrix(row + direction[0], col + direction[1]):
                if discover(
                    matrix[row][col], row + direction[0], col + direction[1], direction
                ):
                    number_of_finds += 1
print(number_of_finds)
assert number_of_finds == 2633
