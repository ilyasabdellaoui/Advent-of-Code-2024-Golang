#!/bin/bash
input="input.txt"
safe_c=0
t_safe_c=0

check_report() {
  IFS=' ' read -a arr <<< "$1"
  local len=${#arr[@]}

  if [ "$len" -le 1 ]; then return 0; fi
  
  i=0
  while [[ ($i -lt $((len-1))) ]]; do
    diff=$((arr[i+1] - arr[i]))
    # Determine direction
    if [[ $i -eq 0 ]]; then
      (( sign = diff > 0 ? 1 : -1 ))
    fi
    (( diff *= sign))
    if [[ ($diff -le 0) || ($diff -gt 3) ]]; then
      return 1
    fi
    ((i++))
  done
  return 0
}

check_report_tolerance() {
  IFS=' ' read -a arr <<< "$1"
  local len=${#arr[@]}

  if [ "$len" -le 2 ]; then return 0; fi
  
  i=0
  c=0

  while [[ ($i -lt $((len-1))) ]]; do
    diff=$((arr[i+1] - arr[i]))
    if [[ $i -eq 0 ]]; then
      (( sign = diff > 0 ? 1 : -1 ))
    fi
    (( diff *= sign))
    if [[ ($diff -le 0) || ($diff -gt 3) ]]; then
      if [[ $c -eq 1 ]]; then 
        return 1
      fi

      if [[ $diff -eq 0 ]]; then
        c=1
      elif [[ $diff -lt 0 ]]; then
        # Check if the old direction is the same as the last one
        if (( (arr[len-1] - arr[len-2]) * sign > 0 )); then
          ((arr[i+1] = arr[i]))
        fi
        c=1
      elif [[ $diff -gt 3 ]]; then
        # Check if we can skip the next number
        if (( (arr[i+2] - arr[i]) * (arr[i+2] - arr[i]) <= 9 )); then
          ((arr[i+1] = arr[i]))
        fi
        c=1
      fi
    fi
    ((i++))
  done

  return 0
}

while IFS=' ' read -r line; do
  if check_report "$line"; then
    ((safe_c++))
  fi
  if check_report_tolerance "$line"; then
    ((t_safe_c++))
  fi
done < "$input"

echo "Number of safe reports (Part1) : $safe_c"
echo "Number of safe reports (Part2) : $t_safe_c"
