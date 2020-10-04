package ogp

import "html/template"

// ProfileBuilder builds a `profile` object.
type ProfileBuilder struct {
	WebsiteBuilder
	firstName string
	lastName  string
	username  string
	gender    string
}

// Title sets the `profile:title` property.
func (b *ProfileBuilder) Title(title string) *ProfileBuilder {
	b.title = title
	return b
}

// URL sets the `profile:url` property.
func (b *ProfileBuilder) URL(url string) *ProfileBuilder {
	b.url = url
	return b
}

// Description sets the `profile:description` property.
func (b *ProfileBuilder) Description(description string) *ProfileBuilder {
	b.description = description
	return b
}

// Determiner sets the `profile:determiner` property.
func (b *ProfileBuilder) Determiner(determiner string) *ProfileBuilder {
	b.determiner = determiner
	return b
}

// Locale sets the `profile:locale` or adds a new `profile:locale:alternate` property.
func (b *ProfileBuilder) Locale(locale string) *ProfileBuilder {
	b.locales = append(b.locales, locale)
	return b
}

// SiteName sets the `profile:site_name` property.
func (b *ProfileBuilder) SiteName(siteName string) *ProfileBuilder {
	b.siteName = siteName
	return b
}

// Image adds a new `profile:image` property.
func (b *ProfileBuilder) Image(image *ImageBuilder) *ProfileBuilder {
	b.images = append(b.images, image)
	return b
}

// Video adds a new `profile:video` property.
func (b *ProfileBuilder) Video(video *VideoBuilder) *ProfileBuilder {
	b.videos = append(b.videos, video)
	return b
}

// Audio adds a new `profile:audio` property.
func (b *ProfileBuilder) Audio(audio *AudioBuilder) *ProfileBuilder {
	b.audios = append(b.audios, audio)
	return b
}

// FirstName sets the `profile:first_name` property.
func (b *ProfileBuilder) FirstName(firstName string) *ProfileBuilder {
	b.firstName = firstName
	return b
}

// LastName sets the `profile:last_name` property.
func (b *ProfileBuilder) LastName(lastName string) *ProfileBuilder {
	b.lastName = lastName
	return b
}

// Username sets the `profile:username` property.
func (b *ProfileBuilder) Username(username string) *ProfileBuilder {
	b.username = username
	return b
}

// Gender sets the `profile:gender` property.
func (b *ProfileBuilder) Gender(gender string) *ProfileBuilder {
	b.gender = gender
	return b
}

// HTML renders the `profile` object to be used in HTML templates.
func (b *ProfileBuilder) HTML() template.HTML {
	return b.meta("og").HTML()
}

func (b *ProfileBuilder) meta(ns string) *metaBuilder {
	var mb metaBuilder
	if ns == "og" {
		mb.Add(ns, "type", "profile")
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
	if b.firstName != "" {
		if ns == "og" {
			mb.Add("profile", "first_name", b.url)
		} else {
			mb.Add(ns, "first_name", b.url)
		}
	}
	if b.lastName != "" {
		if ns == "og" {
			mb.Add("profile", "last_name", b.url)
		} else {
			mb.Add(ns, "last_name", b.url)
		}
	}
	if b.username != "" {
		if ns == "og" {
			mb.Add("profile", "username", b.url)
		} else {
			mb.Add(ns, "username", b.url)
		}
	}
	if b.gender != "" {
		if ns == "og" {
			mb.Add("profile", "gender", b.url)
		} else {
			mb.Add(ns, "gender", b.url)
		}
	}
	return &mb
}
