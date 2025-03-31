
OS Layer -> Resource Management

Multiple resource :-

~ Each resource needs to interact with the server where u need to run the shit.

~ DRY -> Don't repeat yourself( resource management).

~ Isolated Environment while running



OS Manages the layer of resource management.

Abstracted layer between hardware + software(interface). 

A software, peice of code managing the resource. 



Goals of the CPU utilization :-

~ Maximum utilization

~ Process starvation

~ High priority execution



Types

~ Single Process OS -> Single queue

~ Batch OS -> Multiple queue, but still multiple queue merged into single in OS Layer

~ Multi Programming -> Ready Queue

~ Multi Tasking -> Time sharing addition

~ Multi Processing OS -> addition More CPUs

~ Distributed OS

~ RTOS (Real Time OS) -> Low Error chances.



SOME CONCEPTS

Process -> any program when bought into memory is a Process.

Thread -> Its a lightweighted Process, (corresponds multiple cores).

~ Context switching is faster in Threads > Process as same memory is being shared



COMPONENTS

~ UserSpace

~ Kernel

KERNEL :-

~ Process Management

~ Memory Management

~ File Management -> tree structure

~ I/O Management



Kernel Types :-

~ Monolithic -> all functionality in single large binary( all 4 components) 

~ Micro Kernel -> User Mode(File & I/O) Kernel( Memory & Process management)

~ Hybrid Kernel -> User Mode( File management) Kernel (rest 3)

~ Nano Kernel

~ Exo Kernel



SYSTEM CALLS -> by kernels

Types :-

~ Process Control

~  File management

~ Device Management

~ Information maintaince

~ Communication management (pipe, IPC, etc.)



Process of BIOS/UEFI :-

1) POWER ON

2) Loads some settings from memory CMOS

3) Program, synchronization loads -> POST

4) Hand off to BOOT device -> Bootloader from disk/CD/drive

Loads from From MBR(master boot record) used by BIOS or EFI by UEFI

5) BootLoader loads the Full OS



PROCESS CREATION BY OS

1) Load the static program data

2) Allocate runntime stack -> function calls, etc.

3) Allocate Heap -> Part of memory used for dynamic allocation in the program.

4) I/O related task

5) OS handells task to main()



Process Tables -> Active Map of PCBs

PCB -> Process Control Block [ Process ID, Program Counter, Process state, Priority,

Open File List, Open Device List ]



~ Job scheduler -> Long Time scheduler -> Task are long runned

~ CPU scheduler -> STS -> Idle time is low for cpu throttle

~ MTS -> swap in swap out,

seprate swap space for reducing the memory overhead and puts the process there from ready queue



degree of multi programming -> Owned by LTS ( number of jobs in Ready Queue )

Orphan Process -> Async Process where the child process is detached from the parent.

Zombie Process -> Schronization between the parent wait and child exits ( early ), resource has been exited but the Process Table still has it. -> outcome of multiple process by single bash script

