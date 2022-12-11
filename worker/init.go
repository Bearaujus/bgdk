package worker

import "sync"

// Worker is an interface primarily used by NewWorker.
type Worker interface {
	// PushJob is used to push a job to the worker. if an error occurs, it will send to the worker errListenerFunc.
	//
	// Note: if the listener is nil, the error will be dismissed.
	PushJob(job func() error)

	// Wait is used to wait for all workers to finish all jobs in the job channel.
	Wait()

	// Shutdown is used to wait for all workers to finish all jobs in the job channel.
	// after all jobs is finished, shutdown all workers and channels.
	Shutdown()
}

// worker is a struct to implement the Worker interface.
type worker struct {
	numWorker       int
	jobChannel      chan func() error
	maxRetries      int
	errorChannel    chan error
	errListenerFunc func(err error)
	waitGroup       *sync.WaitGroup
	isDead          bool
}

// NewWorker is used to create a new Worker.
// numWorker is the number of workers to execute jobs synchronously (default 1).
// maxRetries is the maximum number of retries to repeat the job if an error occurs (default 1).
// errListenerFunc is a user custom error handling when a job occurs the error (default nil).
//
// After calling NewWorker, this will immediately execute listenJobChannel X times and listenErrorChannel 1 time,
// where X is equal to numWorker. We don't need a function to start the worker,
// since its already started when calling NewWorker.
//
//	w := worker.NewWorker(1, 1, func(err error) {
//		// if the job occurs error, you can do whatever you want with the error
//		// (ex. log them, or push the error to the db)
//	})
//
// Note: if numWorker/maxRetries < default value, it will automatically adjust to the default value.
// if errListenerFunc is nil, it will not listen to the error channel and the error channel will be not activated.
func NewWorker(numWorker int, maxRetries int, errListenerFunc func(err error)) Worker {
	w := &worker{
		numWorker:       numWorker,
		maxRetries:      maxRetries,
		errListenerFunc: errListenerFunc,
		waitGroup:       &sync.WaitGroup{},
	}

	w.initListenJobChannel()
	w.initListenErrorChannel()

	return w
}
