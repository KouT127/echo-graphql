package graphql

import (
	"context"
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
	"github.com/KouT127/gin-sample/backend/util"
	"strconv"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUser(ctx context.Context, user graph.UserInput) (*graph.AddUserPayload, error) {
	panic("implement me")
}

func (r *mutationResolver) AddTask(ctx context.Context, input graph.TaskInput) (*graph.AddTaskPayload, error) {
	panic("implement me")
}

type queryResolver struct{ *Resolver }

type results struct {
	model.User
	model.Task
}

func (r *queryResolver) User(ctx context.Context, id *string) (*graph.User, error) {
	db := database.NewDB()
	user := model.User{}
	var u = &graph.User{}
	var tl []*graph.Task
	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		panic(err)
	}
	query := db.Table("users").
		Select("users.id, users.name, users.gender, tasks.id, tasks.title, tasks.description").
		Joins("left outer join tasks on users.id = tasks.user_refer").
		Where("users.id = ?", id).
		Offset(1).
		Limit(10)
	rows, err := query.Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		res := &results{}
		err := db.ScanRows(rows, res)
		if err != nil {
			panic(err)
		}
		taskId := strconv.Itoa(int(res.Task.ID))
		gt := graph.Task{
			ID:          taskId,
			Title:       res.Task.Title,
			Description: res.Task.Description,
		}
		tl = append(tl, &gt)
	}
	usrId := strconv.Itoa(int(user.ID))
	u = &graph.User{
		ID:     usrId,
		Name:   user.Name,
		Gender: user.Gender,
		Tasks:  tl,
	}
	return u, nil
}

func (r *queryResolver) AllTasks(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	db := database.NewDB()
	var cnt int
	err := db.Model(&model.Task{}).Count(&cnt).Error
	if err != nil {
		panic(err)
	}
	rows, err := db.Model(&model.Task{}).Rows()
	if err != nil {
		panic(err)
	}
	var edges []*graph.TaskEdge
	for rows.Next() {
		task := model.Task{}
		err := db.ScanRows(rows, &task)
		if err != nil {
			panic(err)
		}
		id := strconv.Itoa(int(task.ID))
		t := &graph.Task{
			ID:          id,
			Title:       task.Title,
			Description: task.Description,
		}
		cur := util.Base64Encode(t.ID)
		edge := graph.TaskEdge{
			Cursor: cur,
			Node:   t,
		}
		edges = append(edges, &edge)
	}
	startcur := edges[0].Cursor
	endcur := edges[len(edges)-1].Cursor
	pg := &graph.PageInfo{
		StartCursor: &startcur,
		EndCursor:   &endcur,
		HasNextPage: true,
	}
	con := &graph.TaskConnection{
		TotalCount: cnt,
		Edges:      edges,
		PageInfo:   pg,
	}
	return con, nil
}
