package semver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// SemVer represents the parts of a semantic version string. It includes major, minor, and patch as integers.
// Extra contains any information after the full version string; i.e., release cut information.
type SemVer struct {
	major int
	minor int
	patch int
	extra string
}

func (s SemVer) String() string {
	ver := fmt.Sprintf("v%d.%d.%d", s.major, s.minor, s.patch)
	if s.extra != "" {
		ver += "-" + s.extra
	}
	return ver
}

func (s SemVer) NextMajor() SemVer {
	return SemVer{
		major: s.major + 1,
		minor: 0,
		patch: 0,
	}
}

func (s SemVer) NextMinor() SemVer {
	return SemVer{
		major: s.major,
		minor: s.minor + 1,
		patch: 0,
	}
}

func (s SemVer) NextPatch() SemVer {
	return SemVer{
		major: s.major,
		minor: s.minor,
		patch: s.patch + 1,
	}
}

// Parse parses a version string, returning a SemVer object.
func Parse(ver string) (SemVer, error) {
	var v SemVer
	if ver == "" {
		return v, errors.New("empty version")
	}
	majorIdx, minorIdx, patchIdx := 0, 1, 2
	if ver[0] == 'v' {
		ver = ver[1:]
	}
	split := strings.SplitN(ver, ".", 3)
	if len(split) < 3 {
		return v, fmt.Errorf("invalid version, expected form: X.X.X, got: %s", ver)
	}
	var err error
	v.major, err = strconv.Atoi(split[majorIdx])
	if err != nil {
		return v, fmt.Errorf("invalid major version: %s", split[majorIdx])
	}
	v.minor, err = strconv.Atoi(split[minorIdx])
	if err != nil {
		return v, fmt.Errorf("invalid minor version: %s", split[minorIdx])
	}
	// this means we have something like 1.2.3-rc.1
	if len(split[patchIdx]) > 1 {
		if split[patchIdx][1] != '-' {
			return v, fmt.Errorf("invalid character in patch version: %s", split[patchIdx])
		}
		fullPatch := split[patchIdx]
		split = strings.SplitN(fullPatch, "-", 2)
		v.patch, err = strconv.Atoi(split[0])
		v.extra = split[1]
	} else {
		v.patch, err = strconv.Atoi(split[patchIdx])
		if err != nil {
			return v, fmt.Errorf("invalid patch version: %s", split[patchIdx])
		}
	}
	if v.major < 0 || v.minor < 0 || v.patch < 0 {
		return v, fmt.Errorf("invalid version: negative version numbers are not allowed: %s", ver)
	}
	return v, nil
}
