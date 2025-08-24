## Usage

1. initialize slice (way1): 

   ```
   numbers := []int{1, 2, 2}
   ```

2. initialize slice (way2): use make

   ```
   sums := make([]int, lengthOfNumbers)
   ```

   only one param :capacity

3. Slices can be sliced!

   ```
   tails := numbers[1:]
   ```

   

## Knowledge

1. slice

   The syntax is `slice[low:high]`.

    If you omit the value on one of the sides of the `:` it captures everything to that side of it. 

## Command

1. check coverage 

   ```
   go test -cover
   ```

   > ➜  slice git:(main) ✗ go test -cover
   > PASS
   > coverage: 100.0% of statements
   > ok      github.com/huan11/learn-go-by-tdd/datastructure/slice   0.381s

2. 

## TDD

1. It is important to question the value of your tests.

   It should not be a goal to have as many tests as possible, 

   but rather to have as much *confidence* as possible in your code base

   Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. 

   **Every test has a cost**.