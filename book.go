package ogp

import (
	"html/template"
	"time"
)

// BookBuilder builds a `book` object.
type BookBuilder struct {
	WebsiteBuilder
	isbn        string
	releaseDate *time.Time
	tags        []string
	authors     metaBuilder
}

// Title sets the `book:title` property.
func (b *BookBuilder) Title(title string) *BookBuilder {
	b.title = title
	return b
}

// URL sets the `book:url` property.
func (b *BookBuilder) URL(url string) *BookBuilder {
	b.url = url
	return b
}

// Description sets the `book:description` property.
func (b *BookBuilder) Description(description string) *BookBuilder {
	b.description = description
	return b
}

// Determiner sets the `book:determiner` property.
func (b *BookBuilder) Determiner(determiner string) *BookBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `book:locale` or adds a new `book:locale:alternate` property.
func (b *BookBuilder) Locale(locale string) *BookBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `book:site_name` property.
func (b *BookBuilder) SiteName(siteName string) *BookBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `book:image` property.
func (b *BookBuilder) Image(image *ImageBuilder) *BookBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `book:video` property.
func (b *BookBuilder) Video(video *VideoBuilder) *BookBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `book:audio` property.
func (b *BookBuilder) Audio(audio *AudioBuilder) *BookBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// ISBN sets the `book:isbn` property.
func (b *BookBuilder) ISBN(isbn string) *BookBuilder {
	b.isbn = isbn
	return b
}

// ReleaseDate sets the `book:release_date` property.
func (b *BookBuilder) ReleaseDate(releaseDate time.Time) *BookBuilder {
	b.releaseDate = &releaseDate
	return b
}

// Tag adds a new `book:tag` property.
func (b *BookBuilder) Tag(tag string) *BookBuilder {
	b.tags = append(b.tags, tag)
	return b
}

// Author adds a new `book:author` property.
func (b *BookBuilder) Author(author *ProfileBuilder) *BookBuilder {
	b.authors.Include(author.meta("book:author"))
	return b
}

// HTML renders the `book` object to be used in HTML templates.
func (b *BookBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *BookBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "book")
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
	if b.isbn != "" {
		mb.Add("book", "isbn", b.isbn)
	}
	if b.releaseDate != nil {
		mb.Add("book", "release_date", b.releaseDate.Format(time.RFC3339))
	}
	for _, tag := range b.tags {
		mb.Add("book", "tag", tag)
	}
	mb.Include(&b.authors)
	return &mb
}
