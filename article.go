package ogp

import (
	"html/template"
	"time"
)

// ArticleBuilder builds a `book` object.
type ArticleBuilder struct {
	WebsiteBuilder
	publishedTime  *time.Time
	modifiedTime   *time.Time
	expirationTime *time.Time
	section        string
	tags           []string
	authors        metaBuilder
}

// Title sets the `article:title` property.
func (b *ArticleBuilder) Title(title string) *ArticleBuilder {
	b.title = title
	return b
}

// URL sets the `article:url` property.
func (b *ArticleBuilder) URL(url string) *ArticleBuilder {
	b.url = url
	return b
}

// Description sets the `article:description` property.
func (b *ArticleBuilder) Description(description string) *ArticleBuilder {
	b.description = description
	return b
}

// Determiner sets the `article:determiner` property.
func (b *ArticleBuilder) Determiner(determiner string) *ArticleBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `article:locale` or adds a new `article:locale:alternate` property.
func (b *ArticleBuilder) Locale(locale string) *ArticleBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `article:site_name` property.
func (b *ArticleBuilder) SiteName(siteName string) *ArticleBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `article:image` property.
func (b *ArticleBuilder) Image(image *ImageBuilder) *ArticleBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `article:video` property.
func (b *ArticleBuilder) Video(video *VideoBuilder) *ArticleBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `article:audio` property.
func (b *ArticleBuilder) Audio(audio *AudioBuilder) *ArticleBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// PublishedTime sets the `article:published_time` property.
func (b *ArticleBuilder) PublishedTime(publishedTime time.Time) *ArticleBuilder {
	b.publishedTime = &publishedTime
	return b
}

// ModifiedTime sets the `article:modified_time` property.
func (b *ArticleBuilder) ModifiedTime(modifiedTime time.Time) *ArticleBuilder {
	b.modifiedTime = &modifiedTime
	return b
}

// ExpirationTime sets the `article:expiration_time` property.
func (b *ArticleBuilder) ExpirationTime(expirationTime time.Time) *ArticleBuilder {
	b.expirationTime = &expirationTime
	return b
}

// Section sets the `article:section` property.
func (b *ArticleBuilder) Section(section string) *ArticleBuilder {
	b.section = section
	return b
}

// Tag adds a new `article:tag` property.
func (b *ArticleBuilder) Tag(tag string) *ArticleBuilder {
	b.tags = append(b.tags, tag)
	return b
}

// Author adds a new `article:author` property.
func (b *ArticleBuilder) Author(author *ProfileBuilder) *ArticleBuilder {
	b.authors.Include(author.meta("article:author"))
	return b
}

// HTML renders the `article` object to be used in HTML templates.
func (b *ArticleBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *ArticleBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "article")
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
	if b.publishedTime != nil {
		mb.Add("article", "published_time", b.publishedTime.Format(time.RFC3339))
	}
	if b.modifiedTime != nil {
		mb.Add("article", "modified_time", b.modifiedTime.Format(time.RFC3339))
	}
	if b.expirationTime != nil {
		mb.Add("article", "expiration_time", b.expirationTime.Format(time.RFC3339))
	}
	if b.section != "" {
		mb.Add("article", "section", b.section)
	}
	for _, tag := range b.tags {
		mb.Add("article", "tag", tag)
	}
	mb.Include(&b.authors)
	return &mb
}
