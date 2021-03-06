// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Ingredient struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
}

type IngredientInput struct {
	Description *string `json:"description"`
	Title       *string `json:"title"`
	Image       *string `json:"image"`
}

type Recipe struct {
	ID          string              `json:"id"`
	Title       *string             `json:"title"`
	Description *string             `json:"description"`
	Image       *string             `json:"image"`
	Ingredients []*RecipeIngredient `json:"ingredients"`
}

type RecipeIngredient struct {
	ID               string           `json:"id"`
	Title            *string          `json:"title"`
	Description      *string          `json:"description"`
	Image            *string          `json:"image"`
	MeasurementType  *MeasurementType `json:"measurementType"`
	MeasurementValue *float64         `json:"measurementValue"`
}

type RecipeIngredientInput struct {
	ID string `json:"id"`
}

type RecipeInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type MeasurementType string

const (
	MeasurementTypeTeaspoon    MeasurementType = "TEASPOON"
	MeasurementTypeTablespoon  MeasurementType = "TABLESPOON"
	MeasurementTypeFluidounce  MeasurementType = "FLUIDOUNCE"
	MeasurementTypeGill        MeasurementType = "GILL"
	MeasurementTypeCup         MeasurementType = "CUP"
	MeasurementTypePint        MeasurementType = "PINT"
	MeasurementTypeQuart       MeasurementType = "QUART"
	MeasurementTypeMilliliter  MeasurementType = "MILLILITER"
	MeasurementTypeCenterliter MeasurementType = "CENTERLITER"
	MeasurementTypeDeciliter   MeasurementType = "DECILITER"
	MeasurementTypeLiter       MeasurementType = "LITER"
	MeasurementTypePound       MeasurementType = "POUND"
	MeasurementTypeOunce       MeasurementType = "OUNCE"
	MeasurementTypeMilligram   MeasurementType = "MILLIGRAM"
	MeasurementTypeGram        MeasurementType = "GRAM"
	MeasurementTypeKilogram    MeasurementType = "KILOGRAM"
)

var AllMeasurementType = []MeasurementType{
	MeasurementTypeTeaspoon,
	MeasurementTypeTablespoon,
	MeasurementTypeFluidounce,
	MeasurementTypeGill,
	MeasurementTypeCup,
	MeasurementTypePint,
	MeasurementTypeQuart,
	MeasurementTypeMilliliter,
	MeasurementTypeCenterliter,
	MeasurementTypeDeciliter,
	MeasurementTypeLiter,
	MeasurementTypePound,
	MeasurementTypeOunce,
	MeasurementTypeMilligram,
	MeasurementTypeGram,
	MeasurementTypeKilogram,
}

func (e MeasurementType) IsValid() bool {
	switch e {
	case MeasurementTypeTeaspoon, MeasurementTypeTablespoon, MeasurementTypeFluidounce, MeasurementTypeGill, MeasurementTypeCup, MeasurementTypePint, MeasurementTypeQuart, MeasurementTypeMilliliter, MeasurementTypeCenterliter, MeasurementTypeDeciliter, MeasurementTypeLiter, MeasurementTypePound, MeasurementTypeOunce, MeasurementTypeMilligram, MeasurementTypeGram, MeasurementTypeKilogram:
		return true
	}
	return false
}

func (e MeasurementType) String() string {
	return string(e)
}

func (e *MeasurementType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeasurementType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeasurementType", str)
	}
	return nil
}

func (e MeasurementType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
