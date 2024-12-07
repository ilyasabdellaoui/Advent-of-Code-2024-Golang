#!/bin/bash
input="input.txt"

# Creating a file for each column
while IFS=' ' read -r val1 val2; do
	echo "$val1" >> left.txt
	echo "$val2" >> right.txt
done < "$input"

# Sorting all numbers
sort -n right.txt -o right.txt
sort -n left.txt -o left.txt

sum=0

similarity=0
prev=0
count=0

sum_func() {
	local diff=$(($1-$2))
	(( sum += diff < 0 ? -diff : diff ))
}

similarity_func() {
  # Debug: echo "$prev - $l - $count - $similarity"
  if [ "$1" != "$prev" ]
  then 
    count="$(grep -c "$1" right.txt)"
  fi
  # As numbers are sorted we can use previous count
  (( similarity += count * $1 ))
  prev="$1"
}

# Get numbers and process
exec 3<left.txt
exec 4<right.txt

while read -r -u 3 l && read -r -u 4 r
do 
	sum_func "$r" "$l"
  similarity_func "$l"
done 

# Close files and remove them
exec 3<&-
exec 4<&-

rm right.txt
rm left.txt

# Show results
echo "Total Distance:"
echo "$sum"
echo "Similarity Score:"
echo "$similarity"
