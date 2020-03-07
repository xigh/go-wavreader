package wavreader

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type testData struct {
	name   string
	format uint16
	bits   uint16
	rate   uint32
	chans  uint16
}

var (
	testFiles = []testData{
		testData{
			name:   "hello-16b-44100-1ch.wav",
			format: 1,
			bits:   16,
			rate:   44100,
			chans:  1,
		},
		testData{
			name:   "hello-16b-44100-2ch.wav",
			format: 1,
			bits:   16,
			rate:   44100,
			chans:  2,
		},
		testData{
			name:   "hello-u8b-44100-1ch.wav",
			format: 1,
			bits:   8,
			rate:   44100,
			chans:  1,
		},
	}

	testFilesAlt = []testData{
		testData{
			name:   "hello-alaw-44100-1ch.wav",
			format: 6,
			bits:   0,
			rate:   44100,
			chans:  1,
		},
		testData{
			name:   "hello-mulaw-44100-1ch.wav",
			format: 7,
			bits:   0,
			rate:   44100,
			chans:  1,
		},
	}

	testFilesExt = []testData{
		testData{
			name:   "hello-32b-44100-1ch.wav",
			format: 0xfffe,
			bits:   0,
			rate:   44100,
			chans:  1,
		},
		testData{
			name:   "hello-f32le-44100-1ch.wav",
			format: 0xfffe,
			bits:   0,
			rate:   44100,
			chans:  1,
		},
		testData{
			name:   "hello-f64le-44100-1ch.wav",
			format: 0xfffe,
			bits:   0,
			rate:   44100,
			chans:  1,
		},
	}
)

// TestWavReader ...
func TestWavReader(t *testing.T) {
	testWavFiles(t, testFiles)
}

/*
// TestWavReader2 ...
func TestWavReaderAlt(t *testing.T) {
	testWavFiles(t, testFilesAlt)
}

// TestWavReader3 ...
func TestWavReaderExt(t *testing.T) {
	testWavFiles(t, testFilesExt)
}
*/

func testWavFiles(t *testing.T, testFiles []testData) {
	for _, f := range testFiles {
		fmt.Printf("%s:\n", f.name)
		r, err := os.Open(filepath.Join("tests", f.name))
		if err != nil {
			t.Fatalf("%s: %v", f.name, err)
		}
		defer r.Close()
		wav, err := New(r)
		if err != nil {
			dumpHead(t, r)
			t.Fatalf("%s: %v", f.name, err)
		}
		if wav.format != f.format {
			t.Fatalf("%s: unexpected format: got %d, expected %d",
				f.name, wav.format, f.format)
		}
		if wav.bits != f.bits {
			t.Fatalf("%s: unexpected bits: got %d, expected %d",
				f.name, wav.bits, f.bits)
		}
		if wav.rate != f.rate {
			t.Fatalf("%s: unexpected rate: got %d, expected %d",
				f.name, wav.rate, f.rate)
		}
		if wav.chans != f.chans {
			t.Fatalf("%s: unexpected chans: got %d, expected %d",
				f.name, wav.chans, f.chans)
		}
		fmt.Printf(" - bits:%d, format:%d chans:%d rate:%d\n",
			wav.bits, wav.format, wav.chans, wav.rate)
		fmt.Printf(" - size:%d count:%d duration: %v\n",
			wav.size, wav.samples, wav.Duration())
		dumpHead(t, r)
	}
}

func dumpHead(t *testing.T, r io.ReaderAt) {
	buf := make([]byte, 128)
	n, err := r.ReadAt(buf, 0)
	if err != nil {
		fmt.Printf("dump: could not read header: %v\n", err)
		return
	}
	if n < 24 {
		fmt.Printf("dump: file too small: %d bytes\n", n)
	}
	a := 0
	z := len(buf)
	for z > 0 {
		n := 16
		if n > z {
			n = z
		}
		p0 := make([]string, 16)
		for i := 0; i < 16; i++ {
			if i < n {
				p0[i] = fmt.Sprintf("%02x", buf[a+i])
			} else {
				p0[i] = "  "
			}
		}
		p1 := ""
		for i := 0; i < n; i++ {
			c := buf[a+i]
			if c < 32 || c > 127 {
				c = '.'
			}
			p1 += fmt.Sprintf("%c ", c)
		}
		fmt.Printf("%08x: %s   %s\n", a, strings.Join(p0, " "), p1)
		a += n
		z -= n
	}
}
