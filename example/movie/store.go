package movie

type MovieStore struct {
	movies Movies
}

func NewMovieStore() *MovieStore {
	return &MovieStore{
		movies: Movies{},
	}
}

func (m *MovieStore) Movies() Movies {
	return m.movies
}

func (m *MovieStore) Insert(movie Movie) {
	m.movies = append(m.movies, movie)
}
