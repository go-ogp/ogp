package ogp

import (
	"html/template"
	"time"
)

// VideoTVShowBuilder builds a `video.tv_show` object.
type VideoTVShowBuilder struct {
	WebsiteBuilder
	duration    int
	tags        []string
	releaseDate *time.Time
	actors      metaBuilder
	directors   metaBuilder
	writers     metaBuilder
}

// Title sets the `video:title` property.
func (b *VideoTVShowBuilder) Title(title string) *VideoTVShowBuilder {
	b.title = title
	return b
}

// URL sets the `video:url` property.
func (b *VideoTVShowBuilder) URL(url string) *VideoTVShowBuilder {
	b.url = url
	return b
}

// Description sets the `video:description` property.
func (b *VideoTVShowBuilder) Description(description string) *VideoTVShowBuilder {
	b.description = description
	return b
}

// Determiner sets the `video:determiner` property.
func (b *VideoTVShowBuilder) Determiner(determiner string) *VideoTVShowBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `video:locale` or adds a new `video:locale:alternate` property.
func (b *VideoTVShowBuilder) Locale(locale string) *VideoTVShowBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `video:site_name` property.
func (b *VideoTVShowBuilder) SiteName(siteName string) *VideoTVShowBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `video:image` property.
func (b *VideoTVShowBuilder) Image(image *ImageBuilder) *VideoTVShowBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `video:video` property.
func (b *VideoTVShowBuilder) Video(video *VideoBuilder) *VideoTVShowBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `video:audio` property.
func (b *VideoTVShowBuilder) Audio(audio *AudioBuilder) *VideoTVShowBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// Duration sets the `video:duration` property.
func (b *VideoTVShowBuilder) Duration(duration int) *VideoTVShowBuilder {
	b.duration = duration
	return b
}

// ReleaseDate sets the `video:release_date` property.
func (b *VideoTVShowBuilder) ReleaseDate(releaseDate time.Time) *VideoTVShowBuilder {
	b.releaseDate = &releaseDate
	return b
}

// Tag adds a new `video:tag` property.
func (b *VideoTVShowBuilder) Tag(tag string) *VideoTVShowBuilder {
	b.tags = append(b.tags, tag)
	return b
}

// Actor adds a new `video:actor` property.
func (b *VideoTVShowBuilder) Actor(actor *ProfileBuilder, role string) *VideoTVShowBuilder {
	var mb metaBuilder
	mb.Include(actor.meta("video:actor"))
	if role != "" {
		mb.Add("video:actor", "role", role)
	}
	b.actors.Include(&mb)
	return b
}

// Director adds a new `video:director` property.
func (b *VideoTVShowBuilder) Director(director *ProfileBuilder) *VideoTVShowBuilder {
	b.directors.Include(director.meta("video:director"))
	return b
}

// Writer adds a new `video:writer` property.
func (b *VideoTVShowBuilder) Writer(writer *ProfileBuilder) *VideoTVShowBuilder {
	b.writers.Include(writer.meta("video:writer"))
	return b
}

// HTML renders the `video.tv_show` object to be used in HTML templates.
func (b *VideoTVShowBuilder) HTML() template.HTML {
	return b.meta("og").HTML()
}

func (b *VideoTVShowBuilder) meta(ns string) *metaBuilder {
	var mb metaBuilder
	if ns == "og" {
		mb.Add(ns, "type", "video.tv_show")
		mb.Add(ns, "title", b.title)
		mb.Add(ns, "url", b.url)
	} else {
		mb.Add(ns, "", b.url)
		if b.title != "" {
			mb.Add(ns, "title", b.title)
		}
	}
	if b.description != "" {
		mb.Add(ns, "description", b.description)
	}
	if b.determiner != "" {
		mb.Add(ns, "determiner", b.determiner)
	}
	for index, locale := range b.locales {
		if index == 0 {
			mb.Add(ns, "locale", locale)
		} else {
			mb.Add(ns, "locale:alternate", locale)
		}
	}
	if b.siteName != "" {
		mb.Add(ns, "site_name", b.siteName)
	}
	for _, image := range b.images {
		mb.Include(image.meta(ns))
	}
	for _, video := range b.videos {
		mb.Include(video.meta(ns))
	}
	for _, audio := range b.audios {
		mb.Include(audio.meta(ns))
	}
	if b.duration > 0 {
		mb.Add(ns, "duration", b.duration)
	}
	if b.releaseDate != nil {
		mb.Add(ns, "release_date", b.releaseDate.Format(time.RFC3339))
	}
	for _, tag := range b.tags {
		mb.Add(ns, "tag", tag)
	}
	if ns != "og" {
		mb.Include(&b.actors)
		mb.Include(&b.directors)
		mb.Include(&b.writers)
	}
	return &mb
}
