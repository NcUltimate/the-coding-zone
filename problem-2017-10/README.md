# Intersections

Write a program that lists the intersection polygons when two polygons are rendered on the same graph

## Input

Your application should accept a single command-line argument that will be the name of a text file. The file will look like this:

```
4 
0 0
10 0
10 10
0 10
3
0 0
20 5
0 5
```
Each input polygon is provided in this file as, first, the number of vertices, and second, each vertex in counter-clockwise order as space separated coordinates

## Output

Your program should output a list of all intersection polygons, formatted in a similar fashion to the input polygons

## Rules of the Game

* No polygon will self-intersect
* The polygons themselves are not necessarily convex
* No point will be repeated
* No two edges of a single polygon will intersect (further clarification of "no self-interesection")
