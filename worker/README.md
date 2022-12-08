# WORKER

---

**worker** can execute many jobs asynchronously. This worker can also do retries and can use a custom error handler when the job function occurs error

# MOCKING

---

```shell
mockgen -source={source_path}/github.com/bearaujus/bgdk/worker/init.go -destination={destination_path}/worker_mock.go
```

# HOW TO USE

---

- ### Basic Usage
```go
package main

import (
	"fmt"
	"github.com/bearaujus/bgdk/worker"
	"time"
)

func main() {
	// create new worker
	w := worker.NewWorker(3, 1, nil)

	// iterate 10 times
	numIter := 10
	for i := 1; i <= numIter; i++ {
		idx := i

		// create job func
		job := func() {
			// we want to see when the job is executed, so we added time.sleep for a second
			time.Sleep(time.Second)

			// print the hh:mm:dd from time.now and the idx
			fmt.Printf("[%v] idx: %v\n", time.Now().Format("15:04:05"), idx)
		}

		// push job to the worker
		w.PushJob(job)
	}

	// wait for all workers to finish all jobs in the job channel
	w.Wait()
}
```
```text
[21:43:51] idx: 3
[21:43:51] idx: 1
[21:43:51] idx: 7
[21:43:52] idx: 6
[21:43:52] idx: 2
[21:43:52] idx: 8
[21:43:53] idx: 10
[21:43:53] idx: 4
[21:43:53] idx: 9
[21:43:54] idx: 5
```

- ### Worker With Error Listener
```go
package main

import (
	"errors"
	"fmt"
	"github.com/bearaujus/bgdk/worker"
	"time"
)

func main() {
	// create error listener func
	errListenerFunc := func(err error) {
		// TODO: if the job occurs error, you can do whatever you want with the error (ex. log them, or push the error to the db)
		// in this example, i just printed the error
		fmt.Println(err)
	}

	// create new worker
	w := worker.NewWorker(3, 1, errListenerFunc)

	// iterate 10 times
	numIter := 10
	for i := 1; i <= numIter; i++ {
		idx := i

		// create job func
		jobWithError := func() error {
			// we want to see when the job is executed, so we added time.sleep for a second
			time.Sleep(time.Second)

			// we want to pass error when idx is even number
			if idx%2 == 0 {
				return errors.New(fmt.Sprintf("err at idx %v", idx))
			}

			return nil
		}

		// push job to the worker
		w.PushJobWithError(jobWithError)
	}

	// we can also push errorless job to the worker
	// create job func
	job := func() {
		time.Sleep(time.Second)
	}

	// push errorless job to the worker
	w.PushJob(job)

	// wait for all workers to finish all jobs in the job channel
	w.Wait()
}
```
```text
err at idx 4
err at idx 8
err at idx 6
err at idx 10
err at idx 2
```

- ### Gracefully Shutdown The Worker
```go
package main

import (
	"fmt"
	"github.com/bearaujus/bgdk/worker"
	"time"
)

func main() {
	// create new worker
	w := worker.NewWorker(3, 1, nil)

	// iterate 10 times
	numIter := 10
	for i := 1; i <= numIter; i++ {
		idx := i

		// create job func
		job := func() {
			// we want to see when the job is executed, so we added time.sleep for a second
			time.Sleep(time.Second)

			// print the hh:mm:dd from time.now and the idx
			fmt.Printf("[%v] idx: %v\n", time.Now().Format("15:04:05"), idx)
		}

		// we want to graceful shutdown the worker when idx == 5
		if idx == 5 {
			// the output should be idx 1-4 only, the idx 5-6 will not be executed by the worker
			w.GracefulShutdown()
		}

		// push job to the worker
		w.PushJob(job)
	}
}
```
```text
[22:15:10] idx: 1
[22:15:10] idx: 4
[22:15:10] idx: 2
[22:15:11] idx: 3
```

# INSTALLATION

---

See [https://github.com/bearaujus/bgdk](https://github.com/bearaujus/bgdk)
