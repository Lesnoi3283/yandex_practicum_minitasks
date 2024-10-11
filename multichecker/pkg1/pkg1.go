package pkg1

import (
	"fmt"
	"sort"
)

func mulfunc(i int) (int, error) {
	return i * 2, nil
}

func errCheckFunc() {
	// формулируем ожидания: анализатор должен находить ошибку,
	// описанную в комментарии want
	mulfunc(5)           // want "expression returns unchecked error"
	res, _ := mulfunc(5) // want "assignment with unchecked error"
	fmt.Println(res)     // want "expression returns unchecked error"
	go mulfunc(5)        // want "unchecked error in goroutine call"
	defer mulfunc(5)     // want "unchecked error in defer call"
	sl := []string{"foo", "bar", "buzz"}

	//staticcheck check:
	sl = sort.StringSlice(sl) // sort.StringSlice — это не функция, а тип, выражение не отсортирует sl
	// чтобы отсортировать, нужно сделать sort.StringSlice(sl).Sort()
}
