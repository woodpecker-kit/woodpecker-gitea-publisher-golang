package version_check

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
)

// SemverVersionMinimumSupport
// check the version is supported greater equal than lessVersion
//
//	if the version is empty, return error
//	if the version is not supported, return error
//	if version is less than lessVersion, return error
//	if the version is supported, return nil
func SemverVersionMinimumSupport(version string, lessVersion string) error {

	if version == "" {
		return fmt.Errorf("version is empty, please check")
	}
	if lessVersion == "" {
		return fmt.Errorf("lessVersion is empty, please check")

	}

	targetVersion, errTargetVersion := semver.NewVersion(version)
	if errTargetVersion != nil {
		return fmt.Errorf("can not parse target version: %s err: %v", version, errTargetVersion)
	}

	checkVersion, errLessConstraint := semver.NewConstraint(fmt.Sprintf(">= %s", lessVersion))
	if errLessConstraint != nil {
		return fmt.Errorf("can not parse less version: %s err: %v", lessVersion, errLessConstraint)
	}

	validateOk, errors := checkVersion.Validate(targetVersion)
	if !validateOk {
		return fmt.Errorf("semver version: %s not support, err: %v", version, errors)
	}

	return nil
}

// SemverVersionConstraint
// check the version is supported by constraint
//
//	if the version or minimumVersion maximumVersion is empty, return error
//	if the version is not pass, return error
//	if the version is pass, return nil
func SemverVersionConstraint(version string, minimumVersion, maximumVersion string) error {

	if version == "" {
		return fmt.Errorf("version is empty, please check")
	}
	if maximumVersion == "" {
		return fmt.Errorf("maximum version is empty, please check")
	}
	if minimumVersion == "" {
		return fmt.Errorf("minimum version is empty, please check")
	}

	targetVersion, errTargetVersion := semver.NewVersion(version)
	if errTargetVersion != nil {
		return fmt.Errorf("can not parse target version: %s err: %v", version, errTargetVersion)
	}

	constraint := fmt.Sprintf("<= %s, >= %s", maximumVersion, minimumVersion)
	checkVersion, errConstraint := semver.NewConstraint(constraint)
	if errConstraint != nil {
		return fmt.Errorf("can not parse constraint: %s err: %v", constraint, errConstraint)
	}

	validateOk, errors := checkVersion.Validate(targetVersion)
	if !validateOk {
		return fmt.Errorf("semver version: %s not support, err: %v", version, errors)
	}

	return nil
}
