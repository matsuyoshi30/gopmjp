package main

import (
	"fmt"
	"time"

	"github.com/matsuyoshi30/gopmjp"
)

func main() {
	now := time.Now().Format("2006/01/02")

	fmt.Println("1992/11/12 の日本の首相は", pmjp.ToPmjp(1992, 11, 12))
	fmt.Println(now, "の日本の首相は", pmjp.ToPmjpFromTime(time.Now()))

	fmt.Println("安倍晋三氏の通算在任期間は", now, "時点で", pmjp.CalcTenure("安倍晋三"), "日")
}
