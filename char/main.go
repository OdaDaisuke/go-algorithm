package char

import (
	"fmt"
	"encoding/hex"
	"bytes"
	"encoding/binary"
)

func Start() {
	s := "abcdeqweyi"
	buf := bytes.NewBuffer([]byte("あ"))
	var le uint16
	binary.Read(buf, binary.LittleEndian, &le)
	fmt.Printf("LE: %v \n", le)
	fmt.Printf("str to hex bin %s\n", StrToHexBin(s))
}

// StrtoOctetBin Convert raw string to hex binary data
func StrToHexBin(s string) string {
	return hex.EncodeToString([]byte(s))
}
