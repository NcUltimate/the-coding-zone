# Mini-Mario

Your program needs to find a valid path between Mario and the rightmost
side of an input matrix that contains the most coins.

### Input

The input will be a file with the following format:

```
4 4     # (width) (height)
0 0     # player start
0 1 C   # coin at (0, 1)
0 2 C   # coin at (0, 2)
1 0 B   # block at (1, 0)
1 3 B   # block at (1, 3)
```
Every other space is empty.

### Output

The output should be a single positive integer representing the maximum number of coins attained when following the above objective.

### Traversal

When Mario is above a solid block (or the bottom of the grid), he is capable of
four (4) different moves:

  - Standard Move: `(x, y) ⇒ (x + 1, y)`
  - Standard Jump: `(x, y) ⇒ (x, y + 2)`
  - High Jump: `(x, y) ⇒  (x, y + 2) ⇒ (x + 1, y + 3)`
  - Long Jump: `(x, y) ⇒ (x + 1, y + 1) ⇒ (x + 3, y + 1)`

When Mario ends a move above an empty space, you need to implement *falling*,
which is technically a conditional 5th move type:

  - Falling Move: `(x, y) ⇒ (x, y - 1)`

As long as Mario is above an empty space, he must fall and cannot perform any
other move. The bottom edge of the board (i.e. `y = 0`) will ultimately prevent Mario from falling infinitely.

If a block is encountered along a 'jump path', Mario must stop movement
immediately and handle the space he is in appropriately. For example, if
Mario attempts a high jump from `(x, y)` and there is a block at `(x, y + 2)` he
would stop movement at `(x, y + 1)` and be forced to begin falling.

Mario cannot step outside of the grid boundaries, even if he is performing a
move that would otherwise cause him to land in bounds.

### FAQ

Here are some things you should be aware of:

- It is possible for there to be no coins or for coins to be unreachable.
- Mario does not have to start over a block.
- Every level will have at least one path from player start to a valid ending position on the right side.
- You cannot end in the 'air' (with an empty space beneath you). 

If you have any other questions please contact us in `The Coding Zone` HipChat room!
