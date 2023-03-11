package structs

type Response struct {
	User *User
	Movie *Movie
	Director *Director
	Movies []Movie
	IsMovieWatched bool
	Message string
	Title string
}