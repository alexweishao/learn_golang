package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func GetMd5String(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
