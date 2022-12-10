# WORKER

`worker` can execute many jobs asynchronously.
`worker` also can do retries and use a custom error listener function 
to listen to the job when the job occurs an error.

See [Documentation](https://github.com/bearaujus/bgdk/blob/master/worker/init.go)

# IMPORT

- Package

```go
import "github.com/bearaujus/bgdk/worker"
```

- Mock

```go
import mockWorker "github.com/bearaujus/bgdk/worker/mock"
```

# Examples

- [Basic Usage](#basic-usage)
- [ErrListenerFunc](#errlistenerfunc)
- [Job Retries - Success](#job-retries---success)
- [Job Retries - Triggering ErrListenerFunc](#job-retries---triggering-errlistenerfunc)
- [Wait - Advanced Usage](#wait---advanced-usage)
- [Shutdown The Worker](#shutdown-the-worker)

### Basic Usage

```go
package main

import (
  "fmt"
  "github.com/bearaujus/bgdk/worker"
  "time"
)

func main() {
  numWorker := 2
  maxRetries := 1
  // create a new worker
  w := worker.NewWorker(numWorker, maxRetries, nil)

  for i := 1; i <= 6; i++ {
    idx := i
    // push a job to the worker
    w.PushJob(func() error {
      time.Sleep(time.Second)
      fmt.Printf("[%v] idx: %v\n", time.Now().Format("15:04:05"), idx)
      return nil
    })
  }

  // wait for all workers to finish all jobs in the job channel
  w.Wait()
}
```

```text
[18:13:06] idx: 4
[18:13:06] idx: 2
[18:13:07] idx: 5
[18:13:07] idx: 1
[18:13:08] idx: 6
[18:13:08] idx: 3
```

&nbsp;

### ErrListenerFunc

```go
package main

import (
  "errors"
  "fmt"
  "github.com/bearaujus/bgdk/worker"
)

func main() {
  numWorker := 2
  maxRetries := 1
  errListenerFunc := func(err error) {
    // the errListenerFunc will be triggered
    fmt.Printf("[ERROR] %v\n", err)
  }

  // create a new worker
  w := worker.NewWorker(numWorker, maxRetries, errListenerFunc)

  for i := 1; i <= 6; i++ {
    idx := i
    if i%2 == 0 {
      // if idx is an even number, push a job to the worker with an error
      w.PushJob(func() error {
        return errors.New(fmt.Sprintf("idx is an odd number %v", idx))
      })
      continue
    }
    // if idx is an odd number, push a job to the worker without error
    w.PushJob(func() error {
      return nil
    })
  }

  // wait for all workers to finish all jobs in the job channel
  w.Wait()
}
```

```text
[ERROR] idx is an odd number 6
[ERROR] idx is an odd number 2
[ERROR] idx is an odd number 4
```

&nbsp;

### Job Retries - Success

```go
package main

import (
  "errors"
  "fmt"
  "github.com/bearaujus/bgdk/worker"
)

func main() {
  numWorker := 1
  // set max retires 3 times
  maxRetries := 3
  errListenerFunc := func(err error) {
    // the errListenerFunc is not triggered because the last retry return a nil error.
    // keep in mind that even if the job is returning an error on the first attempt,
    // this function is never triggered unless the last attempt (relative to max retires)
    // is returning an error
    fmt.Printf("[ERROR] %v\n", err)
  }

  // create a new worker
  w := worker.NewWorker(numWorker, maxRetries, errListenerFunc)

  idx := 0
  // push a job to the worker
  w.PushJob(func() error {
    idx++
    // if idx != maxRetries, assume the job had an error
    if idx != maxRetries {
      fmt.Printf("retrying... (attempt %v)\n", idx)
      return errors.New("an error")
    }

    // if idx == maxRetries, assume the job is returning success
    fmt.Println("success")
    return nil
  })

  // wait for all workers to finish all jobs in the job channel
  w.Wait()
}
```

```text
retrying... (attempt 1)
retrying... (attempt 2)
success
```

&nbsp;

### Job Retries - Triggering ErrListenerFunc

```go
package main

import (
  "errors"
  "fmt"
  "github.com/bearaujus/bgdk/worker"
)

func main() {
  numWorker := 5
  // set max retires 3 times
  maxRetries := 3
  errListenerFunc := func(err error) {
    // the errListenerFunc is triggered because the job is keep returning an error 
    // after retrying X times. where X is equal to maxRetries
    fmt.Printf("[ERROR] %v\n", err)
  }

  // create a new worker
  w := worker.NewWorker(numWorker, maxRetries, errListenerFunc)

  idx := 0
  // push a job to the worker
  w.PushJob(func() error {
    idx++
    return errors.New(fmt.Sprintf("an error (attempt %v)", idx))
  })

  // wait for all workers to finish all jobs in the job channel
  w.Wait()
}
```

```text
[ERROR] an error (attempt 3)
```

&nbsp;

### Wait - Advanced Usage

```go
package main

import (
	"fmt"
	"github.com/bearaujus/bgdk/worker"
	"time"
)

func main() {
	numWorker := 2
	maxRetries := 1

	// create a new worker
	w := worker.NewWorker(numWorker, maxRetries, nil)

	// push the first job to the worker
	fmt.Printf("-> push the first job at [%v]\n", time.Now().Format("15:04:05"))
	w.PushJob(func() error {
		time.Sleep(time.Second * 3)
		fmt.Printf("<- the first job executed at [%v]\n", time.Now().Format("15:04:05"))
		return nil
	})

	// waiting for the first job executed
	fmt.Println("waiting for the first job executed...")
	w.Wait()

	// push the second job to the worker
	fmt.Printf("-> push the second job at [%v]\n", time.Now().Format("15:04:05"))
	w.PushJob(func() error {
		time.Sleep(time.Second * 3)
		fmt.Printf("<- the second job executed at [%v]\n", time.Now().Format("15:04:05"))
		return nil
	})

	// waiting for the second job executed
	fmt.Println("waiting for the second job executed...")
	w.Wait()
}
```

```text
-> push the first job at [21:06:11]
waiting for the first job executed...
<- the first job executed at [21:06:14]
-> push the second job at [21:06:14]
waiting for the second job executed...
<- the second job executed at [21:06:17]
```

&nbsp;

### Shutdown The Worker

```go
package main

import (
  "fmt"
  "github.com/bearaujus/bgdk/worker"
)

func main() {
  numWorker := 2
  maxRetries := 1

  // create a new worker
  w := worker.NewWorker(numWorker, maxRetries, nil)

  // push a job to the worker
  w.PushJob(func() error {
    // this job will be triggered because the worker is not yet dead
    fmt.Println("before shutdown")
    return nil
  })

  // wait for all workers to finish all jobs in the job channel.
  // after all jobs is finished, shutdown all workers and channels.
  w.Shutdown()

  // push a job to the worker
  w.PushJob(func() error {
    // this job will never be triggered because the worker is already dead.
    // you cannot wake up the worker. the only way to reactivate the worker is to create new worker
    fmt.Println("after shutdown")
    return nil
  })

  // wait is no longer working since the worker is already dead
  w.Wait()
}
```

```text
before shutdown
```

&nbsp;

[Back to top](#worker) 
