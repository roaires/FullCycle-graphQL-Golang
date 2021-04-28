package graph

import "github.com/roaires/FullCycle-graphQL-Golang/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {

	/* Esse exemplo não irá trabalhar com banco de dados. Para simular serão utilizadas listas para armazenamento */

	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
