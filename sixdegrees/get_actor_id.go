package sixdegrees

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func GetActorId(actorName string) (int, error) {
	url := strings.Replace(getActorURL, "<ActorName>", NormalizeName(actorName), 1)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+authToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	//fmt.Println(string(body))

	actorRes := ActorResults{}
	err = json.Unmarshal(body, &actorRes)
	if err != nil {
		return 0, err
	}

	if len(actorRes.Results) == 0 {
		return 0, errors.New("no actors found")
	}

	for _, actor := range actorRes.Results {
		if actor.KnownForDepartment == "Acting" {
			return actor.ID, nil
		}
	}

	return 0, NotAnActorErr
}
