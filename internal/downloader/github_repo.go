package downloader

import (
	"net/url"
	"net/http"
	"fmt"
	"strings"
	"errors"
	"io"
	"os"
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
		return err
	}

	finalURL := resp.Request.URL.String()

	resp, err = http.Get(finalURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		em := fmt.Sprintf("FromGithubRepo: Status Code %d", resp.StatusCode)
		return errors.New(em)
	}


	// TODO: Let function user control this
	out, err := os.Create("a.zip")
	if err != nil {
		return err
	}
	defer out.Close()


	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
