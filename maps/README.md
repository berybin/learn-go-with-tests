# Maps

The key type is special. It can only be a comparable type because without the ability to tell if 2 keys are equal, we have no way to ensure that we are getting the correct value. Comparable types are explained in depth in the [language spec](https://golang.org/ref/spec#Comparison_operators).

In order to make this pass, we are using an interesting property of the map lookup. It can return 2 values. The second value is a boolean which indicates if the key was found successfully.

---

## "ok" value

```go
    func (d Dictionary) Search(word string) (string, error) {
        definition, ok := d[word]
        if !ok {
            return "", errors.New("could not find the word you were looking for")
        }

        return definition, nil
    }
```

This property allows us to differentiate between a word that doesn't exist and a word that just doesn't have a definition.

## Pointers and copies

An interesting thing about maps is you can modify them without passing a pointer.

This may make them feel like a "reference type", [but as Dave Cheney describes](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it) they are not.

> A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a `nil` value. A `nil` map behaves like an empty map when reading, but attempts to write to a `nil` map will cause a runtime panic. You can read more about maps [here](https://blog.golang.org/go-maps-in-action).

## Constant Errors

```go
    const (
        ErrNotFound   = DictionaryErr("could not find the word you were looking for")
        ErrWordExists = DictionaryErr("cannot add word because it already exists")
    )

    type DictionaryErr string

    func (e DictionaryErr) Error() string {
        return string(e)
    }
```

We made the errors constant; this required us to create our own DictionaryErr type which implements the error interface. You can read more about the details in [this excellent article by Dave Cheney](https://dave.cheney.net/2016/04/07/constant-errors). Simply put, it makes the errors more reusable and immutable.

## Note on declaring a new error for Update

We could reuse ErrNotFound and not add a new error. However, it is often better to have a precise error for when an update fails.

Having specific errors gives you more information about what went wrong. Here is an example in a web app:

> You can redirect the user when ErrNotFound is encountered, but display an error message when ErrWordDoesNotExist is encountered.
