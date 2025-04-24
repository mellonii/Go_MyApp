package main

import (
	"context"
	"fmt"
)

// Структура приложения
type App struct {
	ctx context.Context
}

// NewApp создает новую структуру приложения
func NewApp() *App {
	return &App{}
}

// startup вызывается при запуске приложения. context сохраняется
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet возвращает приветствие для имени
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
