# Arrays and Slices

**Arrays:** store multiple elements of the same type in a particular order

Arrays have a fixed capacity which we define when the variable is declared.
They can be initialised the following ways:

- [N]type{v1, v2, ..., vN}, for example:

```go
    [5]int{1, 2, 3, 4, 5}
```

- [...]type{v1, v2, ..., vN}, for example:

```go
    [...]int{6, 7, 8, 9}
```

## Slices

Most of the time, you probably won't use arrays. Instead, go has slices. These have dynamic sizes.

## Range

`range` lets you iterate over an array. On each iteration, `range` returns two values - the index and the value.

## Misc

- `%v` is the specifier to print the `default` format.

## Wrapping Up

We have covered:

- Arrays
- Slices
  - The various ways to make them
  - How they have a fixed capacity but you can create new slices from old ones using append
  - How to slice, slices!
- `len` to get the length of an array or slice
- Test coverage tool
- `reflect.DeepEqual` and why it's useful but can reduce the type-safety of your code
- `slices.Equal`

We've used slices and arrays with integers but they work with any other type too, including arrays/slices themselves. So you can declare a variable of [][]string if you need to.

[Here is an example](https://play.golang.org/p/bTrRmYfNYCp) of slicing an array and how changing the slice affects the original array; but a "copy" of the slice will not affect the original array. [Another example](https://play.golang.org/p/Poth8JS28sc) of why it's a good idea to make a copy of a slice after slicing a very large slice.
