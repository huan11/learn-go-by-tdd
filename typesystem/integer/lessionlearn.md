# executable documentation (Testable Examples)

## knowledge 

1. Examples are tests

2. examples are functions that reside in a package’s _test.go files.

3. example functions take no arguments and begin with the word Example instead of Test.

   eg

   ```
   package integers
   
   import (
   	"fmt"
   	"testing"
   )
   
   func TestAdder(t *testing.T) {
   	expected := 4
   	sum := Add(2, 2)
   
   	if sum != expected {
   		t.Errorf("expected '%d' but got '%d'", expected, sum)
   	}
   }
   
   func ExampleAdd() {
   	sum := Add(1, 5)
   	fmt.Println(sum)
   	// Output: 6
   }
   ```

   

4. Output comments:The test passes if the test’s output matches its output comment.

   If we remove the output comment entirely,then the example function is compiled but not executed

   Examples without output comments are useful for demonstrating code that cannot run as unit tests, such as that which accesses the network, while guaranteeing the example at least compiles.

5. a “whole file example.” (A whole file example is a file that ends in `_test.go` +  one example function)

   Eg: https://pkg.go.dev/sort#example-package



## Tool

### local 

```
pkgsite
```

http://localhost:8080/



### remote

If you publish your code with examples to a public URL, you can share the documentation of your code at [pkg.go.dev](https://pkg.go.dev/)