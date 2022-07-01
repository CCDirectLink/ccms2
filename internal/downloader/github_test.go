package downloader

import "testing"

func TestHostNameCheck1(t * testing.T) {
	url := "https://github1.com/a/b/tree/dev"
	newUrl, err := fromRepoToArchiveUrl(url);
	if err == nil {
		t.Errorf("fromRepoToArchiveUrl(%s)\n = %s expected error.", url, newUrl);
	}

}
func TestHostNameCheck2(t * testing.T) {
	url := "https://github.com/a/b/tree/dev"
	_, err := fromRepoToArchiveUrl(url);
	if err != nil {
		t.Errorf("fromRepoToArchiveUrl(%s)\n threw %s expected no error.", url, err);
	}
}


func TestSubPathLengthCheck(t * testing.T) {
	url := "https://github.com/a/b/tree"
	newUrl, err := fromRepoToArchiveUrl(url);
	if err == nil {
		t.Errorf("fromRepoToArchiveUrl(%s)\n = %s expected error.", url, newUrl);
	}
}

func TestSubPathLengthCheck2(t * testing.T) {
	url := "https://github.com/a/b/tree/dev"
	_, err := fromRepoToArchiveUrl(url);
	if err != nil {
		t.Errorf("fromRepoToArchiveUrl(%s)\n threw %s expected no error.", url, err);
	}
}

func TestOutput1(t * testing.T) {
	url := "https://github.com/a/b/tree/dev"
	expectedUrl := "https://github.com/a/b/archive/refs/heads/dev.zip"
	newUrl, _ := fromRepoToArchiveUrl(url);
	if newUrl != expectedUrl {
		t.Errorf("fromRepoToArchiveUrl(%s);\n expected %s but got %s", url, expectedUrl, newUrl);
	}
}

func TestOutput2(t * testing.T) {
	url := "https://github.com/a/b/tree/dev/abc"
	expectedUrl := "https://github.com/a/b/archive/refs/heads/dev.zip"
	newUrl, _ := fromRepoToArchiveUrl(url);
	if newUrl != expectedUrl {
		t.Errorf("fromRepoToArchiveUrl(%s);\n expected %s but got %s", url, expectedUrl, newUrl);
	}
}
