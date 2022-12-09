package worker

import "sync"

// Worker is interface used by worker.
type Worker interface {
	// PushJob is used to push errorless job to the worker.
	PushJob(job func())

	// PushJobWithError is used to push a job with an error to the worker. if an error occurs,
	// it will send to the listener.
	//
	// Note: if the listener is nil, the error will be dismissed.
	PushJobWithError(job func() error)

	// Wait is used to wait for all workers to finish all jobs in the job channel.
	Wait()

	// GracefulShutdown is used to wait for all workers to finish all jobs in the job channel
	// and gracefully shutdown all workers and channels.
	GracefulShutdown()
}

// worker is a struct to implement Worker interface.
type worker struct {
	numWorker    int
	jobChannel   chan func() error
	maxRetries   int
	errorChannel chan error
	listenerFunc func(err error)
	wg           *sync.WaitGroup
	isDead       bool
}

// NewWorker is used to create a new Worker.
// numWorker is the number of workers to execute jobs synchronously (default 1).
// maxRetries is the maximum number of retries to repeat the job if an error occurs (default 1).
// errorListenerFunc is a user custom error handling when a job occurs the error (default nil).
//
// After calling NewWorker, this will immediately execute listenJobChannel X times and listenErrorChannel 1 time,
// where X is equal to numWorker. So we don't need a function to start the worker,
// its already started when calling NewWorker.
//
// Note: if numWorker/maxRetries < default value, it will automatically adjust to the default value.
// if errorListenerFunc is nil, it will not listen to the error channel and the error channel will be not activated.
func NewWorker(numWorker int, maxRetries int, errorListenerFunc func(err error)) Worker {
	w := &worker{
		numWorker:    numWorker,
		maxRetries:   maxRetries,
		listenerFunc: errorListenerFunc,
		wg:           &sync.WaitGroup{},
	}

	w.initListenJobChannel()
	w.initListenErrorChannel()

	return w
}
