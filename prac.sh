v=2wZqa read -p 'Enter a : ' a

read -p 'Enter b : ' b

add = $((a + b))

echo Addition of a and b is $add


# Using statemnet 

read -p 'Enter a : ' a
read -p 'Enter b : ' b

if(($a == $b))
then
    echo both a and b are equal
fi

if(($a >= $b))
then
    echo a is greater than b
fi

if(($a%2 ==0))
then
    echo a is even 
fi

# using bitwise operator 

if(($a == 'true' && $b == 'true'))
then
    echo both are true
fi

if(($a == 'true' || $b == 'true'))
then
    echo atleast one of them is true
else
    echo none of them is true 
fi 

# using left and right shift 

read -p 'Enter : c ' c
read -p 'Enter : d ' d

bitwiseAND =$((c&d))
echo bitwise of c AND d is $bitwiseAND

bitwiseOR =$((c|d))
echo bitwise of c OR d is $bitwiseOR

leftShift =$(( c<<1 ))
echo showing the leftShift value is $leftShift

rightShift =$(( c>>1))
echo showing the rightShift value is $rightShift

#using shortcut operators 

read -p 'Enter the file name : ' Filename

if [ -e $Filename ]
then   
    echo file does exist 
else 
    echo file does not exist 

fi

if [ -s $Filename ]
then
    echo the given file is not empty

else 
    echo the given file is empty 

fi

if [ -w $Filename ]
then   
    echo the given file has wrote access 
else 
    echo the givben file has no write access
fi

if [ -r $Filename ]
then 
    echo the file has read access 
else 
    echo the file does not have the read access 

fi 

if [ -x $Filename ]
then 
    echo the file has the execute access 

else 
    echo the file does not have the execute access 

fi 


 # case statement (switch)

 cars="Mustang"

 case "$cars" in
    #case 1
    "bmw") echo "Location - Germany, America" ;;

    #case 2
    "benz") echo "Location - italy, France" ;;

    #case 3 
    "Mustang") echo "Location - Germany, Spain" ;;

esac



