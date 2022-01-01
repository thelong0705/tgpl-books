// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	memo "chap9/memo1"
	"chap9/memotest"
	"fmt"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	fmt.Println("test")
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	fmt.Println("concurrent")
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}