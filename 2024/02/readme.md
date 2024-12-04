## Part 1
the obvious approach is to iterate through report to compare adjacent levels.
a report with one level is a trivial case, it will always be safe. keep
iterating until you hit two levels that are unsafe. anything that's already
been checked is safe, so there is no need to compare each element with both
it's neighbors; just the previous element is sufficient to maintain the safe
"section" of the report.

this problem can be decomposed into many smaller subproblems since only
a pair of numbers needs to be checked for the condition. a strictly monotonic
series will produce a sum equal to the number of adjacent pairs of elements,
so len(report)-1. when the sum is less than len(report)-1, there were both
increases and decreases in value.