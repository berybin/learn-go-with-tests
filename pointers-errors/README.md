## Pointers

In Go, when you call a function or a method the arguments are copied.

When calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we called the method from.

Without getting too computer-sciency, when you create a value - like a wallet, it is stored somewhere in memory. You can find out what the address of that bit of memory with `&myVal`.

---

We can fix this with pointers. [Pointers](https://gobyexample.com/pointers) let us point to some values and then let us change them. So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.

---

Why no dereferencing?

The makers of Go deemed dereferencing notation i.e `(*w)` cumbersome, so the language permits us to write `w.balance`, without an explicit dereference `(*w).balance`. These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.
