# BGDK

**bearaujus golang development kit** or `bgdk` is the development kit
with the main aim is making the development easier and testable.

# INSTALLATION


```shell
go get "github.com/bearaujus/bgdk@v1.0.0"
```

# PACKAGES

> ### [UTIL](https://github.com/bearaujus/bgdk/tree/master/util)
> `util` is an `bgdk` utilities packages.

- [UTIL/JSON](https://github.com/bearaujus/bgdk/tree/master/util/json)

`utilJSON` is utilities for the JSON files. `utilJSON` implementing core functions from [encoding/json](https://cs.opensource.google/go/go/+/master:/src/encoding/json/).
This package also has some additional I/O features such as JSON writer and JSON reader.

&nbsp;

- [UTIL/YAML](https://github.com/bearaujus/bgdk/tree/master/util/yaml)

`utilYAML` is utilities for the YAML files. `utilYAML` implementing core functions from [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3).
This package also has some additional I/O features such as YAML writer and YAML reader.

&nbsp;

> ### [WORKER](https://github.com/bearaujus/bgdk/tree/master/worker)
> `worker` can execute many jobs asynchronously.
> `worker` also can do retries and use a custom error listener function
> to listen to the job when the job occurs an error.

[Back to top](#bgdk) 
