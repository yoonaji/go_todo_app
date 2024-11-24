package service

import (
	"context"
	"fmt"

	"github.com/yoonaji/go_todo_app/7_week/entity"
	"github.com/yoonaji/go_todo_app/7_week/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
