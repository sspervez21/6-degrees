package sixdegrees

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetCast(movieId int) ([]Cast, error) {
	url := strings.Replace(getCastURL, "<MovieID>", strconv.Itoa(movieId), 1)

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

	actorsRes := CastResults{}
	err = json.Unmarshal(body, &actorsRes)
	if err != nil {
		return nil, err
	}

	if len(actorsRes.Cast) == 0 {
		return nil, errors.New("no movies found")
	}

	return actorsRes.Cast, nil
}
