package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/roaires/FullCycle-graphQL-Golang/graph/generated"
	"github.com/roaires/FullCycle-graphQL-Golang/graph/model"
)

func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	// Implementação Lazy loading
	// Será executado apenas quando solicitado na query
	// Quando utilizar SQL/ORM verificar estatégias para tratar N+1

	var courses []*model.Course

	for _, c := range r.Resolver.Courses {
		if c.Category.ID == obj.ID {
			courses = append(courses, c)
		}
	}

	return courses, nil
}

func (r *courseResolver) Chapters(ctx context.Context, obj *model.Course) ([]*model.Chapter, error) {
	// Implementação Lazy loading
	// Será executado apenas quando solicitado na query
	// Quando utilizar SQL/ORM verificar estatégias para tratar N+1

	var chapters []*model.Chapter

	for _, c := range r.Resolver.Chapters {
		if c.Course.ID == obj.ID {
			chapters = append(chapters, c)
		}
	}

	return chapters, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := model.Category{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: input.Description,
	}

	r.Categories = append(r.Categories, &category)

	return &category, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	var category *model.Category

	// Simulando a busca de uma Category pelo id
	for _, c := range r.Categories {
		if c.ID == input.CategoryID {
			category = c
		}
	}

	course := model.Course{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: input.Description,
		Category:    category,
	}

	r.Courses = append(r.Courses, &course)

	return &course, nil
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	var course *model.Course

	// Simulando a busca de um Course pelo id
	for _, c := range r.Courses {
		if c.ID == input.CourseID {
			course = c
		}
	}

	chapter := model.Chapter{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Name:   input.Name,
		Course: course,
	}

	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Categoroes(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}
