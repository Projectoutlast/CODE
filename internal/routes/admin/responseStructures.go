package admin

import "code/internal/repository/models"

type MenuResponse struct {
	PageTitle string
	MenuTypes []models.Menu
}
