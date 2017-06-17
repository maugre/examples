# FizzBuzz

A simple test sometimes given to new programmers to prove they can code according to given instructions.

I will complete this in several languages where I am able. I won't write tests but will use `diff` to compare the output.

For example:

```
$ diff -s <(go run fizzbuzz.go) <(node fizzbuzz.js)
Files /dev/fd/63 and /dev/fd/62 are identical
```

## Specification:

Print the numbers 1 to 100 but for multiples of 3 print 'Fizz' and for multiples of 5 print 'Buzz'. Where the number is divisible by both print 'FizzBuzz'.
 
## Results

I have coded versions in the following languages with confirmed matching output:
 
* Go
* Node.js (ES6)
* Bash
