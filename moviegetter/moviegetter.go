package moviegetter

type Movie struct {
	Year string
}

type MovieGetter struct {
	Engine string
}

func (m MovieGetter) GetMovie(movieName string) Movie {
	return Movie{"2001"}
}
