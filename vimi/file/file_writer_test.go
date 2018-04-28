package file

import "testing"

var (
	s = "test_string\n"
	b = []byte("test_byte\n") // not a const
)

const (
	filePath = "./test.txt"
)

func checkerr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteAppend(t *testing.T) {
	f, err := NewFileWriter(filePath, true)
	checkerr(t, err)
	defer f.Close()

	err = f.WriteString(s)
	checkerr(t, err)

	err = f.WriteByte(b)
	checkerr(t, err)
}

func TestWriteNotAppend(t *testing.T) {
	f, err := NewFileWriter(filePath, false)
	checkerr(t, err)
	defer f.Close()

	err = f.WriteString(s)
	checkerr(t, err)

	err = f.WriteByte(b)
	checkerr(t, err)
}
