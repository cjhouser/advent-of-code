## Part 1
there are two exit conditions:

1. when the guard leaves the map
2. when the guard revisits a position and will leave the position in a direction she has already travelled

count any unvisited positions

## Part 2
brute force approach is to run the program over and over again, adding an obstruction to each space to check if it created a loop.

a more efficient approach would to do the
same thing, but only for positions in the known path. obstacles on the guard's path are the only ones that can affect the path the guard takes

the number of distint positions where an obstacle can be placed is bounded by the length of the unique positions in the guard's path.