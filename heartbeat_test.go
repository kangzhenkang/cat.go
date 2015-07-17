package cat

import "testing"
import "fmt"
import "time"

func TestStatus(t *testing.T) {
	ts := time.Now()
	h := NewHeartbeat("system", "192.168.141.131", nil)
	h.Set("FrameworkThread", "httpthread", "0.0")
	h.Set("FrameworkThread", "catthread", "1.0")
	h.Set("System", "CPU", "0.9")
	h.Complete()
	te := time.Since(ts)
	fmt.Println("result:", string(h.GetData()))
	fmt.Println("time:", te)
}