#!/bin/bash
input="input.txt"

grep -oP "((?<=mul\()\d{1,3},\d{1,3}(?=\))|do\(\)|don't\(\))" $input > output.txt

declare -i result

mul_operation() {
    (( result += ${1%,*} * ${1#*,} ))
    return 0
}

is_enabled=1

# Usage: calculate_sum <part2:true/false>
calculate_sum() {
    result=0

    while IFS=' ' read -r line; do
        if [[ $line == "do()" ]]; then
            is_enabled=1
        elif [[ $line == "don't()" ]]; then
            is_enabled=0
        elif [[ $is_enabled -eq 1 || $1 == "false" ]]; then
            mul_operation "$line"
        fi
    done < "output.txt"

    echo $result
}

echo "Result of the multiplications: $(calculate_sum false)"
echo "Result of the multiplications (switching): $(calculate_sum)"

rm output.txt
