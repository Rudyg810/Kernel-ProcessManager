package fcfs

import (
	"fmt"
	"ProcessManger/types"
)

/*
Assuming its a 32 bit system with 4 bytes of CPU registers
important entities :-
	Process ID
	Burst Time
	Arrival Time
	Waiting Time
	Turnaround Time
	Completion Time
	Response Time

*/


type FCFS struct {
	processes []types.Process
}

func ( f *FCFS) Enqueue(p types.Process) {
	if(len(f.processes) ==4) {
		fmt.Println("Cannot add more than 4 processes")
		return
	}
	f.processes = append(f.processes, p)
}

func ( f *FCFS) Schedule() {
	

		
}

