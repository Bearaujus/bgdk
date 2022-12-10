package worker

func (w *worker) PushJob(job func() error) {
	// if the worker instance is dead, the job is not executed
	if w.isDead {
		return
	}

	// assign the job to the job channel
	w.waitGroup.Add(1)
	go func() {
		w.jobChannel <- job
	}()
}

func (w *worker) Wait() {
	// if the worker instance is not dead, wait for all workers to finish all jobs in the job channel
	if !w.isDead {
		w.waitGroup.Wait()
	}
}

func (w *worker) Shutdown() {
	// if the worker instance is not dead, shutdown the worker instance
	if !w.isDead {
		// set worker instance as dead instance to the worker to prevent PushJob pushing jobs to the job channel
		w.isDead = true

		// wait for all workers to finish all jobs in the job channel
		w.waitGroup.Wait()

		// close job channel and stop all listenJobChannel iteration
		close(w.jobChannel)

		// if error channel is not nil, close error channel and stop listenErrorChannel iteration
		if w.errorChannel != nil {
			close(w.errorChannel)
		}
	}
}

func (w *worker) initListenJobChannel() {
	// create job channel
	w.jobChannel = make(chan func() error)

	// if the num worker is invalid, set to default
	if w.numWorker < 1 {
		w.numWorker = 1
	}

	// if the max retires is invalid, set to default
	if w.maxRetries < 1 {
		w.maxRetries = 1
	}

	// listen to the job channel X times synchronously, where X is the num worker
	for i := 0; i < w.numWorker; i++ {
		go w.listenJobChannel()
	}
}

func (w *worker) listenJobChannel() {
	// listen to the job channel
	for job := range w.jobChannel {
		var err error

		// if an error occurs, keep retrying until max retries.
		// when max retires are reached and the job still got an error, send the last retry error to the error channel
		for retry := w.maxRetries; retry > 0; retry-- {
			err = job()
			if err == nil {
				break
			}
		}

		// if the error channel is not nil and an error occurs, send the error to the error channel
		if w.errorChannel != nil && err != nil {
			w.waitGroup.Add(1)
			w.errorChannel <- err
		}

		w.waitGroup.Done()
	}
}

func (w *worker) initListenErrorChannel() {
	// if listener func is not nil, create and listen to the error channel
	if w.errListenerFunc != nil {
		w.errorChannel = make(chan error)
		go w.listenErrorChannel()
	}
}

func (w *worker) listenErrorChannel() {
	// listen to the error channel
	for err := range w.errorChannel {
		w.errListenerFunc(err)
		w.waitGroup.Done()
	}
}
