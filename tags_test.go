package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/technicallyty/r2d2/semver"
)

func TestReadRepoTags(t *testing.T) {
	t.Setenv(tagsEnv, "v1.2.3\nv1.2.4\nlog/v1.2.5\nx/tx/v1.2.5")
	tags := readRepoTags()
	if len(tags) == 0 {
		t.Error("no tags found")
	}
	fmt.Print(tags)
}

func TestGetLatestTagForPkg(t *testing.T) {
	testCases := []struct {
		name string
		tags []string
		pkg  string
		exp  string
	}{
		{
			name: "valid single",
			tags: []string{"v1.2.4", "log/v1.2.5", "x/tx/v1.2.5"},
			pkg:  "log",
			exp:  "v1.2.5",
		},
		{
			name: "valid multiple",
			tags: []string{"v1.2.4", "log/v1.2.5", "log/v1.2.6", "x/tx/v1.2.5"},
			pkg:  "log",
			exp:  "v1.2.6",
		},
		{
			name: "final match",
			tags: []string{"v1.2.4", "log/v1.2.5", "log/v1.2.6", "x/tx/v1.2.5", "x/tx/v1.2.6"},
			pkg:  "x/tx",
			exp:  "v1.2.6",
		},
		{
			name: "match on root version",
			tags: []string{"v1.2.4", "log/v1.2.5", "log/v1.2.6", "x/tx/v1.2.5", "x/tx/v1.2.6"},
			pkg:  "v",
			exp:  "v1.2.4",
		},
		{
			name: "no match",
			tags: []string{"v1.2.4", "log/v1.2.5", "log/v1.2.6", "x/tx/v1.2.5", "x/tx/v1.2.6"},
			pkg:  "foo",
			exp:  "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			latest, err := getLatestTagForPkg(tc.pkg, tc.tags)
			if tc.exp == "" {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				expectedSemver, err := semver.Parse(tc.exp)
				require.NoError(t, err)
				require.Equal(t, expectedSemver, latest)
			}
		})
	}
}
