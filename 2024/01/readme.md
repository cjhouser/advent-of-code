## Part 1
simple solution is to sort each list then add the elements at the same index in each list and keep a running total.

heaps are more appropriate since operations are performed on the two minimums at each iteration.

```
load the heaps
pop the minimums (or maximum, doesn't really matter)
keep a running total of the differences between each pop'd pair
```

costs O(n) space and O(n log n) time to empty out the heaps.

i dont care to analyze initialization since we are loading data from a file

## Part 2
the following expression represents the total similarity score for an integer n:

n * A.count(n) * B.count(n)

use a dictionary to count n in each list as we pop them from the heaps. sum the above expression for every n in the dictionary. will cost O(n) more space than the heap solution since we need an additional data structure. we can't use the heaps here.
