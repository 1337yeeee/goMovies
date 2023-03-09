package structs

type Response struct {
	User *User
	Movie *Movie
	Producer *Producer
	Movies []Movie
	IsMovieWatched bool
	Message string
}