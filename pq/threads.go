package pq

const MAX_QUEUE_SIZE = 12

type Job interface {
	Do() error
}

type JobChan chan Job

type WorkerChan chan JobChan

var (
	JobQuequ JobChan
	WorkPool WorkerChan
)
