#!/bin/bash

FIZZ=3
BUZZ=5

for i in {1..100}
do
    case 1 in
        $(( $i % $FIZZ == 0 && $i % $BUZZ == 0)))
            echo 'FizzBuzz'
        ;;
        $(( $i % $FIZZ == 0)))
            echo 'Fizz'
        ;;
        $(( $i % $BUZZ == 0)))
            echo 'Buzz'
        ;;
        *)
            echo $i
    esac
done;
