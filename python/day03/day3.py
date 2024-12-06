import re

def read_file(file_path):
    with open(file_path, 'r') as file:
        return file.read()

def extract_mul_pairs(memory, part2=False):
    pattern = re.compile(r"mul\((\d+),(\d+)\)")
    pairs = []
    if not part2:
        matches = pattern.findall(memory)
        return [(int(x), int(y)) for x, y in matches]
    mul_enabled = True
    idx = 0
    while idx < len(memory):
        if memory[idx:].startswith("do()"):
            mul_enabled = True
            idx += 4
            continue
        elif memory[idx:].startswith("don't()"):
            mul_enabled = False
            idx += 7
            continue
        mul_match = pattern.match(memory, idx)
        if mul_match and mul_enabled:
            x, y = map(int, mul_match.groups())
            pairs.append((x, y))
            idx += len(mul_match.group(0))
        else:
            idx += 1
    return pairs

def calculate_sum(pairs):
    return sum(x * y for x, y in pairs)

def process_file(file_path):
    memory = read_file(file_path)
    part1_pairs = extract_mul_pairs(memory)
    part2_pairs = extract_mul_pairs(memory, part2=True)
    return calculate_sum(part1_pairs), calculate_sum(part2_pairs)

file_path = "corrupted_memory.txt"
result_part1, result_part2 = process_file(file_path)
print("Total sum of valid multiplications part 1:", result_part1, ", part 2:", result_part2)
