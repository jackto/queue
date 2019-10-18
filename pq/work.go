package pq

import "fmt"

type Worker struct {
	JobChannel JobChan
	quit       chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			WorkPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if job != nil {
					if err := job.Do(); err != nil {
						fmt.Println("execute job failed with err %v", err)
					}
				}
			}
		}
	}()
}
