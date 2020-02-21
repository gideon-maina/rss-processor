package fetchrss

import (
	"encoding/xml"
)

type CNNRss struct {
	XMLName    xml.Name `xml:"rss"`
	Text       string   `xml:",chardata"`
	Dc         string   `xml:"dc,attr"`
	Content    string   `xml:"content,attr"`
	Atom       string   `xml:"atom,attr"`
	Media      string   `xml:"media,attr"`
	Feedburner string   `xml:"feedburner,attr"`
	Version    string   `xml:"version,attr"`
	Channel    struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        struct {
			Text   string `xml:",chardata"`
			Atom10 string `xml:"atom10,attr"`
			Rel    string `xml:"rel,attr"`
			Type   string `xml:"type,attr"`
			Href   string `xml:"href,attr"`
		} `xml:"link"`
		Image struct {
			Text  string `xml:",chardata"`
			URL   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image"`
		Generator     string `xml:"generator"`
		LastBuildDate string `xml:"lastBuildDate"`
		PubDate       string `xml:"pubDate"`
		Copyright     string `xml:"copyright"`
		Language      string `xml:"language"`
		Ttl           string `xml:"ttl"`
		Info          struct {
			Text string `xml:",chardata"`
			URI  string `xml:"uri,attr"`
		} `xml:"info"`
		Item []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			Guid        struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate string `xml:"pubDate"`
			Group   struct {
				Text    string `xml:",chardata"`
				Content []struct {
					Text   string `xml:",chardata"`
					Medium string `xml:"medium,attr"`
					URL    string `xml:"url,attr"`
					Height string `xml:"height,attr"`
					Width  string `xml:"width,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"content"`
			} `xml:"group"`
			OrigLink string `xml:"origLink"`
		} `xml:"item"`
	} `xml:"channel"`
}
