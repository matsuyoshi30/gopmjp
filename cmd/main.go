package main

import (
	"fmt"
	"time"

	"github.com/matsuyoshi30/gopmjp"
)

func main() {
	fmt.Println("1992/11/12 の日本の首相は ", pmjp.ToPmjp(1992, 11, 12))
	fmt.Println(time.Now().Format("2006/01/02"), "の日本の首相は ", pmjp.ToPmjpFromTime(time.Now()))
}
