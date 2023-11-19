package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const input = `ckczppom`

func main() {
	fmt.Println("Part 1:", solve1(input))
	fmt.Println("Part 2:", solve2(input))
}

func find(input, zeroes string) int {

	//f, _ := os.Create("cpu.prof")
	//_ = pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	// go tool pprof -pdf cpu.prof > cpu.pdf && open cpu.pdf

	buff := bytes.NewBuffer(make([]byte, 0, len(input)+10))
	buff.Write([]byte(input))

	var sum [16]byte

	enc := make([]byte, 32)

	for i := 0; ; i++ {
		buff.Truncate(len(input))
		buff.Write([]byte(strconv.Itoa(i)))

		sum = md5.Sum(buff.Bytes())

		hex.Encode(enc, sum[:])

		if string(enc[:len(zeroes)]) == zeroes {
			return i
		}
	}
}

func solve1(input string) int {
	return find(input, `00000`)
}

func solve2(input string) int {
	return find(input, `000000`)
}
