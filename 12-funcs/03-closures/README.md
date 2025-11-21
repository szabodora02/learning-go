# Closures as proxy

In this exercise, you will create a **load-balancer proxy** function.
The function should be called *proxy* and should satisfy the following criteria:
- it should recieve a vararg of functions that have no parameters and return an int
- it should return a function that returns an int
- when the returned function is called, it should propagate the call to the next in line received function and return it's value
- the functions should be propagated sequentially and should loop back after the last function was called


It is important that you do not externalize the state of the function, instead use closures to accomplish a similar effect.
Just for clarity, here is an example usage of the proxy function.
```python
  one := () -> 1
  two := () -> 2
  three := () -> 3
  loadBalanced := proxy(one,two,three)
  loaBalanced() // should return 1
  loadBalanced() // should return 2
  loadBalanced() // should return 3
  loadBalanced() // should return 1 again
```


Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.