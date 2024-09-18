package command

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/FileFormatInfo/fflint/internal/argtype"
	"github.com/FileFormatInfo/fflint/internal/shared"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
)

var (
	feedFormat = argtype.NewStringSet("Feed file format", "auto", []string{"atom", "auto", "jsonfeed", "rss"})
	feedStrict bool
)

// xmlCmd represents the xml command
var feedCmd = &cobra.Command{
	Args:     cobra.MinimumNArgs(1),
	Use:      "feed [options] files...",
	Short:    "Validate feeds (RSS/Atom/Jsonfeed)",
	Long:     `Checks that your feeds are valid.  RSS, Atom and JSONFeed are supported.`,
	PreRunE:  feedInit,
	RunE:     shared.MakeFileCommand(feedCheck),
	PostRunE: feedCleanup,
}

func AddFeedCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(feedCmd)
	feedCmd.Flags().Var(&feedFormat, "format", feedFormat.HelpText())
	feedCmd.Flags().BoolVar(&feedStrict, "strict", true, "Check contents in addition to parsability")
}

func feedCheck(f *shared.FileContext) {

	data, readErr := f.ReadFile()
	if readErr != nil {
		f.RecordResult("fileRead", false, map[string]interface{}{
			"error": readErr,
		})
		return
	}
	var parseErr error
	var feed *gofeed.Feed

	detectedFeedType := gofeed.DetectFeedType(bytes.NewReader(data))
	if feedFormat.String() == "auto" {
		if detectedFeedType == gofeed.FeedTypeUnknown {
			f.RecordResult("feedDetectType", false, map[string]interface{}{
				"error":        "Unknown feed type",
				"startOfInput": Substr(string(data), 0, 100),
			})
			return
		}
	} else if (feedFormat.String() == "rss" && detectedFeedType != gofeed.FeedTypeRSS) ||
		(feedFormat.String() == "atom" && detectedFeedType != gofeed.FeedTypeAtom) ||
		(feedFormat.String() == "jsonfeed" && detectedFeedType != gofeed.FeedTypeJSON) {
		f.RecordResult("feedTypeMismatch", false, map[string]interface{}{
			"error":    "Feed type mismatch",
			"detected": detectedFeedType,
			"expected": feedFormat.String(),
		})
		return
	}

	p := gofeed.NewParser()
	feed, parseErr = p.Parse(bytes.NewReader(data))

	if parseErr != nil {
		f.RecordResult("feedParse", false, map[string]interface{}{
			"error": parseErr,
		})
		return
	}

	if feed == nil {
		f.RecordResult("feedEmpty", false, map[string]interface{}{
			"error": "Empty feed",
		})
		return
	}

	if !feedStrict {
		return
	}

	if feed.Title == "" {
		f.RecordResult("feedTitle", false, map[string]interface{}{
			"error": "Missing title",
		})
	}

	if feed.Description == "" {
		f.RecordResult("feedTitle", false, map[string]interface{}{
			"error": "Missing title",
		})
	}

	if feedParentLinkErr := IsValidUrl(feed.Link); feedParentLinkErr != nil {
		f.RecordResult("feedParentLink", false, map[string]interface{}{
			"error": feedParentLinkErr,
			"url":   feed.Link,
		})
	}

	if feedSelfLinkErr := IsValidUrl(feed.FeedLink); feedSelfLinkErr != nil {
		f.RecordResult("feedSelfLink", false, map[string]interface{}{
			"error": feedSelfLinkErr,
			"url":   feed.FeedLink,
		})
	}

	if feed.Updated != "" && feed.UpdatedParsed == nil {
		f.RecordResult("feedUpdated", false, map[string]interface{}{
			"error":   "Invalid updated date",
			"rawdate": feed.Updated,
		})
	}

	if feed.Published != "" && feed.PublishedParsed == nil {
		f.RecordResult("feedPublished", false, map[string]interface{}{
			"error":   "Invalid published date",
			"rawdate": feed.Published,
		})
	}

	if feed.Items == nil || len(feed.Items) == 0 {
		f.RecordResult("feedItems", false, map[string]interface{}{
			"error": "No items found",
		})
	} else {
		guidMap := make(map[string]int)
		for i, item := range feed.Items {
			if item.Title == "" {
				f.RecordResult("feedItemTitle", false, map[string]interface{}{
					"error": "Missing title",
					"index": i,
				})
			}
			if item.Description == "" {
				f.RecordResult("feedItemDescription", false, map[string]interface{}{
					"error": "Missing description",
					"index": i,
				})
			}
			if item.Link == "" {
				f.RecordResult("feedItemLink", false, map[string]interface{}{
					"error": "Missing link",
					"index": i,
				})
			}
			if item.Published != "" && item.PublishedParsed == nil {
				f.RecordResult("feedItemPublished", false, map[string]interface{}{
					"error":   "Invalid published date",
					"index":   i,
					"rawdate": item.Published,
				})
			}
			if item.Updated != "" && item.UpdatedParsed == nil {
				f.RecordResult("feedItemUpdated", false, map[string]interface{}{
					"error":   "Invalid updated date",
					"index":   i,
					"rawdate": item.Updated,
				})
			}
			if item.GUID == "" {
				f.RecordResult("feedItemGUID", false, map[string]interface{}{
					"error": "Missing GUID",
					"index": i,
				})
			} else if originalIndex, ok := guidMap[item.GUID]; ok {
				f.RecordResult("feedItemGUID", false, map[string]interface{}{
					"error":          "Duplicate GUID",
					"originalIndex":  originalIndex,
					"duplicateIndex": i,
					"guid":           item.GUID,
				})
			} else {
				guidMap[item.GUID] = i
			}
		}
	}
}

func feedInit(cmd *cobra.Command, args []string) error {
	return nil
}

func feedCleanup(cmd *cobra.Command, args []string) error {
	return nil
}

func IsValidUrl(target string) error {
	if target == "" {
		return fmt.Errorf("URL not set")
	}
	_, err := url.ParseRequestURI(target)
	if err != nil {
		return err
	}
	return nil
}

// UTF8-safe substring
func Substr(input string, start int, length int) string {

	if start == 0 && length >= len(input) {
		return input
	}

	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
