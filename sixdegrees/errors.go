package sixdegrees

import "errors"

var NotKevinsFriendErr = errors.New("Congratulations, you found an actor not within 6 degrees of Kevin Bacon. How did you do this??")
var NotAnActorErr = errors.New("This person does not appear to be an actor")
