### Prerequisites

Golang, Docker and Docker-Compose

# Build

```
make build && make image && make push
```

## Deployment

```
make test-env
```

## Running the tests
There are 2 API in the service, the exercise_for_muscle and muscles_for_exercise APIs, the API takes "AccountID={accound_id}" as the authorization header. For now, this is used as simple authorization. Authorization middleware can be used in the future to implement real authorization.

For exercise_for_muscle API, the API takes an additional parameter of a muscle group name, the muscle group name can be one of the following names:
```
calves
quad_riceps
ham_strings
gluteus
hips_other
lower_back
lats
trapezius
abdominals
pectorals
deltoid
triceps
biceps
forearms
```

For muscles_for_exercise API, the API takes an additional parameter of an exercise name, the exercise name can be one of the following names:
```
squat
leg_press
lunge
deadlift
leg_extension
leg_curl
standing_calf_raise
seated_calf_raise
hip_adductor
bench_press
chest_fly
push_up
pull_down
pull_up
bent_over_row
upright_row
shoulder_press
shoulder_fly
lateral_raise
shoulder_shrug
pushdown
triceps_extension
biceps_curl
crunch
russian_twist
leg_raise
back_extension
```

```
curl -H "Authorization: AccountID=1" http://localhost:8081/api/exercise-service/v1/exercise_for_muscle/calves | jq
{
  "results": [
    "squat",
    "leg_press",
    "deadlift",
    "leg_curl",
    "standing_calf_raise",
    "seated_calf_raise"
  ]
}

curl -H "Authorization: AccountID=1" http://localhost:8081/api/exercise-service/v1/muscles_for_exercise/squat | jq
{
  "results": [
    "calves",
    "quad_riceps",
    "ham_strings",
    "gluteus",
    "hips_other",
    "lower_back",
    "abdominals"
  ]
}
```

