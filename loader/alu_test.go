package loader

import (
	"testing"

	"github.com/go-test/deep"
)

// expected strucutre from testfiles/sample.alu
// generated from https://github.com/helenamariano/multos-sample-app
var sampleAlu = Alu{
	McdNumber: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	Code:      []byte{0x0F, 0x00, 0xB9, 0x01, 0x00, 0x00, 0x39, 0x01, 0x00, 0x00, 0x47, 0x01, 0xFF, 0xFF, 0x04, 0x04},
	Data:      []byte{0x00},
	Dir:       []byte{0x61, 0x0C, 0x4F, 0x01, 0x11, 0x50, 0x07, 0x59, 0x6F, 0x74, 0x69, 0x4B, 0x65, 0x79},
	Fci:       []byte{0x11, 0x22, 0x33},
	Sig:       []byte{},
	Ktu:       []byte{},
}

func TestParseAlu(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *Alu
		wantErr bool
	}{
		{"Valid ALU", args{filename: "testfiles/sample.alu"}, &sampleAlu, false},
		{"Missing data segment", args{filename: "testfiles/sample_missing_data_segment.alu"}, nil, true},
		{"Incomplete code segment", args{filename: "testfiles/sample_incomplete_code_segment.alu"}, nil, true},
		{"File does not exist", args{filename: "fileDoesNotExist"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAlu(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAlu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(tt.want, got); diff != nil {
				t.Error(diff)
			}
		})
	}
}
