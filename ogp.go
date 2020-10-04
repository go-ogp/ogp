package ogp

// Website is the convenient way for creating a WebsiteBuilder.
func Website() *WebsiteBuilder {
	return &WebsiteBuilder{}
}

// Article is the convenient way for creating an ArticleBuilder.
func Article() *ArticleBuilder {
	return &ArticleBuilder{}
}

// Book is the convenient way for creating a BookBuilder.
func Book() *BookBuilder {
	return &BookBuilder{}
}

// Profile is the convenient way for creating a ProfileBuilder.
func Profile() *ProfileBuilder {
	return &ProfileBuilder{}
}

// Song is the convenient way for creating a MusicSongBuilder.
func Song() *MusicSongBuilder {
	return &MusicSongBuilder{}
}

// Album is the convenient way for creating a MusicAlbumBuilder.
func Album() *MusicAlbumBuilder {
	return &MusicAlbumBuilder{}
}

// Playlist is the convenient way for creating a MusicPlaylistBuilder.
func Playlist() *MusicPlaylistBuilder {
	return &MusicPlaylistBuilder{}
}

// RadioStation is the convenient way for creating a MusicRadioStationBuilder.
func RadioStation() *MusicRadioStationBuilder {
	return &MusicRadioStationBuilder{}
}

// Movie is the convenient way for creating a VideoMovieBuilder.
func Movie() *VideoMovieBuilder {
	return &VideoMovieBuilder{}
}

// TVShow is the convenient way for creating a VideoTVShowBuilder.
func TVShow() *VideoTVShowBuilder {
	return &VideoTVShowBuilder{}
}

// Episode is the convenient way for creating a VideoEpisodeBuilder.
func Episode() *VideoEpisodeBuilder {
	return &VideoEpisodeBuilder{}
}

// VideoOther is the convenient way for creating a VideoOtherBuilder.
func VideoOther() *VideoOtherBuilder {
	return &VideoOtherBuilder{}
}
