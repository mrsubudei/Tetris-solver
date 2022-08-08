# Tetris-solver

A Tetrimino block is a shape made up of 4 consecutive characters. For the purpose of this puzzle, Tetrominoes are considered fixed, meaning that they cannot be rotated and there are a total of 19 possible pieces that can be provided as input. This program finds the smallest possible square board that the tetrominoes can fit in. If it isn't possible to form a complete square, the program leaves spaces between the tetrominoes. Every peace is changed to letters, first peace to A, second to B etc.

To run the program, you should have sample.txt file with at least one tetromino inside. From folder cmd write in terminal "go run . ./" Argument "./" is the path where sample.txt is located.

Example:
```
$ cat -e sample.txt
...#$
...#$
...#$
...#$
$
....$
....$
....$
####$
$
.###$
...#$
....$
....$
$
....$
..##$
.##.$
....$
$
....$
.##.$
.##.$
....$
$
....$
....$
##..$
.##.$
$
##..$
.#..$
.#..$
....$
$
....$
###.$
.#..$
....$
$ go run . sample.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
$
```
