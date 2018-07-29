package timeofday

import(
"time"
"errors"
"fmt"
middleware "github.com/go-openapi/runtime/middleware"
"github.com/ahlusar1989/swagger-ui/models"
"github.com/ahlusar1989/swagger-ui/restapi/operations"
"github.com/go-openapi/swag"
)



func getTimeOfDay(tz *string) (*string, error) {
	defaultTZ := "UTC"

	t := time.Now()
	if tz == nil {
		tz = &defaultTZ
	}

	utc, err := time.LoadLocation(*tz)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Time zone not found: %s", *tz))
	}

	thetime := t.In(utc).String()
	return &thetime, nil
}


func GetClock(params operations.TimeGetParams) middleware.Responder {
	var tz *string = nil

	if (params.Timezone != nil) {
		tz = params.Timezone
	}

	thetime, err := getTimeOfDay(tz)
	if err != nil {
		return operations.NewTimeGetNotFound().WithPayload(
			&models.ErrorResponse {
				int32(operations.TimeGetNotFoundCode),
				swag.String(fmt.Sprintf("%s", err)),
			})
	}

	return operations.NewTimeGetOK().WithPayload(
		&models.Timeofday{
			Timeofday: *thetime,
		})
}


func PostClock(params operations.TimePostParams) middleware.Responder {
	var tz *string = nil

	if (params.Timezone != nil) {
		tz = params.Timezone.Timezone
	}

	thetime, err := getTimeOfDay(tz)
	if err != nil {
		return operations.NewTimePostNotFound().WithPayload(
			&models.ErrorResponse {
				int32(operations.TimePostNotFoundCode),
				swag.String(fmt.Sprintf("%s", err)),
			})
	}

	return operations.NewTimePostOK().WithPayload(
		&models.Timeofday{
			Timeofday: *thetime,
		})
}