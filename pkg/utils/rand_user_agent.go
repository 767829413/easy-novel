package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	operatingSystems = []string{
		"Windows NT 10.0; Win64; x64",
		"Windows NT 11.0; Win64; x64",
		"Macintosh; Intel Mac OS X 12_6",
		"Macintosh; Intel Mac OS X 13_4",
		"X11; Linux x86_64",
		"X11; Ubuntu; Linux x86_64",
	}

	browsers = []string{"Chrome", "Firefox", "Safari", "Edge"}

	minVersion = 100 // Minimum browser version
	maxVersion = 131 // Maximum browser version

)

// GenerateRandomUA generates a random User-Agent string
func GenerateRandomUA() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	operatingSystem := operatingSystems[random.Intn(len(operatingSystems))]
	browser := browsers[random.Intn(len(browsers))]
	majorVersion := random.Intn(maxVersion-minVersion+1) + minVersion
	minorVersion := random.Intn(10)   // Minor version is 0-9
	buildVersion := random.Intn(1000) // Build version is 0-999

	switch browser {
	case "Chrome":
		return fmt.Sprintf(
			"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s/%d.0.%d Safari/537.36",
			operatingSystem,
			browser,
			majorVersion,
			buildVersion,
		)
	case "Firefox":
		return fmt.Sprintf("Mozilla/5.0 (%s; rv:%d.0) Gecko/20100101 %s/%d.0",
			operatingSystem, majorVersion, browser, majorVersion)
	case "Safari":
		return fmt.Sprintf(
			"Mozilla/5.0 (%s) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%d.1 Safari/605.1.15",
			operatingSystem,
			majorVersion,
		)
	case "Edge":
		return fmt.Sprintf(
			"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s/%d.0.%d.0 Safari/537.36",
			operatingSystem,
			browser,
			majorVersion,
			minorVersion,
		)
	default:
		return "Unknown User-Agent"
	}
}
