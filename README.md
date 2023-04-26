# Emmy - a math interpreter

Emmy supports a subset of mathematical methods and operations and is intended to compute simple mathematical expressions.

> **Warning**
>
> Emmy is a work in progress, currently only the scanner is implemented.

## Name

Emmy is named after [Emmy Noether](https://en.wikipedia.org/wiki/Emmy_Noether), a german mathematician who contributed to abstract algebra.

## Usage

Invoke `emmy` without any arguments to start the repl and with an argument to instantly get the computed result:

```console
$ emmy "25+25"
50
```

> the emmy shell is prefixed with `ε>`

```emmy
ε> 25+25
50.0
```

### Reference

Emmy is interpreted on a line by line basis, hitting enter will make emmy forget everything put in before.

### Operations

| Name           | Symbol | Description                                                                                                                         | Example     |
| -------------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| Addition       | `+`    | Adds the preceding and the succeeding operands                                                                                      | `1+2` = 3   |
| Subtraction    | `-`    | Subtracts the succeeding operand from the preceding operand                                                                         | `3-1` = 2   |
| Multiplication | `*`    | Multiplies the preceding operand by the succeeding operand                                                                          | `2*2` = 4   |
| Division       | `/`    | Dividies the preceding operand by the succeeding operand                                                                            | `4/2` = 2   |
| Modulo         | `%`    | Dividies the preceding operand by the succeeding operand, returns remainder, read more [here](https://en.wikipedia.org/wiki/Modulo) | `16%15` = 1 |
| Power          | `^`    | Raises the preceding operand to the power of the succeeding operand                                                                 | `2^7` = 128 |

### Constants

| Name | Symbol | Description            | Example            |
| ---- | ------ | ---------------------- | ------------------ |
| π    | `@pi`  | holds the value of pi  | `@pi` = 3.1415...  |
| e    | `@e`   | holds the value of e   | `@e` = 2.7182...   |
| phi  | `@phi` | holds the value of phi | `@phi` = 1.6180... |

### Functions

> `x,y` represent the functions arguments

| Name | Symbol     | Description                                   | Example              |
| ---- | ---------- | --------------------------------------------- | -------------------- |
| sqrt | `@sqrt{x}` | calculates the square root of x               | `@sqrt{25}` = 5      |
| sin  | `@sin{x}`  | calculates the sinus of x                     | `@sin{5}`= -0.132352 |
| cos  | `@cos{x}`  | calculates the cos of x                       | `@cos{5}`= 0.991203  |
| tan  | `@tan{x}`  | calculates the tan of x                       | `@tan{5}`= -0.133526 |
| lb   | `@lb{x}`   | calculates the binary logarithm (log_2) of x  | `@lb{5}`= 2.321928   |
| ln   | `@ln{x}`   | calculates the natural logarithm (log_e) of x | `@ln{5}`= 1.609438   |
| lg   | `@lg{x}`   | calculates the common logarithm (log_10) of x | `@lg{5}`= 0.69897    |

### Planned features

- convert to other number representations:
  - binary
  - hexadecimal
  - octal
  - base32
  - base64
- bit operations (and, or, xor, nor)
- arc functions:
  - arccos
  - arcsin
  - arctan
