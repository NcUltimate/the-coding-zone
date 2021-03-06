# Simplified Chess

You're given an 8x8 chess board containing white rooks, white bishops, a white king and a black king. Assuming it's white's turn to move, you must write a program that will determine if there is a mate in one move.

Solutions to be presented on May 25, 2017!

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
- The third character will always be a space
- The fourth character will denote the _column_ (or _file_ in chess lingo)
- The fifth character will denote the _row_ (or _rank_ in chess lingo)

The output of your program should produce a list of one-step checkmates using positional notation:
```
b3 - b8
```

In this example moving the white roook from b3 to b8 will result in checkmate.
Your application should list all possible checkmates, one per line. The image below summarizes this case:

![test-case](https://git.enova.com/raw/fun/the-coding-zone/master/problem-2017-05/board.png)

## Limits

- There is no limit to the number of pieces in the board-input. For example, there could be 30 white rooks on the board.
- You may assume the board will contain at least a white king and a black king
- You may assume all inputs are correctly formatted values (e.g. first character of a line is always W or B)
- You may assume that the black king is not in check on the supplied board (otherwise it could not be white's move!)


## Credits

Special thanks to:
 - [Dustin Fox](mailto:dfox@enova.com) for proposing this problem.
 - [Nick Matelli](mailto:nmatelli@enova.com) for the test-cases

