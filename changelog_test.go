package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPkgName(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "valid",
			path:     "log/changelog.md",
			expected: "log",
		},
		{
			name:     "root",
			path:     "changelog.md",
			expected: "",
		},
		{
			name:     "long path",
			path:     "foo/bar/baz/tests/changelog.md",
			expected: "foo/bar/baz/tests",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(changeLogFileNameEnv, tc.path)
			pkgName := getPkgName()
			if pkgName != tc.expected {
				t.Errorf("expected: %s, got: %s", tc.expected, pkgName)
			}
		})
	}
}

func TestGetRequestedTag(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(dir+"/changelog.md", []byte(changelog), 0644)
	bz, err := os.ReadFile(dir + "/changelog.md")
	require.NoError(t, err)
	t.Setenv(changeLogEnv, string(bz))
	tag, err := getRequestedTag()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, "1.2.1", tag.String())
}

var changelog = `# Changelog

## [Unreleased]

## [v1.2.1](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv1.2.1)

## Bug Fixes

* [#24737](https://github.com/cosmos/cosmos-sdk/pull/24737) Ensure that map memory will never be reused unintentionally.

## [v1.2.0](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv1.2.0)

### Improvements

* [#24081](https://github.com/cosmos/cosmos-sdk/pull/24081) Remove "cosmossdk.io/core" dependency.

## [v1.1.0](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv1.1.0)

### Improvements

* [#23515](https://github.com/cosmos/cosmos-sdk/pull/23515) Bring in "collections/protocodec" go module as package within "collections" module.

## [v1.0.0](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv1.0.0)

### Features

* [#22641](https://github.com/cosmos/cosmos-sdk/pull/22641) Add reverse iterator support for "Triple".
* [#17656](https://github.com/cosmos/cosmos-sdk/pull/17656) Introduces "Vec", a collection type that allows to represent a growable array on top of a KVStore.
* [#18933](https://github.com/cosmos/cosmos-sdk/pull/18933) Add LookupMap implementation. It is basic wrapping of the standard Map methods but is not iterable.
* [#19343](https://github.com/cosmos/cosmos-sdk/pull/19343) Simplify IndexedMap creation by allowing to infer indexes through reflection.
* [#19861](https://github.com/cosmos/cosmos-sdk/pull/19861) Add "NewJSONValueCodec" value codec as an alternative for "codec.CollValue" from the SDK for non protobuf types.
* [#21090](https://github.com/cosmos/cosmos-sdk/pull/21090) Introduces "Quad", a composite key with four keys.
* [#20704](https://github.com/cosmos/cosmos-sdk/pull/20704) Add "ModuleCodec" method to "Schema" and "HasSchemaCodec" interface in order to support "cosmossdk.io/schema" compatible indexing.
* [#20538](https://github.com/cosmos/cosmos-sdk/pull/20538) Add "Nameable" variations to "KeyCodec" and "ValueCodec" to allow for better indexing of "collections" types.
* [#22544](https://github.com/cosmos/cosmos-sdk/pull/22544) Schema's "ModuleCodec" will now also return Enum descriptors to be registered with the indexer.

## [v0.4.0](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv0.4.0)

### Features

* [#17024](https://github.com/cosmos/cosmos-sdk/pull/17024) Introduces "Triple", a composite key with three keys.

### API Breaking

* [#17290](https://github.com/cosmos/cosmos-sdk/pull/17290) Collections iteration methods (Iterate, Walk) will not error when the collection is empty.

### Improvements

* [#17021](https://github.com/cosmos/cosmos-sdk/pull/17021) Make collections implement the "appmodule.HasGenesis" interface.

## [v0.3.0](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv0.3.0)

### Features

* [#16074](https://github.com/cosmos/cosmos-sdk/pull/16607) Introduces "Clear" method for "Map" and "KeySet"
* [#16773](https://github.com/cosmos/cosmos-sdk/pull/16773)
    * Adds "AltValueCodec" which provides a way to decode a value in two ways.
    * Adds the possibility to specify an alternative way to decode the values of "KeySet", "indexes.Multi", "indexes.ReversePair".

## [v0.2.0](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv0.2.0)

### Features

* [#16074](https://github.com/cosmos/cosmos-sdk/pull/16074)  Makes the generic Collection interface public, still highly unstable.

### API Breaking

* [#16127](https://github.com/cosmos/cosmos-sdk/pull/16127)  In the "Walk" method the call back function being passed is allowed to error.

## [v0.1.0](https://github.com/cosmos/cosmos-sdk/releases/tag/collections%2Fv0.1.0)

Collections "v0.1.0" is released! Check out the [docs](https://docs.cosmos.network/main/build/packages/collections) to know how to use the APIs.
`
