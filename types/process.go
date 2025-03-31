package types

type Process struct {
	ID             int
	BurstTime      int
	ArrivalTime    int
	WaitingTime    int
	TurnaroundTime int
	CompletionTime int
}

type Scheduler struct {
	ID        int
	Processes []Process
	Key       *string
}

type SchedulerProcessMap map[int]string

var GlobalSchedulerMap = make(SchedulerProcessMap)
