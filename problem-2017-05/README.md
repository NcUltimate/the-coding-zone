# Simplified Chess

You're given an 8x8 chess board containing white rooks, white bishops, a white king and a black king. Assuming it's white's turn to move, you must write a program that will determine if there is a mate in one move.

## Description

Input should be taken in the form of multiline ASCII text:

```
BK h8
WK f2
WR b3
WB e4
WB h6
```

Each line represents the position of a piece.
- The first character denotes the color `(B=black, W=white)`.
- The second character denotes the piece `(K=king, B=bishop, R=rook)`.
- The third character will always gonna be a space
- The fourth character will denote the _column_ (aka _file_ in chess lingo)
- The fifth character will denote the _row_ (aka _rank_ in chess lingo)

The output of your program should produce a list of one-step checkmates using positional notation:
```
b3 - b8
```

In this example moving the white roook from b3 to b8 will result in checkmate.
Your application should list all possible checkmates, one per line.

## Credits

Special thanks to [Dustin Fox](@dfox) for proposing this problem.

