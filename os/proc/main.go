package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	v, ok := os.LookupEnv("MYAPP")
	if ok {
		fmt.Println(v)
	} else {
		//prepare env values
		name, err := os.Executable()
		if err != nil {
			log.Fatalf("os.Executable err: %v", err)
		}
		newEnv := make([]string, 1)
		newEnv = append(newEnv, "MYAPP="+name)

		//start process
		p, err := os.StartProcess(name, []string{name}, &os.ProcAttr{Env: newEnv})
		if err != nil {
			log.Fatalf("os.StartProcess err: %v", err)
		}

		//wait for process end and print result
		r, err := p.Wait()
		if err != nil {
			log.Printf("p.Wait err: %v", err)
		}
		fmt.Printf("Child process exited with status %v", r.ExitCode())
	}
}
