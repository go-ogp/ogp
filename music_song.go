package ogp

import (
	"html/template"
)

// MusicSongBuilder builds a `music.song` object.
type MusicSongBuilder struct {
	WebsiteBuilder
	duration  int
	albums    metaBuilder
	musicians metaBuilder
}

// Title sets the `music:title` property.
func (b *MusicSongBuilder) Title(title string) *MusicSongBuilder {
	b.title = title
	return b
}

// URL sets the `music:url` property.
func (b *MusicSongBuilder) URL(url string) *MusicSongBuilder {
	b.url = url
	return b
}

// Description sets the `music:description` property.
func (b *MusicSongBuilder) Description(description string) *MusicSongBuilder {
	b.description = description
	return b
}

// Determiner sets the `music:determiner` property.
func (b *MusicSongBuilder) Determiner(determiner string) *MusicSongBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `music:locale` or adds a new `music:locale:alternate` property.
func (b *MusicSongBuilder) Locale(locale string) *MusicSongBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `music:site_name` property.
func (b *MusicSongBuilder) SiteName(siteName string) *MusicSongBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `music:image` property.
func (b *MusicSongBuilder) Image(image *ImageBuilder) *MusicSongBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `music:video` property.
func (b *MusicSongBuilder) Video(video *VideoBuilder) *MusicSongBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `music:audio` property.
func (b *MusicSongBuilder) Audio(audio *AudioBuilder) *MusicSongBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// Duration sets the `music:duration` property.
func (b *MusicSongBuilder) Duration(duration int) *MusicSongBuilder {
	b.duration = duration
	return b
}

// Album adds a new `music:album` property.
func (b *MusicSongBuilder) Album(url string, disc, track int) *MusicSongBuilder {
	var mb metaBuilder
	if url != "" {
		mb.Add("music:album", "", url)
	}
	if disc > 0 {
		mb.Add("music:album", "disc", disc)
	}
	if track > 0 {
		mb.Add("music:album", "track", track)
	}
	b.albums.Include(&mb)
	return b
}

// Musician adds a new `music:musician` property.
func (b *MusicSongBuilder) Musician(musician *ProfileBuilder) *MusicSongBuilder {
	b.musicians.Include(musician.meta("music:musician"))
	return b
}

// HTML renders the `music.song` object to be used in HTML templates.
func (b *MusicSongBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *MusicSongBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "music.song")
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
		mb.Add("music", "duration", b.duration)
	}
	mb.Include(&b.albums)
	mb.Include(&b.musicians)
	return &mb
}
