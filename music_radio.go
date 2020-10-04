package ogp

import "html/template"

// MusicRadioStationBuilder builds a `music.radio_station` object.
type MusicRadioStationBuilder struct {
	WebsiteBuilder
	creators metaBuilder
}

// Title sets the `music:title` property.
func (b *MusicRadioStationBuilder) Title(title string) *MusicRadioStationBuilder {
	b.title = title
	return b
}

// URL sets the `music:url` property.
func (b *MusicRadioStationBuilder) URL(url string) *MusicRadioStationBuilder {
	b.url = url
	return b
}

// Description sets the `music:description` property.
func (b *MusicRadioStationBuilder) Description(description string) *MusicRadioStationBuilder {
	b.description = description
	return b
}

// Determiner sets the `music:determiner` property.
func (b *MusicRadioStationBuilder) Determiner(determiner string) *MusicRadioStationBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `music:locale` or adds a new `music:locale:alternate` property.
func (b *MusicRadioStationBuilder) Locale(locale string) *MusicRadioStationBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `music:site_name` property.
func (b *MusicRadioStationBuilder) SiteName(siteName string) *MusicRadioStationBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `music:image` property.
func (b *MusicRadioStationBuilder) Image(image *ImageBuilder) *MusicRadioStationBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `music:video` property.
func (b *MusicRadioStationBuilder) Video(video *VideoBuilder) *MusicRadioStationBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `music:audio` property.
func (b *MusicRadioStationBuilder) Audio(audio *AudioBuilder) *MusicRadioStationBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// Creator adds a new `music:creator` property.
func (b *MusicRadioStationBuilder) Creator(creator *ProfileBuilder) *MusicRadioStationBuilder {
	b.creators.Include(creator.meta("music:creator"))
	return b
}

// HTML renders the `music.radio_station` object to be used in HTML templates.
func (b *MusicRadioStationBuilder) HTML() template.HTML {
	return b.meta().HTML()
}

func (b *MusicRadioStationBuilder) meta() *metaBuilder {
	var mb metaBuilder
	mb.Add("og", "type", "music.radio_station")
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
	mb.Include(&b.creators)
	return &mb
}
