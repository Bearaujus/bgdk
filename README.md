# BGDK

**bgdk** or bearaujus golang development kit is the development kit that I use when working personally. 
But I think this repository can help many developers to work, so I decided to make this repository public

My goal in creating this repository is to make development easier when mocking and testing several projects that implement the **bgdk** package

# INSTALLATION

```shell
go get "github.com/bearaujus/bgdk"
```

# PACKAGES

> ### [WORKER](https://github.com/bearaujus/bgdk/tree/master/worker)
**worker** can execute many jobs asynchronously. This worker can also do retries and can use a custom error handler when the job function occurs error

See [Documentation](https://github.com/bearaujus/bgdk/tree/master/worker)

# TODO

Things to do on my mind right now

- Move all bearaujus golang package here
- Create Package: `bgdk/nsq`
- Create Package: `bgdk/io`
- Create Unit Tests
