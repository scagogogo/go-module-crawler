package go_module_crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// https://index.golang.org/index?since=2019-04-10T20:50:56.247227Z

// Package {"Path":"github.com/beorn7/perks","Package":"v0.0.0-20180321164747-3a771d992973","Timestamp":"2019-04-10T20:50:56.247227Z"}
type Package struct {
	Path      string `json:"Path"`
	Version   string `json:"Version"`
	Timestamp string `json:"Timestamp"`
}

func (x *Repository) Index(ctx context.Context, since string) ([]*Package, error) {
	targetUrl := x.BuildIndexURL(since)
	responseBytes, err := x.Request(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	packageSlice := make([]*Package, 0)
	split := strings.Split(string(responseBytes), "\n")
	for _, line := range split {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		r := &Package{}
		err := json.Unmarshal([]byte(line), &r)
		if err != nil {
			return nil, fmt.Errorf("format json %s error: %s", line, err.Error())
		}
		packageSlice = append(packageSlice, r)
	}
	return packageSlice, nil
}

func (x *Repository) BuildIndexURL(since string) string {
	return fmt.Sprintf("%s/index?since=%s", x.options.IndexServerURL, since)
}
