package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/technicallyty/r2d2/semver"
)

const (
	changeLogEnv         = "COSMOS_CHANGELOG"
	changeLogFileNameEnv = "COSMOS_CHANGELOG_FILE"
)

// getRequestedTag returns the tag the user wishes to create.
func getRequestedTag() (semver.SemVer, error) {
	var ver semver.SemVer
	changelog := os.Getenv(changeLogEnv)
	if changelog == "" {
		return ver, fmt.Errorf("changelog was not stored in env var %q", changeLogEnv)
	}
	var tag string
	for _, line := range strings.Split(changelog, "\n") {
		if strings.HasPrefix(line, "## [") {
			leftBrace := strings.Index(line, "[")
			rightBrace := strings.Index(line, "]")
			tag = line[leftBrace+1 : rightBrace]
			if strings.ToLower(tag) == "unreleased" {
				continue
			}
			break
		}
	}
	if tag == "" {
		return ver, fmt.Errorf("changelog did not contain a tag")
	}
	ver, err := semver.Parse(tag)
	if err != nil {
		return ver, fmt.Errorf("unable to parse tag: %w", err)
	}
	return ver, nil
}

func getPkgName() string {
	name := os.Getenv(changeLogFileNameEnv)
	split := strings.Split(name, "/")
	if len(split) == 1 {
		return ""
	}
	return strings.Join(split[0:len(split)-1], "/")
}
