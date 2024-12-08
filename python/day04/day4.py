import re

def read_input(filename):
    with open(filename, 'r') as f:
        return [line.strip() for line in f]

def get_diagonals(matrix):
    rows, cols = len(matrix), len(matrix[0])
    diagonals = []
    for d in range(-(rows - 1), cols):
        diag = ''.join(
            matrix[i][i - d] 
            for i in range(max(0, d), min(rows, cols + d))
        )
        diagonals.append(diag)
    
    for d in range(rows + cols - 1):
        diag = ''.join(
            matrix[i][d - i] 
            for i in range(max(0, d - cols + 1), min(rows, d + 1))
        )
        diagonals.append(diag)
    
    return diagonals

def count_pattern_occurrences(lines, pattern):
    return sum(
        len(re.findall(pattern, ''.join(line) if isinstance(line, tuple) else line)) + 
        len(re.findall(pattern, ''.join(line)[::-1] if isinstance(line, tuple) else line[::-1]))
        for line in lines
    )

def first_part(matrix):
    return count_pattern_occurrences(
        [
            *matrix,  
            *zip(*matrix), 
            *get_diagonals(matrix) 
        ],
        r'XMAS'
    )

def check_x_mas_point(matrix, i, j):
    rows, cols = len(matrix), len(matrix[0])
    if i - 1 < 0 or i + 1 >= rows or j - 1 < 0 or j + 1 >= cols:
        return False

    x_mas_patterns = {"MAS", "SAM"}
    
    l_diag = ''.join([
        matrix[i - 1][j - 1],
        matrix[i][j],
        matrix[i + 1][j + 1]
    ])
    
    r_diag = ''.join([
        matrix[i - 1][j + 1],
        matrix[i][j],
        matrix[i + 1][j - 1]
    ])

    return any(
        (l_diag in x_mas_patterns or l_diag[::-1] in x_mas_patterns) and
        (r_diag in x_mas_patterns or r_diag[::-1] in x_mas_patterns)
        for _ in [None]
    )

def second_part(matrix):
    rows, cols = len(matrix), len(matrix[0])
    return sum(
        check_x_mas_point(matrix, i, j)
        for i in range(1, rows - 1)
        for j in range(1, cols - 1)
    )

matrix = read_input("XMAS_input.txt")
first_part_result = first_part(matrix)
second_part_result = second_part(matrix)

print("First part:",first_part_result)
print("Second part:",second_part_result)