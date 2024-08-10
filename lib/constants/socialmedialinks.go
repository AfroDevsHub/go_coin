package constants

import "github.com/dlclark/regexp2"

type SocialMediaLink int

const (
	GITHUB_LINK SocialMediaLink = iota
	FACEBOOK_LINK
	TWITTER
	INSTAGRAM
	LINKEDIN
	PINTEREST
	TIKTOK
	REDDIT
	DISCORD
	TELEGRAM
	YOUTUBE
	SLACK
	TWITCH
	SPOTIFY
	SOUNDCLOUD
)

var social_media_links = map[SocialMediaLink]regexp2.Regexp{
	GITHUB_LINK:   *regexp2.MustCompile(`^https?:\/\/(www\.)?github\.com\/[a-zA-Z0-9_-]+$`, regexp2.IgnoreCase),
	FACEBOOK_LINK: *regexp2.MustCompile(`^https?:\/\/(www\.)?facebook\.com\/[a-zA-Z0-9._-]+$`, regexp2.IgnoreCase),
	TWITTER:       *regexp2.MustCompile(`^https?:\/\/(www\.)?twitter\.com\/[a-zA-Z0-9_]+$`, regexp2.IgnoreCase),
	INSTAGRAM:     *regexp2.MustCompile(`^https?:\/\/(www\.)?instagram\.com\/[a-zA-Z0-9._]+$`, regexp2.IgnoreCase),
	LINKEDIN:      *regexp2.MustCompile(`^https?:\/\/(www\.)?linkedin\.com\/in\/[a-zA-Z0-9-]+$`, regexp2.IgnoreCase),
	PINTEREST:     *regexp2.MustCompile(`^https?:\/\/(www\.)?pinterest\.com\/[a-zA-Z0-9_]+$`, regexp2.IgnoreCase),
	TIKTOK:        *regexp2.MustCompile(`^https?:\/\/(www\.)?tiktok\.com\/@[a-zA-Z0-9_]+$`, regexp2.IgnoreCase),
	REDDIT:        *regexp2.MustCompile(`^https?:\/\/(www\.)?reddit\.com\/user\/[a-zA-Z0-9_]+$`, regexp2.IgnoreCase),
	DISCORD:       *regexp2.MustCompile(`^https?:\/\/(www\.)?discord\.com\/users\/[0-9]+$`, regexp2.IgnoreCase),
	TELEGRAM:      *regexp2.MustCompile(`^https?:\/\/(t(?:elegram)?\.me|telegram\.org)\/[a-zA-Z0-9_]+$`, regexp2.IgnoreCase),
	YOUTUBE:       *regexp2.MustCompile(`^https?:\/\/(www\.)?youtube\.com\/(c|channel|user)\/[a-zA-Z0-9_-]+$`, regexp2.IgnoreCase),
	SLACK:         *regexp2.MustCompile(`^https?:\/\/[a-zA-Z0-9-]+\.slack\.com\/?$`, regexp2.IgnoreCase),
	TWITCH:        *regexp2.MustCompile(`^https?:\/\/(www\.)?twitch\.tv\/[a-zA-Z0-9_]+$`, regexp2.IgnoreCase),
	SPOTIFY:       *regexp2.MustCompile(`^https?:\/\/open\.spotify\.com\/user\/[a-zA-Z0-9_]+$`, regexp2.IgnoreCase),
	SOUNDCLOUD:    *regexp2.MustCompile(`^https?:\/\/(www\.)?soundcloud\.com\/[a-zA-Z0-9_-]+$`, regexp2.IgnoreCase),
}
