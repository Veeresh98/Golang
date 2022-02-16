echo "Hello world"

echo -n "enter the first number:"
read first_number

echo -n "enter the second number:"
read second_number

sum = $(($first_number + $second_number))

echo "sum of $first_number and $second_number is :"$sum