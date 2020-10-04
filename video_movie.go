package ogp

import (
	"html/template"
	"time"
)

// VideoMovieBuilder builds a `video.movie` object.
type VideoMovieBuilder struct {
	WebsiteBuilder
	duration    int
	tags        []string
	releaseDate *time.Time
	actors      metaBuilder
	directors   metaBuilder
	writers     metaBuilder
}

// Title sets the `video:title` property.
func (b *VideoMovieBuilder) Title(title string) *VideoMovieBuilder {
	b.title = title
	return b
}

// URL sets the `video:url` property.
func (b *VideoMovieBuilder) URL(url string) *VideoMovieBuilder {
	b.url = url
	return b
}

// Description sets the `video:description` property.
func (b *VideoMovieBuilder) Description(description string) *VideoMovieBuilder {
	b.description = description
	return b
}

// Determiner sets the `video:determiner` property.
func (b *VideoMovieBuilder) Determiner(determiner string) *VideoMovieBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `video:locale` or adds a new `video:locale:alternate` property.
func (b *VideoMovieBuilder) Locale(locale string) *VideoMovieBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `video:site_name` property.
func (b *VideoMovieBuilder) SiteName(siteName string) *VideoMovieBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `video:image` property.
func (b *VideoMovieBuilder) Image(image *ImageBuilder) *VideoMovieBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `video:video` property.
func (b *VideoMovieBuilder) Video(video *VideoBuilder) *VideoMovieBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `video:audio` property.
func (b *VideoMovieBuilder) Audio(audio *AudioBuilder) *VideoMovieBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// Duration sets the `video:duration` property.
func (b *VideoMovieBuilder) Duration(duration int) *VideoMovieBuilder {
	b.duration = duration
	return b
}

// ReleaseDate sets the `video:release_date` property.
func (b *VideoMovieBuilder) ReleaseDate(releaseDate time.Time) *VideoMovieBuilder {
	b.releaseDate = &releaseDate
	return b
}

// Tag adds a new `video:tag` property.
func (b *VideoMovieBuilder) Tag(tag string) *VideoMovieBuilder {
	b.tags = append(b.tags, tag)
	return b
}

// Actor adds a new `video:actor` property.
func (b *VideoMovieBuilder) Actor(actor *ProfileBuilder, role string) *VideoMovieBuilder {
	var mb metaBuilder
	mb.Include(actor.meta("video:actor"))
	if role != "" {
		mb.Add("video:actor", "role", role)
	}
	b.actors.Include(&mb)
	return b
}

// Director adds a new `video:director` property.
func (b *VideoMovieBuilder) Director(director *ProfileBuilder) *VideoMovieBuilder {
	b.directors.Include(director.meta("video:director"))
	return b
}

// Writer adds a new `video:writer` property.
func (b *VideoMovieBuilder) Writer(writer *ProfileBuilder) *VideoMovieBuilder {
	b.writers.Include(writer.meta("video:writer"))
	return b
}

// HTML renders the `video.movie` object to be used in HTML templates.
func (b *VideoMovieBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *VideoMovieBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "video.movie")
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
	if b.duration > 0 {
		mb.Add("video", "duration", b.duration)
	}
	if b.releaseDate != nil {
		mb.Add("video", "release_date", b.releaseDate.Format(time.RFC3339))
	}
	for _, tag := range b.tags {
		mb.Add("video", "tag", tag)
	}
	mb.Include(&b.actors)
	mb.Include(&b.directors)
	mb.Include(&b.writers)
	return &mb
}
