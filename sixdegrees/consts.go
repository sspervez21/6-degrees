package sixdegrees

const authToken = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI4NTBmMjY1ZTFiNjhmZDQ3MmIyNTdlMzQ1NGU2NjkyYyIsIm5iZiI6MTczOTA1MDkwMC43NzIsInN1YiI6IjY3YTdjZjk0YmYwYTlmNmY1NGUwODBhNyIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.xtvk8XLGeUomkKerUS8rk28CXvyCMZJ1qEdJ0bKWk6Y"

const getActorURL = "https://api.themoviedb.org/3/search/person?query=<ActorName>&include_adult=false&language=en-US&page=1"

// const getMoviesURL = "https://api.themoviedb.org/3/person/<ActorID>/combined_credits?language=en-US"
const getMoviesURL = "https://api.themoviedb.org/3/person/<ActorID>/movie_credits?language=en-US"

const getCastURL = "https://api.themoviedb.org/3/movie/<MovieID>>/credits?language=en-US"
