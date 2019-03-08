package loader

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
)

const mcdNumberLength = 8
const segmentSizeLength = 2
const maxSegmentSize = math.MaxUint16

// Alu is the Strucure of an Application Load Unit as defined
// https://www.multos.com/uploads/GALU.pdf
type Alu struct {
	McdNumber []byte
	Code      []byte
	Data      []byte
	Dir       []byte
	Fci       []byte
	Sig       []byte
	Ktu       []byte
}

type errReader struct {
	io.Reader
	err error
}

func (e *errReader) readBytes(numBytes int) ([]byte, error) {
	if e.err != nil {
		return nil, e.err
	}

	bytes := make([]byte, numBytes)
	var numBytesRead int
	numBytesRead, e.err = e.Reader.Read(bytes)
	if numBytes != numBytesRead {
		e.err = fmt.Errorf("Expected to read %d bytes but acutally read %d", numBytes, numBytesRead)
		return nil, e.err
	}
	return bytes, e.err
}

func (e *errReader) readSegment() ([]byte, error) {
	if e.err != nil {
		return nil, e.err
	}

	segmentSizeBytes, _ := e.readBytes(segmentSizeLength)
	if e.err != nil {
		return nil, e.err
	}
	segmentSize := int(binary.BigEndian.Uint16(segmentSizeBytes))
	segment, _ := e.readBytes(segmentSize)
	return segment, e.err
}

// ParseAlu reads an ALU file and returs an Alu structure as defined
// in https://www.multos.com/uploads/GALU.pdf
func ParseAlu(filename string) (*Alu, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	reader = bufio.NewReaderSize(reader, maxSegmentSize)
	er := errReader{Reader: reader}

	mcdNumber, _ := er.readBytes(mcdNumberLength)
	code, _ := er.readSegment()
	data, _ := er.readSegment()
	dir, _ := er.readSegment()
	fci, _ := er.readSegment()
	sig, _ := er.readSegment()
	ktu, _ := er.readSegment()

	if er.err != nil {
		return nil, er.err
	}

	alu := Alu{
		McdNumber: mcdNumber,
		Code:      code,
		Data:      data,
		Dir:       dir,
		Fci:       fci,
		Sig:       sig,
		Ktu:       ktu,
	}
	return &alu, er.err
}
