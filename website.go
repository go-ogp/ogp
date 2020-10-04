package ogp

import "html/template"

// WebsiteBuilder builds a `website` object.
type WebsiteBuilder struct {
	title       string
	url         string
	description string
	determiner  string
	locales     []string
	siteName    string
	images      []*ImageBuilder
	videos      []*VideoBuilder
	audios      []*AudioBuilder
}

// Title sets the `og:title` property.
func (b *WebsiteBuilder) Title(title string) *WebsiteBuilder {
	b.title = title
	return b
}

// URL sets the `og:url` property.
func (b *WebsiteBuilder) URL(url string) *WebsiteBuilder {
	b.url = url
	return b
}

// Description sets the `og:description` property.
func (b *WebsiteBuilder) Description(description string) *WebsiteBuilder {
	b.description = description
	return b
}

// Determiner sets the `og:determiner` property.
func (b *WebsiteBuilder) Determiner(determiner string) *WebsiteBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `og:locale` or adds a new `og:locale:alternate` property.
func (b *WebsiteBuilder) Locale(locale string) *WebsiteBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `og:site_name` property.
func (b *WebsiteBuilder) SiteName(siteName string) *WebsiteBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `og:image` property.
func (b *WebsiteBuilder) Image(image *ImageBuilder) *WebsiteBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `og:video` property.
func (b *WebsiteBuilder) Video(video *VideoBuilder) *WebsiteBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `og:audio` property.
func (b *WebsiteBuilder) Audio(audio *AudioBuilder) *WebsiteBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// HTML renders the `website` object to be used in HTML templates.
func (b *WebsiteBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *WebsiteBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "website")
	mb.Add("og", "title", b.title)
	mb.Add("og", "url", b.url)
	if b.description != "" {
		mb.Add("og", "description", b.description)
	}
	if b.determiner != "" {
		mb.Add("og", "determiner", b.determiner)
	}
	for index, locale := range b.locales {
		if index == 0 {
			mb.Add("og", "locale", locale)
		} else {
			mb.Add("og", "locale:alternate", locale)
		}
	}
	if b.siteName != "" {
		mb.Add("og", "site_name", b.siteName)
	}
	for _, image := range b.images {
		mb.Include(image.meta("og"))
	}
	for _, video := range b.videos {
		mb.Include(video.meta("og"))
	}
	for _, audio := range b.audios {
		mb.Include(audio.meta("og"))
	}
	return &mb
}
