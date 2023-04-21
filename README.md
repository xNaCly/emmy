# Emmy - a math interpreter

Emmy supports a subset of mathematical methods and operations and is intended to compute simple mathematical expressions.

## Name

Emmy is named after [Emmy Noether](https://en.wikipedia.org/wiki/Emmy_Noether), a german mathematician who contributed to abstract algebra.

## Usage

Invoke `emmy` without any arguments to start the repl and with an argument to instantly get the computed result:

```sh
emmy
# or
emmy "25+14"
```

### Examples:

#### Arithmetik:

the emmy shell is prefixed with `E>`

```text
E> 25+25
= 50.0
E> sqrt{5}
= 2.236068
E> log{25} / cos{7} * sin{3}
1.657187
E> 2^7
= 128
E> 5%2
= 1
E> scalarProd{[3,2,1], [1,2,3]}
= 10
E> crossProd{[3,2,1], [1,2,3]}
= {4,-8,4}
```
