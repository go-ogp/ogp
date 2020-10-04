package ogp

// Image -----------------------------------------------------------------------

// ImageBuilder builds an `og:image` object.
type ImageBuilder struct {
	url       string
	secureURL string
	mime      string
	alt       string
	width     int
	height    int
}

// URL sets the `og:image:url` property.
func (b *ImageBuilder) URL(url string) *ImageBuilder {
	b.url = url
	return b
}

// SecureURL sets the `og:image:secure_url` property.
func (b *ImageBuilder) SecureURL(url string) *ImageBuilder {
	b.secureURL = url
	return b
}

// MIME sets the `og:image:type` property.
func (b *ImageBuilder) MIME(mime string) *ImageBuilder {
	b.mime = mime
	return b
}

// Alt sets the `og:image:alt` property.
func (b *ImageBuilder) Alt(alt string) *ImageBuilder {
	b.alt = alt
	return b
}

// Width sets the `og:image:width` property.
func (b *ImageBuilder) Width(width int) *ImageBuilder {
	b.width = width
	return b
}

// Height sets the `og:image:height` property.
func (b *ImageBuilder) Height(height int) *ImageBuilder {
	b.height = height
	return b
}

func (b *ImageBuilder) meta(ns string) *metaBuilder {
	var mb metaBuilder
	mb.Add(ns, "image", b.url)
	if b.secureURL != "" {
		mb.Add(ns, "image:secure_url", b.secureURL)
	}
	if b.mime != "" {
		mb.Add(ns, "image:type", b.mime)
	}
	if b.alt != "" {
		mb.Add(ns, "image:alt", b.alt)
	}
	if b.width > 0 {
		mb.Add(ns, "image:width", b.width)
	}
	if b.height > 0 {
		mb.Add(ns, "image:height", b.height)
	}
	return &mb
}

// Video -----------------------------------------------------------------------

// VideoBuilder builds an `og:video` object.
type VideoBuilder struct {
	url       string
	secureURL string
	mime      string
	alt       string
	width     int
	height    int
}

// URL sets the `og:video:url` property.
func (b *VideoBuilder) URL(url string) *VideoBuilder {
	b.url = url
	return b
}

// SecureURL sets the `og:video:secure_url` property.
func (b *VideoBuilder) SecureURL(url string) *VideoBuilder {
	b.secureURL = url
	return b
}

// MIME sets the `og:video:type` property.
func (b *VideoBuilder) MIME(mime string) *VideoBuilder {
	b.mime = mime
	return b
}

// Alt sets the `og:video:alt` property.
func (b *VideoBuilder) Alt(alt string) *VideoBuilder {
	b.alt = alt
	return b
}

// Width sets the `og:video:width` property.
func (b *VideoBuilder) Width(width int) *VideoBuilder {
	b.width = width
	return b
}

// Height sets the `og:video:height` property.
func (b *VideoBuilder) Height(height int) *VideoBuilder {
	b.height = height
	return b
}

func (b *VideoBuilder) meta(ns string) *metaBuilder {
	var mb metaBuilder
	mb.Add(ns, "video", b.url)
	if b.secureURL != "" {
		mb.Add(ns, "video:secure_url", b.secureURL)
	}
	if b.mime != "" {
		mb.Add(ns, "video:type", b.mime)
	}
	if b.alt != "" {
		mb.Add(ns, "video:alt", b.alt)
	}
	if b.width > 0 {
		mb.Add(ns, "video:width", b.width)
	}
	if b.height > 0 {
		mb.Add(ns, "video:height", b.height)
	}
	return &mb
}

// Audio -----------------------------------------------------------------------

// AudioBuilder builds an `og:audio` object.
type AudioBuilder struct {
	url       string
	secureURL string
	mime      string
}

// URL sets the `og:audio:url` property.
func (b *AudioBuilder) URL(url string) *AudioBuilder {
	b.url = url
	return b
}

// SecureURL sets the `og:audio:secure_url` property.
func (b *AudioBuilder) SecureURL(url string) *AudioBuilder {
	b.secureURL = url
	return b
}

// MIME sets the `og:audio:type` property.
func (b *AudioBuilder) MIME(mime string) *AudioBuilder {
	b.mime = mime
	return b
}

func (b *AudioBuilder) meta(ns string) *metaBuilder {
	var mb metaBuilder
	mb.Add(ns, "audio", b.url)
	if b.secureURL != "" {
		mb.Add(ns, "audio:secure_url", b.secureURL)
	}
	if b.mime != "" {
		mb.Add(ns, "audio:type", b.mime)
	}
	return &mb
}
