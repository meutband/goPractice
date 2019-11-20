package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

//like fmt.Sprint(values) but adds commas
func intsToString(values []int) string {
	//initialize buffer
	var buf bytes.Buffer

	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
