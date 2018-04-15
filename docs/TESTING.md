# Testing

This document contains the Continuous Engineering Platform code testing guidelines. It should answer any questions you may have as an aspiring contributor. 

Contributing includes testing your changes. If you change the cePlatform code, you may need to add a new test or modify an existing test. Your contribution could even be adding tests to cePlatform. For this reason, you need to know a little about cePlatform's test infrastructure.

## Understand how to test cePlatform

cePlatform tests use the Go language's test framework. In this framework, files whose names end in `_test.go` contain test code; you'll find test files like this throughout the cePlatform repo. Use these files for inspiration when writing your own tests. For information on Go's test framework, see [Go's testing package documentation][Testing] and the [go test help][GoTestHelp].

You are responsible for _unit testing_ your contribution when you add new or change existing cePlatform code. A unit test is a piece of code that invokes a single, small piece of code (_unit of work_) to verify the unit works as expected.

Depending on your contribution, you may need to add _integration tests_. These are tests that combine two or more work units into one component. These work units each have unit tests and then, together, integration tests that test the interface between the components. 

Testing is its own specialty. If you aren't familiar with testing techniques, there is a lot of information available to you on the Web. For now, you should understand that, the Docker maintainers may ask you to write a new test or change an existing one.

## Test suites

The Continuous Engineering Platform has two test suites:

* Unit tests - use standard `go test` and [testify][Testify] assertions. They are located in the package they test. Unit tests should be fast and test only their own  package.
* Integration tests - use standard `go test` and [testify][Testify] assertions. They are located in `./integration/<component>` directories.

## Writing new tests

Most code changes will fall into one of the following categories.

### Naming test functions

The naming convention for a test function should have a name that describes the test case, in the format:

`Test<Function under test>ToReturn<Expected output><Conditions>`

### Writing tests for new features

New code should be covered by unit tests. If the code is difficult to test with a unit tests then that is a good sign that it should be refactored to make it easier to reuse and maintain. Consider accepting unexported interfaces instead of structs so that fakes can be provided for dependencies.

If the new feature includes a completely new API endpoint then a new API integration test should be added to cover the success case of that endpoint.

If the new feature does not include a completely new API endpoint consider adding the new API fields to the existing test for that endpoint. A new integration test should **not** be added for every new API field or API error case. Error cases should be handled by unit tests.

### Writing tests for bug fixes

Bugs fixes should include a unit test case which exercises the bug.

A bug fix may also include new assertions in an existing integration test.

## Running tests

Before submitting a pull request with a code change, you should run the entire cePlatform  test suite. The `Makefile` contains a target for the entire test suite, named `test`. Also, it contains several targets for testing:

| Target                 | What this target does                          |
| ---------------------- | ---------------------------------------------- |
| `test`                 | Run the unit, integration, and docker-py tests |
| `test-unit`            | Run just the unit tests                        |
| `test-integration`     | Run the integration tests for the CLI          |

To run the unit test suite, do the following:

```
make test-unit
```

To run the integration test suite, do the following:

```
make test-integration
```

Running the entire test suite on your current repository can take over half an hour. To run the test suite, do the following:

```
make test
```
<!--
## Run unit tests

We use golang standard [testing][Testing] package for our unit tests.

You can use the `TESTDIRS` environment variable to run unit tests for a single package.

```bash
$ TESTDIRS='opts' make test-unit
```

You can also use the `TESTFLAGS` environment variable to run a single test. The flag's value is passed as arguments to the `go test` command. For example, from your local host you can run the `TestBuild` test with this command:

```bash
$ TESTFLAGS='-test.run ^TestValidateIPAddress$' make test-unit
```

On unit tests, it's better to use `TESTFLAGS` in combination with `TESTDIRS` to make it quicker to run a specific test.

```bash
$ TESTDIRS='opts' TESTFLAGS='-test.run ^TestValidateIPAddress$' make test-unit
```

## Run integration tests

We use golang standard [testing][Testing] package for our integration tests. You can use the `TESTFLAGS` environment variable to run a single test. The flag's value is passed as arguments to the `go test` command. For example, from your local host you can run the `TestBuild` test with this command:

```bash
$ TESTFLAGS='-check.f DockerSuite.TestBuild*' make test-integration
```
-->

## Where to go next

Congratulations, you have successfully completed the basics you need to understand the cePlatform test framework.

[Testify]: https://github.com/stretchr/testify
[Testing]: https://golang.org/pkg/testing/
[gocheck]: https://labix.org/gocheck
[GoTestHelp]: http://golang.org/cmd/go/#hdr-Test_packages