package semver

import "testing"

func TestParse(t *testing.T) {
	testCases := []struct {
		name     string
		ver      string
		expected SemVer
		expErr   string
	}{
		{
			name: "valid",
			ver:  "v0.1.0",
			expected: SemVer{
				major: 0,
				minor: 1,
				patch: 0,
			},
		},
		{
			name: "valid with extra",
			ver:  "v1.2.3-rc.1",
			expected: SemVer{
				major: 1,
				minor: 2,
				patch: 3,
				extra: "rc.1",
			},
		},
		{
			name:   "invalid, missing patch",
			ver:    "v1.2",
			expErr: "invalid version, expected form: X.X.X, got: 1.2",
		},
		{
			name:   "invalid, extra is malformed",
			ver:    "v1.2.3+rc-1",
			expErr: "invalid character in patch version: 3+rc-1",
		},
		{
			name:   "empty version string",
			ver:    "",
			expErr: "empty version",
		},
		{
			name:   "missing major",
			ver:    ".2.3",
			expErr: "invalid major version: ",
		},
		{
			name:   "missing minor",
			ver:    "1..3",
			expErr: "invalid minor version: ",
		},
		{
			name:   "non-integer major",
			ver:    "vA.2.3",
			expErr: "invalid major version: A",
		},
		{
			name:   "non-integer minor",
			ver:    "v1.B.3",
			expErr: "invalid minor version: B",
		},
		{
			name:   "non-integer patch",
			ver:    "v1.2.C",
			expErr: "invalid patch version: C",
		},
		{
			name:   "negative version numbers",
			ver:    "v1.-2.3",
			expErr: "invalid version: negative version numbers are not allowed: 1.-2.3",
		},
		{
			name: "valid with long extra",
			ver:  "v9.8.7-beta.alpha.42",
			expected: SemVer{
				major: 9,
				minor: 8,
				patch: 7,
				extra: "beta.alpha.42",
			},
		},
		{
			name: "valid without leading v",
			ver:  "1.2.3",
			expected: SemVer{
				major: 1,
				minor: 2,
				patch: 3,
			},
		},
		{
			name: "valid with dash in extra only",
			ver:  "1.2.3-rc-1",
			expected: SemVer{
				major: 1,
				minor: 2,
				patch: 3,
				extra: "rc-1",
			},
		},
		{
			name: "multiple dashes in patch-extras",
			ver:  "2.3.4-feature-x-1.5",
			expected: SemVer{
				major: 2,
				minor: 3,
				patch: 4,
				extra: "feature-x-1.5",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, err := Parse(tc.ver)
			if err != nil {
				if tc.expErr == "" {
					t.Errorf("unexpected error: %v", err)
				} else if err.Error() != tc.expErr {
					t.Errorf("expected error: %v, got: %v", tc.expErr, err)
				}
			} else {
				if v != tc.expected {
					t.Errorf("expected: %v, got: %v", tc.expected, v)
				}
			}
		})
	}
}
