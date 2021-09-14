package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSth(t *testing.T) {
	now := time.Now()
	//local := time.Local() //没有local方法了
	fmt.Println("now : ",now)
}
