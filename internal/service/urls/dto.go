package urls

type ShortenUrlIn struct {
	OriginalURL string
}

type ShortenURLOut struct {
	Alias string
}

type GetOriginalURLIn struct {
	ShortURL string
}

type GetOriginalURlOut struct {
	OriginalURL string
}
