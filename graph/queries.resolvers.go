package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tryy3/Recipe-Maker-Backend/graph/generated"
	"github.com/tryy3/Recipe-Maker-Backend/graph/model"
)

func (r *queryResolver) Ingredients(ctx context.Context) ([]*model.Ingredient, error) {
	col := r.Database.Collection("ingredients")
	refs, err := col.DocumentRefs(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}

	ingredients := make([]*model.Ingredient, len(refs))

	for i, ref := range refs {
		doc, err := ref.Get(context.Background())
		if err != nil {
			return nil, err
		}

		ingredient := model.Ingredient{
			ID: ref.ID,
		}
		err = doc.DataTo(&ingredient)
		if err != nil {
			return nil, err
		}

		ingredients[i] = &ingredient
	}

	return ingredients, nil
}

func (r *queryResolver) Ingredient(ctx context.Context, id string) (*model.Ingredient, error) {
	doc := r.Database.Doc("ingredients/" + id)
	snap, err := doc.Get(context.Background())
	if err != nil {
		return nil, err
	}

	ingredient := model.Ingredient{
		ID: doc.ID,
	}

	err = snap.DataTo(&ingredient)
	if err != nil {
		return nil, err
	}

	return &ingredient, nil
}

func (r *queryResolver) Recipes(ctx context.Context) ([]*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Recipe(ctx context.Context, id string) (*model.Recipe, error) {
	doc := r.Database.Doc("recipes/" + id)
	snap, err := doc.Get(context.Background())
	if err != nil {
		return nil, err
	}

	recipe := model.Recipe{
		ID: doc.ID,
	}

	err = snap.DataTo(&recipe)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
