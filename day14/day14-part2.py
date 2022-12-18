def read_from_file(file_path):
    with open(file_path, "r") as f:
        lines = f.readlines()

        segments = []
        maxx = 0
        maxy = 0

        rock_coords = []
        for line in lines:
            points = line.split("->")
            for idx in range(0, len(points) - 1):
                p1 = points[idx]
                p2 = points[idx + 1]
                p1 = p1.strip()
                p2 = p2.strip()

                y1, x1 = p1.split(",")
                y1, x1 = int(y1), int(x1)
                y2, x2 = p2.split(",")
                y2, x2 = int(y2), int(x2)

                x1, y1 = y1, x1
                x2, y2 = y2, x2

                if x1 > maxx:
                    maxx = x1
                if x2 > maxx:
                    maxx = x2

                if y1 > maxy:
                    maxy = y1
                if y2 > maxy:
                    maxy = y2

                if x1 == x2:
                    for y in range(min(y1, y2), max(y1, y2) + 1):
                        rock_coords.append((x1, y))
                else:
                    for x in range(min(x1, x2), max(x1, x2) + 1):
                        rock_coords.append((x, y1))

                segments.append((x1, y1, x2, y2))

        matrix = []
        for y in range(maxy + 3):
            line = []
            for x in range(2 * maxx + 1):
                line.append(0)
            matrix.append(line)

        for x, y in rock_coords:
            matrix[y][x] = 1

        matrix[len(matrix) - 1] = [1] * len(matrix[0])

        # sand source
        matrix[0][500] = 2
        return matrix, segments


def print_matrix(matrix):
    line = ""
    for r in range(len(matrix)):
        for c in range(len(matrix[0])):
            if matrix[r][c] == 1:
                line += "#"
            elif matrix[r][c] == 2:
                line += "+"
            elif matrix[r][c] == 3:
                line += "O"
            else:
                line += "."
        line += "\n"
    return line


def valid_coords(matrix, x, y):
    return x >= 0 and x < len(matrix[1]) and y >= 0 and y < len(matrix)


def solve(matrix):
    it = 0
    while True:
        start_x, start_y = 500, 0
        not_settled = True
        while not_settled:
            not_settled = False
            candidates = [
                (start_x, start_y + 1),
                (start_x - 1, start_y + 1),
                (start_x + 1, start_y + 1),
            ]

            for candidate in candidates:
                x, y = candidate
                if valid_coords(matrix, x, y) and matrix[y][x] == 0:
                    start_x, start_y = x, y
                    not_settled = True
                    break
        if start_x == 500 and start_y == 0:
            break
        matrix[start_y][start_x] = 3
        it += 1

    return matrix, it


matrix, segments = read_from_file("./in-day14.txt")
print("Start solve...")
matrix, it = solve(matrix)

print(it + 1)

with open("./test.txt", "w+") as f:
    f.write(print_matrix(matrix))
