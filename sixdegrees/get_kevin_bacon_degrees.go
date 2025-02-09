package sixdegrees

import (
	"fmt"
	"strconv"
)

func GetKevinBaconDegrees(actorName string) (int, *MovieExtended, error) {
	_, err := GetActorId(actorName)
	if err != nil {
		return 0, nil, err
	}

	id, err := GetActorId("Kevin+Bacon")
	if err != nil {
		return 0, nil, err
	}

	tempMovies, err := GetMovies(id)
	if err != nil {
		return 0, nil, err
	}

	KevinBaconMovies := make([]MovieExtended, 0)

	for _, movie := range tempMovies {
		KevinBaconMovies = append(KevinBaconMovies, MovieExtended{Movie: movie, OriginalMovie: nil})
		fmt.Println(strconv.Itoa(movie.ID) + "::" + movie.Title)
	}

	for i := range 6 {
		fmt.Println(strconv.Itoa(i) + "::" + strconv.Itoa(len(KevinBaconMovies)))
		j := 1
		NextMovies := make([]MovieExtended, 0)
		for _, movie := range KevinBaconMovies {
			fmt.Println(strconv.Itoa(j) + "::GetCast:" + movie.Movie.Title)
			j++
			cast, err := GetCast(movie.Movie.ID)
			if err != nil {
				return 0, nil, err
			}

			//if len(cast) > 10 {
			//cast = cast[:10]
			//}

			for _, actor := range cast {
				if actor.Name == actorName {
					return i + 1, &movie, nil
				}

				actorMovies, err := GetMovies(actor.ID)
				if err != nil {
					return 0, nil, err
				}
				for _, actorMovie := range actorMovies {
					NextMovies = append(NextMovies, MovieExtended{Movie: actorMovie, OriginalMovie: &movie})
				}
			}
		}
		KevinBaconMovies = NextMovies
	}

	return 0, nil, NotKevinsFriendErr
}
