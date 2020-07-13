package gongram

import (
	"errors"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := map[string]struct {
		input string
		gramSize int
		wantNgrams []string
		wantErr error
	}{"English Regular": {input: "The way to get started is to quit talking and begin doing", gramSize: 5,
	wantNgrams: []string{"The way to get started",
		"way to get started is",
		"to get started is to",
		"get started is to quit",
		"started is to quit talking",
		"is to quit talking and",
		"to quit talking and begin",
		"quit talking and begin doing"}, wantErr: nil},
		"Bengali Regular": {input:"সাত কোটি বাঙালীরে হে মুগ্ধ জননী রেখেছ বাঙালী করে মানুষ করোনি", gramSize: 4,
			wantNgrams: []string{"সাত কোটি বাঙালীরে হে",
				"কোটি বাঙালীরে হে মুগ্ধ",
				"বাঙালীরে হে মুগ্ধ জননী",
				"হে মুগ্ধ জননী রেখেছ",
				"মুগ্ধ জননী রেখেছ বাঙালী",
				"জননী রেখেছ বাঙালী করে",
				"রেখেছ বাঙালী করে মানুষ",
				"বাঙালী করে মানুষ করোনি"}, wantErr: nil},
				"Err: Bigger gramSize than number of tokens": {input: "this is small", gramSize: 8, wantNgrams: nil, wantErr: errors.New("gram size is greater than the number of tokens")},
				"Err: Invalid gramSize": {input: "doesn't matter", gramSize: -1, wantNgrams: nil, wantErr: errors.New("please enter a valid gram size")},

}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Generate(tc.input, tc.gramSize)
			if !reflect.DeepEqual(tc.wantNgrams, got) {
				t.Fatalf("expected: %v, got: %v", tc.wantNgrams, got)
			}
			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Fatalf("expected error is: %v, but got: %v", tc.wantErr, err)
			}
		})
	}
}
