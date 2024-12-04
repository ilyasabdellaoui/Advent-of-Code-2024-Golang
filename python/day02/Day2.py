
def is_safe(report):
    diffs = [report[i+1] - report[i] for i in range(len(report) - 1)]
    # Check all differences are between -3 and -1 
    # or between 1 and 3
    if all(-3 <= diff <= -1 for diff in diffs) or all(1 <= diff <= 3 for diff in diffs):
        return True
    return False
    
def count_safe_reports(file_path):
    safe_count = 0
    with open(file_path, 'r') as file:
        for line in file:
            report = list(map(int, line.split()))
            if is_safe(report):
                safe_count += 1
    return safe_count

def is_safe_with_dampener(report):
    if is_safe(report):
        return True
    # the same rules apply as before, except if removing a single level from an unsafe report would make it safe
    # remove each level and check if safe
    for i in range(len(report)):
        modified_report = report[:i] + report[i+1:]
        if is_safe(modified_report):
            return True
    return False

def count_safe_reports_with_dampener(file_path):
    safe_count = 0
    with open(file_path, 'r') as file:
        for line in file:
            report = list(map(int, line.split()))
            if is_safe_with_dampener(report):
                safe_count += 1
    return safe_count

file_path = 'reports.txt'

safe_reports = count_safe_reports(file_path)
safe_reports_with_dampener = count_safe_reports_with_dampener(file_path)

print(f"Number of safe reports (Part 1): {safe_reports}")
print(f"Number of safe reports with tolerence (Part 2): {safe_reports_with_dampener}")
