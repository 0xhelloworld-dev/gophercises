package smUtils

import "strings"

func IsLinkScanned(targetLink string, scannedLinks []string) bool {
	for _, link := range scannedLinks {
		if link == targetLink {
			return true
		}
	}
	return false
}

// Transforms all links to a universal format: https://{domain}/{path}
func NormalizeHref(href string, targetURL string) string {
	if strings.HasPrefix(href, targetURL) {
		//Addresses "http://test.com/about" cases
		return href
	} else if strings.HasPrefix(href, "/") {
		//Addresses "/about" cases. Returns "https://test.com/about"
		formattedURL := targetURL[:len(targetURL)-1] + href
		return formattedURL
	} else {
		//Addresses "#" cases. Returns https://test.com/#
		formattedURL := targetURL + href
		return formattedURL
	}
}

// Accepts a raw Href from link.Href
// Ideally you should check if it is in scope before processing with
func InScope(href string, targetURL string) bool {
	if strings.HasPrefix(href, targetURL) { //does href have prefix of https://target.com?
		return true
	} else if strings.HasPrefix(href, "/") { //is it a relative href "/about"
		return true
	} else {
		return false
	}

	//need to account for cases where relative link is "#" or ","
}
