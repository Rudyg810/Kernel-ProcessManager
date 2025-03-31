package dataStructures

import (
	"ProcessManger/types"
)

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

func compareProcess(process1 types.Process, process2 types.Process) bool {
	key := types.GlobalSchedulerMap[process1.ID]
	switch key {
	case "FCFS":
		return process1.ArrivalTime < process2.ArrivalTime
	case "SJF":
		return process1.BurstTime < process2.BurstTime
	case "RR":
		return process1.ArrivalTime < process2.ArrivalTime
	default:
		return process1.ArrivalTime < process2.ArrivalTime
	}
}
