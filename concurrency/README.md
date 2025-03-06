# Concurrency

Basics:

- **goroutines**, the basic unit of concurrency in Go, which let us manage more than one website check request.
- **anonymous functions**, which we used to start each of the concurrent processes that check websites.
- **channels**, to help organize and control the communication between the different processes, allowing us to avoid a race condition bug.
- the **race detector** which helped us debug problems with concurrent code
