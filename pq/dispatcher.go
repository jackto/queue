package pq

const MaxWorkPoolSize = 20

type Dispatcher struct {
	Workers []*Worker
	quit    chan bool
}

func (d *Dispatcher) Run() {
	for i := 0; i < MaxWorkPoolSize; i++ {
		worker := new(Worker)
		worker.JobChannel = make(JobChan, 10)
		worker.quit = make(chan bool)
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}

	for {
		select {
		case job := <-JobQuequ:

			worker := <-WorkPool
			worker <- job
			/*
				go func(job Job) {
					worker := <- WorkPool
					worker <- job
				}(job)
			*/

		case <-d.quit:
			return
		}
	}
}
