package sixdegrees

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetMovies(actorId int) ([]Movie, error) {
	url := strings.Replace(getMoviesURL, "<ActorID>", strconv.Itoa(actorId), 1)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+authToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))

	moviesRes := MovieResults{}
	err = json.Unmarshal(body, &moviesRes)
	if err != nil {
		return nil, err
	}

	if len(moviesRes.Cast) == 0 {
		return nil, errors.New("no movies found")
	}

	return moviesRes.Cast, nil
}
