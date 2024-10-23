package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// PrimeToFile записывает в файл fname простые числа,
// которые меньше или равны n.
func PrimeToFile(n int, fname string) error {
	// допишите код
	// ...
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", fname, err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	written := 0
loop:
	for i := 0; i <= n; i++ {
		for k := 2; float64(k) <= math.Sqrt(float64(i)); k++ {
			if i%k == 0 {
				continue loop
			}
		}
		fmt.Fprintf(w, "%d", i)
		written++
		if written == 10 {
			written = 0
			fmt.Fprintln(w, "")
		} else {
			fmt.Fprintf(w, " ")
		}
	}
	err = w.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush file %s: %w", fname, err)
	}
	return nil
}

func main() {
	if err := PrimeToFile(10000, "prime.txt"); err != nil {
		panic(err)
	}
}
