package suits

type Category struct {
	Id           int
	MenuType     string
	CategoryName string
}

var Categories = []Category{
	{Id: 1, MenuType: "Основное меню", CategoryName: "Салаты"},
	{Id: 2, MenuType: "Основное меню", CategoryName: "Закуски"},
	{Id: 3, MenuType: "Основное меню", CategoryName: "Супы"},
	{Id: 4, MenuType: "Кейтеринг", CategoryName: "Сеты"},
	{Id: 5, MenuType: "Банкетное меню", CategoryName: "Десерты"},
	{Id: 6, MenuType: "Кейтеринг", CategoryName: "Пицца"},
}
