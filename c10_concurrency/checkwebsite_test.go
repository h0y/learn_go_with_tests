package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://www.baidu.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string {
		"http://www.wechat.com",
		"http://www.qq.com",
		"http://www.baidu.com",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	// want := len(websites)
	// got := len(actualResults)

	// if got != want {
	// 	t.Fatalf("got '%v' want '%v'", got, want)
	// }

	expectedResults := map[string]bool {
		"http://www.wechat.com": true,
		"http://www.qq.com": true,
		"http://www.baidu.com": false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("got %v want %v", actualResults, expectedResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}