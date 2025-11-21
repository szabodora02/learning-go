# Generic functions

In this exercise you  will have to implement a generic *flatten* function.
The flatten function receives as arguments a slice of type E and returns a slice of type R,
where E is of type []T and T is of type []R and R could be any type.

Flatten is a function that receives a list of lists and returns a list of one dimension.
For example:
```python
    values := {{1,2,3},{4,5,6}}
    flattened := flatten(values) // produces {1,2,3,4,5,6}
```


Place your code into the file exercise.go near the placeholder // INSERT YOUR CODE HERE.