package site

import "log/slog"

const baseHTMLLayout string = "./ui/html/site/base.layout.html"

type MainHandlers struct {
	log *slog.Logger
}

func NewMainHandlers(log *slog.Logger) *MainHandlers {
	return &MainHandlers{log: log}
}
