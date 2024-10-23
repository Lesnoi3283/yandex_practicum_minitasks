package main

import (
	"log"
	"os"
	"os/exec"
)

//task:
//В примере выше пайп сделан для обеих команд — echo и cat.
//Для минимального решения задачи — перенаправления одного потока — достаточно одного пайпа.
//Это позволит избавиться от запуска специальной горутины для io.Copy.
//Минимизируйте код примера так, чтобы можно было обойтись одним пайпом и не запускать отдельных горутин.

// Solution: I`ve commented unnecessary code and written "cmdin.Stdin = stdout"
func main() {
	cmdout := exec.Command("echo", "Hello, world!")
	stdout, err := cmdout.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmdin := exec.Command("cat")
	// указываем текущую консоль для стандартного вывода
	cmdin.Stdout = os.Stdout
	//stdin, err := cmdin.StdinPipe()
	cmdin.Stdin = stdout
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var wg sync.WaitGroup
	//wg.Add(1)
	if err = cmdout.Start(); err != nil {
		log.Fatal(err)
	}
	if err = cmdin.Start(); err != nil {
		log.Fatal(err)
	}

	//go func() {
	//	// перенаправляем потоки данных
	//	if _, err = io.Copy(stdin, stdout); err != nil {
	//		log.Fatal(err)
	//	}
	//	wg.Done()
	//	// закрываем, чтобы завершился процесс cat
	//	stdin.Close()
	//}()
	//wg.Wait()
	cmdout.Wait()
	cmdin.Wait()
}
