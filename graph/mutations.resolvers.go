package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	"github.com/tryy3/Recipe-Maker-Backend/graph/generated"
	"github.com/tryy3/Recipe-Maker-Backend/graph/model"
	"github.com/tryy3/go-cloudinary"
)

func (r *mutationResolver) UpdateIngredient(ctx context.Context, id string, ingredient model.IngredientInput) (*model.Ingredient, error) {
	doc := r.Database.Doc("ingredients/" + id)

	var data map[string]interface{}
	dataJSON, err := json.Marshal(ingredient)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(dataJSON, &data)

	// Remove nil values from map
	for key, v := range data {
		if v == nil {
			delete(data, key)
		}
	}

	_, err = doc.Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return nil, err
	}

	snap, err := doc.Get(context.Background())
	if err != nil {
		return nil, err
	}

	ingredientModel := model.Ingredient{
		ID: doc.ID,
	}

	err = snap.DataTo(&ingredientModel)
	if err != nil {
		return nil, err
	}

	return &ingredientModel, nil
}

func (r *mutationResolver) CreateIngredient(ctx context.Context, ingredient model.IngredientInput) (*model.Ingredient, error) {
	col := r.Database.Collection("ingredients")

	var data map[string]interface{}
	dataJSON, err := json.Marshal(ingredient)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(dataJSON, &data)

	// Remove nil values from map
	for key, v := range data {
		if v == nil {
			delete(data, key)
		}
	}

	doc, _, err := col.Add(context.Background(), data)
	if err != nil {
		return nil, err
	}

	snap, err := doc.Get(context.Background())
	if err != nil {
		return nil, err
	}

	ingredientModel := model.Ingredient{
		ID: doc.ID,
	}

	err = snap.DataTo(&ingredientModel)
	if err != nil {
		return nil, err
	}

	return &ingredientModel, nil
}

func (r *mutationResolver) DeleteIngredient(ctx context.Context, id string) (bool, error) {
	doc := r.Database.Doc("ingredients/" + id)
	_, err := doc.Delete(context.Background())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateRecipe(ctx context.Context, id string, recipe model.RecipeInput) (*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRecipe(ctx context.Context, recipe model.RecipeInput) (*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRecipe(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddRecipeIngredient(ctx context.Context, id string, ingredient model.RecipeIngredientInput) (*model.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadFiles(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	var out = []*model.File{}
	for _, file := range files {
		//create destination file making sure the path is writeable.
		path, err := r.CloudinaryService.Upload(file.Filename, file.File, "", true, cloudinary.ImageType)
		if err != nil {
			return nil, err
		}

		out = append(out, &model.File{ID: path, Name: file.Filename})
	}
	return out, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
