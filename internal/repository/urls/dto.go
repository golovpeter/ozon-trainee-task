package urls

type ShortenUrlIn struct {
	OriginalURL string
	Alias       string
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
