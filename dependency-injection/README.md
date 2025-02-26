# Dependencyh Injection

Our function doesn't need to care where or how the printing happens, so we should accept an interface rather than a concrete type.

If we do that, we can then change the implementation to print to something we control so that we can test it. In "real life" you would inject in something that writes to stdout.

<https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection>
