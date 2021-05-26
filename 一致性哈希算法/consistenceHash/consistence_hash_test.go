package consistenceHash

import (
	"fmt"
	"hash/crc32"
	"testing"
)

func TestNew(t *testing.T) {
/*	data := New(3, func(data []byte) uint32 {
		num, _ := strconv.Atoi(string(data))
		return uint32(num)
	})*/
	data := New(50, crc32.ChecksumIEEE)

	data.Set("1.1.1.1", "2.2.2.2", "3.3.3.3")

	fmt.Println(data.Get("/url/picture1.jpg"))
}
