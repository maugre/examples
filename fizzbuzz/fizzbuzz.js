'use strict';

const FIZZ = 3, BUZZ = 5;

for (let i = 1; i <= 100; i++) {
    switch (true) {
        default:
            console.log(i);
            break;
        case i % FIZZ == 0 && i % BUZZ == 0:
            console.log('FizzBuzz');
            break;
        case i % FIZZ == 0:
            console.log('Fizz');
            break;
        case i % BUZZ == 0:
            console.log('Buzz');
            break;
    }
}
