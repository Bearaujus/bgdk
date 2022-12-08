package worker

func (w *worker) PushJob(job func()) {
	// convert job to the func() error and call PushJobWithError
	w.PushJobWithError(func() error {
		job()
		return nil
	})
}

func (w *worker) PushJobWithError(job func() error) {
	// if the worker instance is dead, the job is not executed
	if w.isDead {
		return
	}

	// assign the job to the job channel
	w.wg.Add(1)
	go func() {
		w.jobChannel <- job
	}()
}

func (w *worker) Wait() {
	// if the worker instance is not dead, wait for all workers to finish all jobs in the job channel
	if !w.isDead {
		w.wg.Wait()
	}
}

func (w *worker) GracefulShutdown() {
	// if the worker instance is not dead, shutdown the worker instance
	if !w.isDead {
		// give dead signal to the worker to prevent PushJob and PushJobWithError pushing jobs to the job channel
		w.isDead = true

		// wait for all workers to finish all jobs in the job channel
		w.wg.Wait()

		// gracefully close job channel and stop job listener
		close(w.jobChannel)

		// gracefully close error channel and stop error listener
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
			w.wg.Add(1)
			w.errorChannel <- err
		}

		w.wg.Done()
	}
}

func (w *worker) initListenErrorChannel() {
	// if the listener func is invalid, then do not listen to the error channel
	if w.listenerFunc == nil {
		return
	}

	// create error channel
	w.errorChannel = make(chan error)

	// listen to the error channel
	go w.listenErrorChannel()
}

func (w *worker) listenErrorChannel() {
	// listen to the error channel
	for err := range w.errorChannel {
		w.listenerFunc(err)
		w.wg.Done()
	}
}
