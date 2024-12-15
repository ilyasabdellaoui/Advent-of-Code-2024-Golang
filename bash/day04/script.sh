#!/bin/bash
declare -A arr
input="word_search.txt"
res=0

i=0
while IFS=' ' read -r line; do
  arr[$i]=$line
  ((i++))
done < $input
n=$i
m=${#arr[0]}

# Usage: search_foo <i, j, target_word> 

search_horizontal() {
  local target=$3
  len=${#target}
  if [[ ${arr[$1]:$2:$len} == $target ]]; then ((res++)); fi
}

search_vertical() {
  local a=$1
  local target=$3
  len=${#target}

  if [[ $(($1+$len-1)) -lt $n ]]; then 
    while [[ $a -lt $(($1+$len)) ]] ; do
      if [[ "${arr[$a]:$2:1}" != "${target:$((a-$1)):1}" ]]; then return 1; fi
      ((a++))
    done
    ((res++))
  fi
}

search_diagonal_right() {
  local a=0
  local target=$3
  len=${#target}

  if [[ $(($1+$len-1)) -lt $n && $(($2+$len-1)) -lt $m ]]; then 
    while [[ $a -lt $len ]] ; do
      if [[ "${arr[$(($1+$a))]:$(($2+a)):1}" != "${target:$a:1}" ]]; then return 1; fi
      ((a++))
    done
    ((res++))
  fi
}

search_diagonal_left() {
  local a=0
  local target=$3
  len=${#target}

  if [[ $(($1+$len-1)) -lt $n && $(($2-$len+1)) -ge 0 ]]; then 
    while [[ $a -lt $len ]] ; do
      if [[ "${arr[$(($1+$a))]:$(($2-$a)):1}" != "${target:$a:1}" ]]; then return 1; fi
      ((a++))
    done
    ((res++))
  fi
}

solve_part_one() {
  word="XMAS"
  r_word=$(echo $word | rev)

  local i=0
  while [[ $i -lt $n ]]; do
    local j=0
    while [[ $j -lt $m ]]; do
      if [[ ${arr[$i]:$j:1} == ${word:0:1} ]]; then
        search_horizontal $i $j $word
        search_vertical $i $j $word
        search_diagonal_right $i $j $word
        search_diagonal_left $i $j $word
      elif [[ ${arr[$i]:$j:1} == ${r_word:0:1} ]]; then
        search_horizontal $i $j $r_word
        search_vertical $i $j $r_word
        search_diagonal_right $i $j $r_word
        search_diagonal_left $i $j $r_word
      fi
      ((j++))
    done
    ((i++))
  done
  echo $res
  res=0
}

solve_part_two() {
  word="MAS"
  r_word=$(echo $word | rev)

  local i=0
  while [[ $i -lt $n ]]; do
    local j=0
    while [[ $j -lt $m ]]; do
      old=$res
      if [[ ${arr[$i]:$j:1} == ${word:0:1} || ${arr[$i]:$j:1} == ${r_word:0:1} ]]; then
        search_diagonal_right $i $j $word
        search_diagonal_right $i $j $r_word
        if [[ $((old+1)) -eq $res ]]; then
          ((res--))
          search_diagonal_left $i $((j+2)) $word
          search_diagonal_left $i $((j+2)) $r_word
        fi
      fi
      ((j++))
    done
    ((i++))
  done
  echo $res 
  res=0
}

echo "Part 1 Solution:"
solve_part_one

echo "Part 2 Solution:"
solve_part_two
