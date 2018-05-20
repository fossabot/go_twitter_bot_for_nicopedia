package bot

import (
	"strings"

	"github.com/pkg/errors"
)

// Behavior is bot business logic behavior switch types.
type Behavior struct {
	TweetFormat     string
	FeedURL         string
	CheckRedirect   bool
	FollowRedirect  bool
	IgnoreDuplicate bool
	StorePrefix     string
	StoreTTL        int
}

var (
	// Gunyapetter is Nicopedia oekakiko tweet account. https://twitter.com/gunyapetter
	Gunyapetter = &Behavior{
		"%s%s に %s というお絵カキコが投稿されたよ。%s",
		"https://dic.nicovideo.jp/feed/rss/n/oekaki",
		false,
		false,
		false,
		"gunyapetter:",
		-1,
	}

	// DulltterTmp is Nicopedia pikokakiko tweet account. https://twitter.com/dulltter_tmp
	DulltterTmp = &Behavior{
		"%s%s に %s というピコカキコが投稿されたよ。%s",
		"https://dic.nicovideo.jp/feed/rss/n/mml",
		false,
		false,
		false,
		"dulltter:",
		-1,
	}

	// NicopetterNewArticle is Nicopedia new general article tweet account.
	NicopetterNewArticle = &Behavior{
		"%s の記事ができたよ。%s",
		"https://dic.nicovideo.jp/feed/rss/n/a",
		true,
		false,
		true,
		"nicopetter_new:",
		-1,
	}

	// NicopetterNewRedirectArticle is Nicopedia general article is to redirect tweet account.
	NicopetterNewRedirectArticle = &Behavior{
		"%s から %s へのリダイレクトができたよ。 %s",
		"https://dic.nicovideo.jp/feed/rss/n/a",
		true,
		true,
		true,
		"nicopetter_new_redirect:",
		86400,
	}
	// NicopetterModifyRedirectArticle is Nicopedia exists article is modify the redirect.
	NicopetterModifyRedirectArticle = &Behavior{
		"%s から %s へのリダイレクトができたよ。 %s",
		"https://dic.nicovideo.jp/feed/rss/u/a",
		true,
		true,
		true,
		"nicopetter_new_redirect:",
		86400,
	}
)

// NewBehavior is cli string from Behavior pointers.
func NewBehavior(mode string) (*Behavior, error) {
	switch strings.ToLower(mode) {
	case "gunyapetter":
		return Gunyapetter, nil
	case "dulltter":
		return DulltterTmp, nil
	case "nicopetter_new":
		return NicopetterNewArticle, nil
	case "nicopetter_new_redirect":
		return NicopetterNewRedirectArticle, nil
	case "nicopetter_modify_redirect":
		return NicopetterModifyRedirectArticle, nil
	default:
		return nil, errors.New("mode is invalid string")
	}
}
