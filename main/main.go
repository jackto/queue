package main

import (
	. "awesomeProject/pq"
	"fmt"
)

type RealJob struct {
	times int
	job   Job
}

func (job RealJob) Do() error {
	job.times++
	fmt.Printf("job done %d \n", job.times)
	return nil
}

var quit chan int = make(chan int, 10)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	quit <- 0
}

func producer() {
	for i := 0; i < 1000; i++ {
		var job RealJob
		job.times = i
		JobQuequ <- job
		//fmt.Println("new job added !")
	}
	defer close(JobQuequ)
}

func comsumer(c chan Job) {
	hasMore := true
	var p Job
	for hasMore {
		if p, hasMore = <-c; hasMore {
			p.Do()
		}
	}
}

func main() {

	/*
		JobQuequ = make(JobChan,100)

		for i:=0;i<100;i++{
			var job RealJob
			JobQuequ <- job
		}
		WorkPool = make(WorkerChan,10)
		dispatcher:= new(Dispatcher)
		dispatcher.Run()
	*/

	/*
		go loop()
		<- quit
	*/
	JobQuequ = make(JobChan, 1000)
	WorkPool = make(WorkerChan, 20)

	go producer()
	dispatcher := new(Dispatcher)
	dispatcher.Run()
}
