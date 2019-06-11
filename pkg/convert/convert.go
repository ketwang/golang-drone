package convert

import (
	"fmt"
	"math"
	"time"
)

var (
	KB = uint64(math.Pow(2, 10))
	MB = uint64(math.Pow(2, 20))
	GB = uint64(math.Pow(2, 30))
	TB = uint64(math.Pow(2, 40))
	PB = uint64(math.Pow(2, 50))

	NS = US / 1000
	US = MS / 1000
	MS = SE / 1000
	SE = uint64(time.Second.Nanoseconds())
)

type Size uint64
type Time uint64

func (s Size) BytesToKB() float64 {
	return float64(s) / float64(KB)
}

func (s Size) BytesToMB() float64 {
	return float64(s) / float64(MB)
}

func (s Size) BytesToGB() float64 {
	return float64(s) / float64(GB)
}

func (s Size) BytesToTB() float64 {
	return float64(s) / float64(TB)
}

func (s Size) BytesToPB() float64 {
	return float64(s) / float64(PB)
}

func (s Size) String() string {
	switch {
	case uint64(s) < KB:
		return fmt.Sprintf("%f bytes", float64(s))
	case uint64(s) < MB:
		return fmt.Sprintf("%f KB", s.BytesToKB())
	case uint64(s) < GB:
		return fmt.Sprintf("%f MB", s.BytesToMB())
	case uint64(s) < TB:
		return fmt.Sprintf("%f GB", s.BytesToGB())
	case uint64(s) < PB:
		return fmt.Sprintf("%f TB", s.BytesToTB())
	default:
		return fmt.Sprintf("%f PB", s.BytesToPB())
	}
}

func (t Time) NsToUs() float64 {
	return float64(t) / float64(US)
}

func (t Time) NsToMs() float64 {
	return float64(t) / float64(MS)
}

func (t Time) NsToSecond() float64 {
	return float64(t) / float64(SE)
}

func (t Time) String() string {
	switch {
	case uint64(t) < US:
		return fmt.Sprintf("%f ns", float64(t))
	case uint64(t) < MS:
		return fmt.Sprintf("%f us", t.NsToUs())
	case uint64(t) < SE:
		return fmt.Sprintf("%f ms", t.NsToMs())
	default:
		return fmt.Sprintf("%f s", t.NsToSecond())
	}
}
