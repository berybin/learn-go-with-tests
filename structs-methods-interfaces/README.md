## What are methods?

So far we have only been writing functions but we have been using some methods. When we call t.Errorf we are calling the method Errorf on the instance of our t (testing.T).

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".

---

The syntax for declaring methods is almost the same as functions and that's because they're so similar. The only difference is the syntax of the method receiver func (receiverName ReceiverType) MethodName(args).

When your method is called on a variable of that type, you get your reference to its data via the receiverName variable. In many other programming languages this is done implicitly and you access the receiver via this.

It is a convention in Go to have the receiver variable be the first letter of the type.

## Interfaces

Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.

In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

---

Notice how our helper does not need to concern itself with whether the shape is a Rectangle or a Circle or a Triangle. By declaring an interface, the helper is decoupled from the concrete types and only has the method it needs to do its job.

This kind of approach of using interfaces to declare only what you need is very important in software design and will be covered in more detail in later sections.

## Table Driven Tests

[Table driven tests](https://go.dev/wiki/TableDrivenTests) are useful when you want to build a list of test cases that can be tested in the same manner.
