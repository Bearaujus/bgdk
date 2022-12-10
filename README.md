# BGDK

`bgdk` or **bearaujus golang development kit** is the development kit
with the main aim is making the development easier and testable.

The `bgdk` packages have already been built with mock files,
so you can call them in your unit tests for mocking the `bgdk` packages.
The `bgdk` has been built as simply as possible 
to help developers easy to understand the context of each package.

# INSTALLATION

- Get latest `bgdk` stable version

```shell
go get "github.com/bearaujus/bgdk@v1.0.0"
```

- Alternatively, if got a timeout you can try

```shell
GO111MODULE=on GOPROXY=direct GOSUMDB=off go get "github.com/bearaujus/bgdk@v1.0.0"
```

# PACKAGES

> ### [WORKER](https://github.com/bearaujus/bgdk/tree/master/worker)

`worker` can execute many jobs asynchronously.
`worker` also can do retries and use a custom error listener function
to listen to the job when the job occurs an error.

See [Documentation](https://github.com/bearaujus/bgdk/tree/master/worker#documentation)

&nbsp;

> ### UTIL

- [UTIL/JSON](https://github.com/bearaujus/bgdk/tree/master/util/json)

`utilJSON` is utilities for the JSON files. `utilJSON` implementing core functions from [encoding/json](https://cs.opensource.google/go/go/+/master:/src/encoding/json/).
This package also has some additional I/O features such as JSON writer and JSON reader.

See [Documentation](https://github.com/bearaujus/bgdk/tree/master/util/json#documentation)

&nbsp;

- [UTIL/YAML](https://github.com/bearaujus/bgdk/tree/master/util/yaml)

`utilYAML` is utilities for the YAML files. `utilYAML` implementing core functions from [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3).
This package also has some additional I/O features such as YAML writer and YAML reader.

See [Documentation](https://github.com/bearaujus/bgdk/tree/master/util/yaml#documentation)

&nbsp;

[Back to top](#bgdk) 
