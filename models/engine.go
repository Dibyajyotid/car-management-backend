package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json: "engine_id"`
	Displacement  int64     `json: "displacement"`
	NoOfCylinders int64     `json: "no_of_cylinders"`
	CarRange      int64     `json: "carRange"`
}

type EngineRequest struct {
	Displacement  int64 `json: "displacement"`
	NoOfCylinders int64 `json: "no_of_cylinders"`
	CarRange      int64 `json: "carRange"`
}

func validateEgineRequest(engineReq EngineRequest) error {
	if err := validateDisplacement(engineReq.Displacement); err != nil {
		return err
	}

	if err := validateNoOfCylinders(engineReq.NoOfCylinders); err != nil {
		return err
	}

	if err := validateCarRange(engineReq.CarRange); err != nil {
		return err
	}

	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement < 0 {
		return errors.New("Displacement must be greater than zero")
	}

	return nil
}

func validateNoOfCylinders(niOfCylinders int64) error {
	if niOfCylinders <= 0 {
		return errors.New("No of clinders must be greater than zero")
	}

	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("Car range must be greater than zero")
	}

	return nil
}
