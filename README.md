# go-lo

A collection of common datatype functions for slices.

## Each

Iterates over each element in a slice, invoking a function on each element.

```go

    input := []string{
        "one",
        "two",
        "three",
    }

    expectedCount := 0
    EachString(input, func(in string) {
        expectedCount = expectedCount + 1
    })

    // expectedCount == 3
    
```

## Filter

Returns a slice containing each element that the function returned truthy for.

```go

    input := []string{
        "one",
        "two",
        "three",
    }

    output := FilterString(input, func(in string) bool {
        return in == "two"
    })

    // output == []string{"two"}

```

## Find

Returns the first element from a slice that the function returned truthy for.

```go

    input := []string{
        "one",
        "two",
        "three",
    }

    output := FindString(input, func(in string) bool {
        return in == "two"
    })

    // output == "two"

```

## Unique

Returns a slice containing the unique elements from the supplied slice.

```go

    input := []string{
        "one",
        "two",
        "two",
        "three",
        "two",
        "three",
        "three",
    }

    output := UniqueString(input)

    // output == []string{"one", "two", "three"}

```

## Index Of

Returns an int for the position where the element was found in the slice

```go

    input := []string{
        "one",
        "two",
        "three",
    }

    output := IndexOfString(input, "two")

    // output == 1

```

## Flatten

Combines two slices into one slice

```go
    input1 := []string{
        "test1",
        "test2",
        "test3",
    }

    input2 := []string{
        "test1",
        "test2",
        "test3",
    }

    output := FlattenString(input1, input2)

    /* output == 
    []string{
        "test1",
        "test2",
        "test3",
        "test1",
        "test2",
        "test3",
    }
    */
```
