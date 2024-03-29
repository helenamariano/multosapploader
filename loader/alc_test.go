package loader

import (
	"testing"

	"github.com/go-test/deep"
)

// expected strucutre from testfiles/sample.alc
// generated from https://github.com/helenamariano/multos-sample-app
var sampleAlc = Alc{
	0x01, 0x22, 0x43, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0xF2, 0x00, 0xF2, 0x00, 0x00,
	0x48, 0x00, 0x80, 0x00, 0x00, 0x01, 0x03, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x10, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x03, 0x00, 0x4E, 0x4E, 0x00,
	0x07, 0x00, 0x72, 0x45, 0xC6, 0xF8, 0x97, 0xF3, 0x92, 0x8E, 0x1D, 0xE9, 0xC8, 0x8E, 0x51, 0x4F,
	0x8C, 0xC0, 0x11, 0x08, 0x1E, 0x95, 0x55, 0x24, 0x17, 0x97, 0xB7, 0x9B, 0x6D, 0x8F, 0x55, 0xBF,
	0x20, 0xEA, 0x9B, 0xD4, 0xE7, 0x2B, 0x78, 0xF4, 0x34, 0x45, 0x0A, 0xB6, 0xA7, 0xE1, 0x34, 0xFD,
	0xD8, 0x37, 0x45, 0x82, 0x84, 0xFC, 0x39, 0xC6, 0x01, 0x22, 0x12, 0xF4, 0xA6, 0xC6, 0x10, 0xF0,
	0x47, 0xBE, 0xEA, 0xE1, 0x79, 0xF8, 0x2F, 0x53, 0x43, 0xA8, 0x91, 0x75, 0xCD, 0x2E, 0xF6, 0x07,
	0x52, 0x43, 0x40, 0xD4, 0xC1, 0xC2, 0xAA, 0xF1, 0xB1, 0x63, 0xD3, 0xDA, 0xF6, 0x2C, 0x4D, 0xA2,
	0x97, 0x2F, 0x4A, 0x56, 0x1B, 0xF1, 0xAD, 0xD1, 0xFD, 0x27, 0x94, 0xDA, 0x99, 0x03, 0xA6, 0x85,
	0xD0, 0x81, 0x07, 0x1A, 0xC2, 0x8E, 0x5A, 0x43, 0x73, 0x83, 0x53, 0x98, 0x60, 0x36, 0xE8, 0x14,
	0xA3, 0x5D,
}

func TestParseAlc(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Alc
		wantErr bool
	}{
		{"Valid ALC", args{filename: "testfiles/sample.alc"}, sampleAlc, false},
		{"File does not exist", args{filename: "fileDoesNotExist"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAlc(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAlc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(tt.want, got); diff != nil {
				t.Error(diff)
			}
		})
	}
}
