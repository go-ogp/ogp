package ogp

import (
	"html/template"
	"time"
)

// MusicAlbumBuilder builds a `music.album` object.
type MusicAlbumBuilder struct {
	WebsiteBuilder
	releaseDate *time.Time
	songs       metaBuilder
	musicians   metaBuilder
}

// Title sets the `music:title` property.
func (b *MusicAlbumBuilder) Title(title string) *MusicAlbumBuilder {
	b.title = title
	return b
}

// URL sets the `music:url` property.
func (b *MusicAlbumBuilder) URL(url string) *MusicAlbumBuilder {
	b.url = url
	return b
}

// Description sets the `music:description` property.
func (b *MusicAlbumBuilder) Description(description string) *MusicAlbumBuilder {
	b.description = description
	return b
}

// Determiner sets the `music:determiner` property.
func (b *MusicAlbumBuilder) Determiner(determiner string) *MusicAlbumBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `music:locale` or adds a new `music:locale:alternate` property.
func (b *MusicAlbumBuilder) Locale(locale string) *MusicAlbumBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `music:site_name` property.
func (b *MusicAlbumBuilder) SiteName(siteName string) *MusicAlbumBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `music:image` property.
func (b *MusicAlbumBuilder) Image(image *ImageBuilder) *MusicAlbumBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `music:video` property.
func (b *MusicAlbumBuilder) Video(video *VideoBuilder) *MusicAlbumBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `music:audio` property.
func (b *MusicAlbumBuilder) Audio(audio *AudioBuilder) *MusicAlbumBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// ReleaseDate sets the `music:release_date` property.
func (b *MusicAlbumBuilder) ReleaseDate(releaseDate time.Time) *MusicAlbumBuilder {
	b.releaseDate = &releaseDate
	return b
}

// Song adds a new `music:song` property.
func (b *MusicAlbumBuilder) Song(url string, disc, track int) *MusicAlbumBuilder {
	var mb metaBuilder
	if url != "" {
		mb.Add("music:song", "", url)
	}
	if disc > 0 {
		mb.Add("music:song", "disc", disc)
	}
	if track > 0 {
		mb.Add("music:song", "track", track)
	}
	b.songs.Include(&mb)
	return b
}

// Musician adds a new `music:musician` property.
func (b *MusicAlbumBuilder) Musician(musician *ProfileBuilder) *MusicAlbumBuilder {
	b.musicians.Include(musician.meta("music:musician"))
	return b
}

// HTML renders the `music.album` object to be used in HTML templates.
func (b *MusicAlbumBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *MusicAlbumBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "music.album")
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
	if b.releaseDate != nil {
		mb.Add("music", "release_date", b.releaseDate.Format(time.RFC3339))
	}
	mb.Include(&b.songs)
	mb.Include(&b.musicians)
	return &mb
}
