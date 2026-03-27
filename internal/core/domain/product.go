package domain

import (
	"fmt"
	"time"

	core_errors "github.com/WilliardT/go-mvp/internal/core/errors"
)

type Product struct {
	ID      int
	Version int

	Title       string
	Description *string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      int
}

func NewProduct(
	id int,
	version int,
	title string,
	description *string,
	price float64,
	createdAt time.Time,
	updatedAt time.Time,
	userID int,
) Product {
	return Product{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Price:       price,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		UserID:      userID,
	}
}

func NewProductUninitialized(
	title string,
	description *string,
	price float64,
	userID int,
) Product {
	return NewProduct(
		UninitializedID,
		UninitializedVersion,
		title,
		description,
		price,
		time.Time{},
		time.Time{},
		userID,
	)
}

func (p *Product) Validate() error {
	titleLength := len([]rune(p.Title))

	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf(
			"invalid `Title` len: %d: %w",
			titleLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Description != nil {
		descriptionLength := len([]rune(*p.Description))

		if descriptionLength < 1 || descriptionLength > 100 {
			return fmt.Errorf(
				"invalid `Description` len: %d: %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if p.Price <= 0 {
		return fmt.Errorf(
			"invalid `Price`: %f: %w",
			p.Price,
			core_errors.ErrInvalidArgument,
		)
	}

	if p.UserID <= 0 {
		return fmt.Errorf(
			"invalid `UserID`: %d: %w",
			p.UserID,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type ProductPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Price       Nullable[float64]
	UserID      Nullable[int]
}

func NewProductPatch(
	title Nullable[string],
	description Nullable[string],
	price Nullable[float64],
	userID Nullable[int],
) ProductPatch {
	return ProductPatch{
		Title:       title,
		Description: description,
		Price:       price,
		UserID:      userID,
	}
}

func (p *ProductPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf(
			"Title can`t be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Price.Set && p.Price.Value == nil {
		return fmt.Errorf(
			"Price can`t be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.UserID.Set && p.UserID.Value == nil {
		return fmt.Errorf(
			"UserID can`t be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Title.Set && p.Title.Value != nil {
		titleLength := len([]rune(*p.Title.Value))

		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf(
				"invalid `Title` len: %d: %w",
				titleLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if p.Description.Set && p.Description.Value != nil {
		descriptionLength := len([]rune(*p.Description.Value))

		if descriptionLength < 1 || descriptionLength > 100 {
			return fmt.Errorf(
				"invalid `Description` len: %d: %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if p.Price.Set && p.Price.Value != nil && *p.Price.Value <= 0 {
		return fmt.Errorf(
			"invalid `Price`: %f: %w",
			*p.Price.Value,
			core_errors.ErrInvalidArgument,
		)
	}

	if p.UserID.Set && p.UserID.Value != nil && *p.UserID.Value <= 0 {
		return fmt.Errorf(
			"invalid `UserID`: %d: %w",
			*p.UserID.Value,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (p *Product) ApplyPatch(patch ProductPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate product patch: %w", err)
	}

	tmp := *p

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Price.Set {
		tmp.Price = *patch.Price.Value
	}

	if patch.UserID.Set {
		tmp.UserID = *patch.UserID.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate product after applying patch: %w", err)
	}

	*p = tmp

	return nil
}
