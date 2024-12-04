def calculate_total_distance(file_path):
    left_list = []
    right_list = []
    with open(file_path, 'r') as file:
        for line in file:
            values = line.strip().split()  
            if len(values) == 2: 
                left_list.append(int(values[0]))
                right_list.append(int(values[1]))

    left_list.sort()
    right_list.sort()
    
    total_distance = 0
    for left, right in zip(left_list, right_list):
        total_distance += abs(left - right)
    
    return total_distance
    
def calculate_similarity_score(left_list, right_list):
    right_count = {}
    for num in right_list:
        if num in right_count:
            right_count[num] += 1
        else:
            right_count[num] = 1
    similarity_score = 0
    for num in left_list:
        if num in right_count:
            similarity_score += num * right_count[num]
    
    return similarity_score
    
def calculate_similarity_score_from_file(file_path):
    left_list = []
    right_list = []

    with open(file_path, 'r') as file:
        for line in file:
            values = line.strip().split()  
            if len(values) == 2: 
                left_list.append(int(values[0]))
                right_list.append(int(values[1]))
    return calculate_similarity_score(left_list, right_list)

file_path = 'location_data.txt'  
result = calculate_total_distance(file_path)
print(f"The total distance between the lists is: {result}")
similarity_score_from_file = calculate_similarity_score_from_file(file_path)
print(f"The similarity score for the file data is: {similarity_score_from_file}")



