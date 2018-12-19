package nibble_test

import (
	"bytes"
	"testing"

	"github.com/bakerolls/nibble"
)

func TestWriter_Write(t *testing.T) {
	tt := map[string]struct{ in, out []byte }{
		"two 1":   {[]byte{0xff, 0x00}, []byte{0x0f}},
		"two 2":   {[]byte{0x00, 0xff}, []byte{0xf0}},
		"two 3":   {[]byte{0xab, 0xcd}, []byte{0xca}},
		"three 1": {[]byte{0xa0, 0xb1, 0xc2}, []byte{0xba, 0x0c}},
		"three 2": {[]byte{0xab, 0xcd, 0xef}, []byte{0xca, 0x0e}},
		"four":    {[]byte{0xa0, 0xb1, 0xc2, 0xd3}, []byte{0xba, 0xdc}},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			var b bytes.Buffer
			nw := nibble.New(&b)
			if _, err := nw.Write(tc.in); err != nil {
				t.Fatal(err)
			}
			if _, err := nw.Flush(); err != nil {
				t.Fatal(err)
			}
			res := make([]byte, len(tc.out))
			if _, err := b.Read(res); err != nil {
				t.Fatal(err)
			}
			for i, b := range res {
				if tc.out[i] != b {
					t.Fatalf("expected %02x, got %02x", tc.out[i], b)
				}
			}
		})
	}
}
