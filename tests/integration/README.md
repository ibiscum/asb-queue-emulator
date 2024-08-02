## Integration tests

The integration tests are automated tests for the end to end flow, that includes servers or rabbitmq parts that must be running.

### Running them
To run them, run the `runTest.sh` file from the project root like `./tests/integration/runTests.sh`. This creates and spins up the rabbitmq docker container, run the tests and stops the docker container.

### Adding tests
To add tests, create a new folder and file. Then create a structure that implements the `base.TestSuite` interface.

For example:

```go
package simpletest

import (
    "abs-queue-emulator/tests/integration/base"
)

type OneTest struct {
    // Put properties to be used accrosed the tests here
}

var _ base.TestSuite = &OneTest{}

func (t *OneTest) BeforeSuite() error {

}


func (t *OneTest) RunSuite() error {

}

func (t *OneTest) AfterSuite() error {

}

```

Do test setup and cleanup on the `BeforeSuite` and `AfterSuite` respectively. Even on failure, the `AfterSuite` will/should always run.

To compare results, use the utility functions from the `utils` package inside the folder of the same name inside this module. 

Finally to include the testing suite when running the `runTests.sh`, add an instance of the struct to the `tests` variable on `main.go`.

For example:

```go
package main
import (
    // ...imports
    "abs-queue-emulator/tests/integration/onetest"
    /// imports...
)

//...
var tests = []base.TestSuite{
    //...
    &onetest.OneTest{/*init the necesary parameters*/},
    //...
}
```