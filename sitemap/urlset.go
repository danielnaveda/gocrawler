package sitemap

// Urlset is the structure expected from the sitemap
type Urlset struct {
	URL []struct {
		Loc        string `xml:"loc"`
		Lastmod    string `xml:"lastmod"`
		Priority   string `xml:"priority"`
		Changefreq string `xml:"changefreq"`
	} `xml:"url"`
}
