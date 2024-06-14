package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	PProfHeapFile = "PPROF_HEAP_FILE"
	PProfCPUFile  = "PPROF_CPU_FILE"
)

func instrumentPprof() func() {
	heapFile := createProfileFile(os.Getenv(PProfHeapFile))
	cpuFile := createProfileFile(os.Getenv(PProfCPUFile))
	if cpuFile != nil {
		if err := pprof.StartCPUProfile(cpuFile); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}

	return func() {
		if cpuFile != nil {
			pprof.StopCPUProfile()
			if err := cpuFile.Close(); err != nil {
				fmt.Printf("failed to close CPU profile: %v\n", err)
			}
		}
		if heapFile != nil {
			// get up-to-date statistics
			runtime.GC() // nolint:revive
			if err := pprof.WriteHeapProfile(heapFile); err != nil {
				log.Fatal("could not write memory profile: ", err)
			}
			if err := heapFile.Close(); err != nil {
				fmt.Printf("failed to close heap profile: %v\n", err)
			}
		}
	}
}

func createProfileFile(path string) *os.File {
	if path == "" {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed to create %q file for profile: %v", path, err)
	}

	return f
}
