# Deflectors
You are given a rectangular grid of arbitrary size that contains an target object at a fixed location. Your goal is to fire a laser beam into the grid to hit that target. However, the grid also contains 45-degree deflectors that alter the path of the laser beam. You need to write a program that can read such a grid and determine if it is possible to hit the target.

# Input

The input will be single, multi-line text file like the following:

```
test.txt
--------
8 13
6  9
1  2  /
2  6  \
2  9  /
2 12  \
4  4  /
4  9  /
6  6  \
6 12  /
8  4  /
8  9  \
```

The first line contains the dimensions of the rectangle: _rows_ and _columns_. So in this example there are 8 rows and 13 columns. The second row contains the position of the target. In this example the target resides in the 6th row and 9th column. The remaining rows of the file contain the locations of the reflectors, one per row. The first two fields are the location of the deflector, and the third field indicates the direction of the deflector. Here is a visual display of the above text example: 

![example](https://git.enova.com/raw/fun/the-coding-zone/master/problem-2017-06/deflectors.png)

# Output

Your application must output multi-line ASCII text indicating all positions from which a horizontal or vertical laser beam would reach the target. To communicate this clearly, label the four sides of the grid as N, S, E, W so that the NW corner is at position (1, 1). Each line of the output should contain two fields: _side_ and _position_, where position is simply the row or column number for that side.

In the example above, a laser beam fired from the east side at row 4 would hit the target, so this would be outputted as `E 4`. Similarly beams shot from the positions `E 9` and `W 9` would hit the target (check for yourself!). No other positions will result in the target being hit. Hence a valid output of the program would be:

```
E 4
E 9
W 9
```

# Limits

1. You may assume that no position on the grid will contain multiple objects. Each position will either be empty, contain a single deflector, or contain the target.
2. You may assume all the positional coordinates are within the boundaries of the grid specified by the first line.

Please drop solutions into the [solutions directory](https://git.enova.com/raw/fun/the-coding-zone/master/problem-2017-06/solutions/) in the form of a PR.
