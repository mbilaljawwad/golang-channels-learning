package tests

import (
	"reflect"
	"testing"
	"time"

	"github.com/mbilaljawwad/golang-channels-learning/internal/concurrency"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func mockWebsiteChecker(url string) bool {
	blacklistedWebsites := [3]string{
		"https://torrent.com",
		"https://utorrent.com",
		"https://jtorrent.com",
	}
	allowed := true
	for _, bw := range blacklistedWebsites {
		if bw == url {
			allowed = false
		}
	}
	return allowed
}

func TestConcurrency(t *testing.T) {
	websites := []string{
		"https://facebook.com",
		"https://getmyboat.com",
		"https://utorrent.com",
		"https://linkedin.com",
	}

	results := concurrency.CheckWebsites(mockWebsiteChecker, websites)
	expected := map[string]bool{
		"https://facebook.com":  true,
		"https://getmyboat.com": true,
		"https://utorrent.com":  false,
		"https://linkedin.com":  true,
	}

	if ok := reflect.DeepEqual(results, expected); !ok {
		t.Errorf("Result: %v, Expected: %v", results, expected)
	}
}

func BenchmarkConcurrency(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
