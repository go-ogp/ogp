package ogp

import (
	"html/template"
)

// MusicPlaylistBuilder builds a `music.playlist` object.
type MusicPlaylistBuilder struct {
	WebsiteBuilder
	songs    metaBuilder
	creators metaBuilder
}

// Title sets the `music:title` property.
func (b *MusicPlaylistBuilder) Title(title string) *MusicPlaylistBuilder {
	b.title = title
	return b
}

// URL sets the `music:url` property.
func (b *MusicPlaylistBuilder) URL(url string) *MusicPlaylistBuilder {
	b.url = url
	return b
}

// Description sets the `music:description` property.
func (b *MusicPlaylistBuilder) Description(description string) *MusicPlaylistBuilder {
	b.description = description
	return b
}

// Determiner sets the `music:determiner` property.
func (b *MusicPlaylistBuilder) Determiner(determiner string) *MusicPlaylistBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `music:locale` or adds a new `music:locale:alternate` property.
func (b *MusicPlaylistBuilder) Locale(locale string) *MusicPlaylistBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `music:site_name` property.
func (b *MusicPlaylistBuilder) SiteName(siteName string) *MusicPlaylistBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `music:image` property.
func (b *MusicPlaylistBuilder) Image(image *ImageBuilder) *MusicPlaylistBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `music:video` property.
func (b *MusicPlaylistBuilder) Video(video *VideoBuilder) *MusicPlaylistBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `music:audio` property.
func (b *MusicPlaylistBuilder) Audio(audio *AudioBuilder) *MusicPlaylistBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// Song adds a new `music:song` property.
func (b *MusicPlaylistBuilder) Song(url string, disc, track int) *MusicPlaylistBuilder {
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

// Creator adds a new `music:creator` property.
func (b *MusicPlaylistBuilder) Creator(creator *ProfileBuilder) *MusicPlaylistBuilder {
	b.creators.Include(creator.meta("music:creator"))
	return b
}

// HTML renders the `music.playlist` object to be used in HTML templates.
func (b *MusicPlaylistBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *MusicPlaylistBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "music.playlist")
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
	mb.Include(&b.songs)
	mb.Include(&b.creators)
	return &mb
}
