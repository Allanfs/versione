package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VersionUpdater interface {
	Update(string) string
}

type MajorUpdater struct {
	Regex []regexp.Regexp
}

func (MajorUpdater) Update(versao string) string {

	pos := strings.Split(versao, ".")
	major, _ := strconv.Atoi(pos[0])
	pos[0] = strconv.Itoa(major + 1)

	return fmt.Sprintf("%s.0.0", pos[0])
}

type MinorUpdater struct {
	Regex *regexp.Regexp
}

func (MinorUpdater) Update(versao string) string {

	pos := strings.Split(versao, ".")
	minor, _ := strconv.Atoi(pos[1])
	pos[1] = strconv.Itoa(minor + 1)

	return fmt.Sprintf("%s.%s.0", pos[0], pos[1])
}

type PatchUpdater struct {
	Regex *regexp.Regexp
}

func (PatchUpdater) Update(versao string) string {

	pos := strings.Split(versao, ".")
	patch, _ := strconv.Atoi(pos[2])
	pos[2] = strconv.Itoa(patch + 1)

	return fmt.Sprintf("%s.%s.%s", pos[0], pos[1], pos[2])
}

func StringRegexMatchBranch(regexes []string, branch string) bool {

	for _, regex := range regexes {
		rgx := regexp.MustCompile(regex)
		if rgx.MatchString(branch) {
			return true
		}
	}

	return false
}

func GetUpdaterByName(name string) VersionUpdater {
	name = strings.ToLower(name)
	switch name {
	case "major":
		return MajorUpdater{}
	case "minor":
		return MinorUpdater{}
	case "patch":
	default:
		return PatchUpdater{}

	}
	return PatchUpdater{}

}
