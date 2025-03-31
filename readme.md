# Process Manager

A robust and distributed process management system implemented in Go, following SOLID principles and modern software architecture practices.

## 🚀 Features

- **Multiple Scheduling Algorithms**
  - First Come First Serve (FCFS)
  - Shortest Job First (SJF)
  - Round Robin (RR)
  - Extensible for new algorithms

- **Process Management**
  - Process creation and tracking
  - Real-time process state monitoring
  - Process metrics calculation
  - Thread-safe operations

- **Distributed Architecture**
  - Modular design for distributed deployment
  - Interface-based architecture
  - Scalable components

## 🏗️ Architecture

The system follows SOLID principles and is organized into the following packages:

### Core Components

- **types**: Core data structures and interfaces
- **store**: Process storage operations
- **scheduler**: Scheduling algorithms
- **queue**: Process queue operations
- **metrics**: Process metrics calculations
- **comparator**: Process comparison logic
- **manager**: High-level process management

### Key Interfaces

```go
type ProcessStore interface {
    AddProcess(process Process) error
    GetProcess(id int) (Process, error)
    UpdateProcess(process Process) error
    RemoveProcess(id int) error
    GetAllProcesses() []Process
}

type Scheduler interface {
    Schedule(processes []Process) []Process
    GetNextProcess() (Process, error)
    IsEmpty() bool
    GetAlgorithmType() string
}

type ProcessQueue interface {
    Enqueue(process Process) error
    Dequeue() (Process, error)
    Peek() (Process, error)
    IsEmpty() bool
    Size() int
}
```

## 🛠️ Installation

```bash
git clone https://github.com/yourusername/ProcessManager.git
cd ProcessManager
go mod download
```

## 📖 Usage

```go
// Create a new process manager with FCFS algorithm
pm := manager.NewProcessManager("FCFS")

// Add processes
process := types.Process{
    ID: 1,
    BurstTime: 10,
    ArrivalTime: 0,
}
pm.AddProcess(process)

// Schedule processes
scheduledProcesses := pm.ScheduleProcesses()

// Get metrics
avgWaitingTime, avgTurnaroundTime, cpuUtilization := pm.CalculateMetrics()
```

## 🔧 Configuration

The system supports various configuration options:

- **Scheduling Algorithm**: Choose between FCFS, SJF, or RR
- **Storage Backend**: Configure different storage implementations
- **Metrics Collection**: Enable/disable specific metrics
- **Queue Implementation**: Select different queue strategies

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

## 📊 Performance Metrics

The system tracks various performance metrics:

- Average Waiting Time
- Average Turnaround Time
- CPU Utilization
- Process Throughput
- Response Time

## 🔒 Thread Safety

All operations are thread-safe with proper mutex usage:

- Read/Write locks for process storage
- Atomic operations for metrics
- Safe concurrent access to queues

## 🌐 Distributed Deployment

The system can be deployed in a distributed manner:

1. **Storage Layer**: Use distributed databases
2. **Scheduling Layer**: Deploy as a separate service
3. **Queue Layer**: Implement with message queues
4. **Metrics Layer**: Integrate with monitoring systems

## 📝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Support

For support, please open an issue or contact the maintainers.
