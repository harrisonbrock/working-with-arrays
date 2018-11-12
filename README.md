**Basic of working with Array or Slice**

I'm going to show some basic of working with array or slice.

**Create an array or slice:**

```go
// this creates array of ints with the size of 5
var arr = [5]{0,1,3,4}

// this creates an slice of ints with no fixed size
var mySlice = []int{8, 1, 10, 5, 3, 2, 6, 7, 4, 9}
```

**Iterating Over Slice Or Array.**

```go
func print(arr []int) {
	for index, number := range arr {
		fmt.Printf("Index: %v and Number: %v\n", index, number)

	}
}
```

**Sum Numbers in Slice Or Array.**
```go
func sum(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}
```

**Find The Average in Slice or Array**
```go
func average(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum / len(arr)

}
```