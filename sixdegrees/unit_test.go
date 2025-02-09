package sixdegrees

import (
	"testing"
)

func TestGetKevinBaconId(t *testing.T) {
	id, err := GetActorId("Kevin Bacon")

	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	if id != 4724 {
		t.Fatalf("Wrong Id: %d\n", id)
	}
}

func TestGetFakeActorId(t *testing.T) {
	_, err := GetActorId("Salman Pervez")

	if err == nil {
		t.Fatalf("Error expected\n")
	}
}

func TestGetKevinBaconMovies(t *testing.T) {
	id, err := GetActorId("Kevin Bacon")
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	movies, err := GetMovies(id)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	// Warning: Kevin Bacon is still making movie. Replace with someone who is not.
	if len(movies) != 117 {
		t.Fatalf("Wrong # Movies: %d\n", len(movies))
	}
}

func TestMysticRiverCast(t *testing.T) {
	// Mystic River movie id is 322
	cast, err := GetCast(322)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	if len(cast) != 59 {
		t.Fatalf("Wrong # Movies: %d\n", len(cast))
	}
}

func TestKevinBaconDegrees1(t *testing.T) {
	i, _, err := GetKevinBaconDegrees("Sean Penn")
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	if i != 1 {
		t.Fatalf("Sean Penn not found within 1 degree of separation from Kevin Bacon %d\n", i)
	}

	//fmt.Println(m.movie.Title)
}

func TestKevinBaconDegrees2(t *testing.T) {
	i, _, err := GetKevinBaconDegrees("Dennis Quaid")
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	if i != 2 {
		t.Fatalf("Dennis Quaid not found within 2 degrees of separation from Kevin Bacon %d\n", i)
	}

	//fmt.Println(m.movie.Title)
}
