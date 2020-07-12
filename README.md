# Check

Check is a simple unit test framework to make it a bit cleaner writing unit tests in golang.

Example:
```{go}
func TestSomething(t *testing.T) {
	check.Check(t, 1, 2)
	// This line will also run
}

func TestSomethingElse(t *testing.T) {
	check.Assert(t, 1, 2)
	// This line won't run.
}
```
