package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/technicallyty/r2d2/semver"
)

const (
	orgEnv  = "REPO_OWNER"
	repoEnv = "REPO_NAME"
)

func main() {
	// godotenv.Load()

	pkg := getPkgName()
	requestedVersion, err := getRequestedTag()
	if err != nil {
		log.Fatal(err)
	}
	latestTag, err := getLatestTagForPkg(pkg, readRepoTags())
	hasLatestTag := true
	if err != nil {
		if errors.Is(err, ErrTagNotFound) {
			hasLatestTag = false
		} else {
			log.Fatal(err)
		}
	}
	// if we have a latest tag, we need to verify that the requested tag matches the update type.
	if hasLatestTag {
		updateType := getUpdateType()
		err = verifyUpdate(updateType, requestedVersion, latestTag)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = updateTag(pkg, requestedVersion, os.Getenv(orgEnv), os.Getenv(repoEnv))
	if err != nil {
		log.Fatal(err)
	}
}

func verifyUpdate(updateType UpdateType, requestedVersion, latestTag semver.SemVer) error {
	var nextVersion semver.SemVer
	switch updateType {
	case Major:
		nextVersion = latestTag.NextMajor()
	case Minor:
		nextVersion = latestTag.NextMinor()
	case Patch:
		nextVersion = latestTag.NextPatch()
	default:
		panic("invalid update type")
	}

	if nextVersion != requestedVersion {
		return fmt.Errorf("requested tag does not match intended update type: %q", updateType)
	}
	return nil
}
