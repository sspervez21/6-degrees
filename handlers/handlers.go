package handlers

import (
	"CrossRiver/6-degrees/models"
	"CrossRiver/6-degrees/restapi/operations"
	"CrossRiver/6-degrees/sixdegrees"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
)

func GetDegreesOfSeparationHandler(params operations.GetDegreesOfSeparationParams) middleware.Responder {

	fmt.Println("HERE222")

	i, m, err := sixdegrees.GetKevinBaconDegrees(sixdegrees.DenormalizeName(params.ActorName))

	fmt.Println("HERE333")

	if err != nil {
		// TODO: specialize behavior for the user's benefit
		if err == sixdegrees.NotAnActorErr {
			return operations.NewGetDegreesOfSeparationNotFound().WithPayload(&models.ActorNotFound{Message: err.Error()})
		} else if err == sixdegrees.NotKevinsFriendErr {
			return operations.NewGetDegreesOfSeparationNotFound().WithPayload(&models.ActorNotFound{Message: err.Error()})
		} else {
			return operations.NewGetDegreesOfSeparationNotFound().WithPayload(&models.ActorNotFound{Message: err.Error()})
		}
	}

	if m == nil {
		// This should not happen...
		return operations.NewGetDegreesOfSeparationNotFound().WithPayload(&models.ActorNotFound{Message: "Internal server error. Please try again?"})
	}

	movies := make([]*models.Actor, 0)

	fmt.Println("HERE444")

	for {
		fmt.Println("HERE555::" + m.Movie.Title)

		movies = append(movies, &models.Actor{Movie: m.Movie.Title})
		m = m.OriginalMovie
		if m == nil {
			break
		}
	}

	return operations.NewGetDegreesOfSeparationOK().WithPayload(&models.ActorFound{DegreesOfSeparation: int64(i), Movies: movies})
}
