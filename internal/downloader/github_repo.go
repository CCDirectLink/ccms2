package downloader

import (
	"net/url"
	"net/http"
	"fmt"
	"strings"
	"errors"
)

func fromRepoToArchiveUrl(repoUrl string) (string, error) {
	u, err := url.Parse(repoUrl)
	if err != nil {
		return "", err
	}

	if u.Hostname() != "github.com" {
		return "", errors.New("fromRepoToArchiveUrl: Invalid host name provided")

	}
	parts := strings.Split(u.Path, "/")

	if len(parts) < 5 {
		return "", errors.New("fromRepoToArchiveUrl: Invalid url provided")
	}

	if parts[3] != "tree" {
		return "", errors.New("fromRepoToArchiveUrl: Invalid url provided")
	}

	np := append([]string{}, parts[0:3]...)
	np = append(np, []string{"archive", "refs", "heads"}...)
	np = append(np, parts[4])

	nru := strings.Join(np, "/")
	nru = fmt.Sprintf("%s://%s%s.zip", u.Scheme, u.Host, nru)

	return nru, nil
}

func FromGithubRepo(repoUrl string) error {
	url, err := fromRepoToArchiveUrl(repoUrl)
	if err != nil {
		return err
	}

	resp, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	finalURL := resp.Request.URL.String()
	fmt.Println(finalURL)
	return nil
}
