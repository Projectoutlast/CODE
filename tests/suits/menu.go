package suits

type Menu struct {
	Id       int
	MenuType string
}

var MenuTypes = []Menu{
	{Id: 1, MenuType: "Основное меню"},
	{Id: 2, MenuType: "Меню доставки"},
	{Id: 3, MenuType: "Кейтеринг"},
	{Id: 4, MenuType: "Банкетное меню"},
	{Id: 5, MenuType: "Детское меню"},
}
