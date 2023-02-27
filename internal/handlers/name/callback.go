package name

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Callback = "name"

func ProcessCallback(context *appcontext.Context, data string) {
	rand.Seed(time.Now().UTC().UnixNano())
	msgId := context.RawUpdate.CallbackQuery.Message.MessageID

	msg := tgbotapi.NewEditMessageText(context.ChatID, msgId, "")
	switch data {
	case "female":
		index := rand.Intn(len(femaleNames))
		msg.Text = femaleNames[index]
	case "male":
		index := rand.Intn(len(maleNames))
		msg.Text = maleNames[index]
	}

	context.CustomAnswer(msg)
}

var femaleNames = []string{
	"Авдотья", "Агафья", "Аида", "Аксинья", "Алевтина", "Александра", "Алёна", "Алина", "Алла", "Анастасия",
	"Ангелина", "Анисья", "Анна", "Антонина", "Анфиса", "Аполлинария", "Арина", "Ария", "Ася", "Аэлита",
	"Богдана", "Валентина", "Валерия", "Варвара", "Василина", "Василиса", "Венера", "Вера", "Вета",
	"Викторина", "Виктория", "Вилена", "Влада", "Владана", "Владислава", "Галина", "Глафира", "Дарья",
	"Дина", "Домника", "Евгения", "Евдокия", "Екатерина", "Елена", "Елизавета", "Есения", "Зинаида",
	"Злата", "Зоя", "Изабелла", "Илина", "Иллирика", "Илья", "Инесса", "Инна", "Иоанна", "Ира", "Ираида",
	"Ирина", "Искра", "Ия", "Карина", "Кира", "Кристина", "Ксения", "Кузьма", "Лада", "Лара", "Лариса", "Лера",
	"Лидия", "Лика", "Лина", "Лукерья", "Людмила", "Ляля", "Магдалeна", "Майя", "Макария", "Маргарита", "Марина",
	"Мария", "Марфа", "Мила", "Милада", "Милана", "Милена", "Милица", "Мира", "Мирослава", "Мирра", "Надежда",
	"Наталья", "Нелли", "Ника", "Нина", "Нонна", "Оксана", "Октябрина", "Олеся", "Ольга",
	"Павлина", "Пелагея", "Платонида", "Полина", "Прасковья", "Рада", "Раиса", "Рената", "Римма", "Русалина",
	"Руслана", "Сабина", "Савва", "Светлана", "Серафима", "Соня", "Софья", "Стелла", "Таисия", "Татьяна", "Таяна",
	"Ульяна", "Устинья", "Фаина", "Федора", "Цветана", "Юлия", "Юния", "Яна", "Янина", "Ярина", "Ярослава", "Виталина (Виталия)",
}
var maleNames = []string{
	"Аким", "Александр", "Алексей", "Анатолий", "Андрей", "Антон", "Анфим",
	"Аркадий", "Арсений", "Артём", "Артемий", "Богдан", "Борис", "Борислав",
	"Вадим", "Валентин", "Валерий", "Василий", "Виктор", "Виталий", "Влад",
	"Владимир", "Владислав", "Владлен", "Влас", "Всеволод", "Вячеслав", "Гавриил",
	"Геннадий", "Георгий", "Герасим", "Глеб", "Гордей", "Григорий", "Дамир", "Даниил",
	"Данил", "Данислав", "Демид", "Демьян", "Денис", "Джереми (Иеремия)", "Дмитрий",
	"Евгений", "Евдоким", "Евстахий", "Егор", "Елисей", "Емельян", "Еремей", "Ефим", "Захар",
	"Зиновий", "Иван", "Игнат", "Игнатий", "Игорь", "Иероним (Джером)", "Иннокентий", "Кирилл",
	"Константин", "Лев", "Леонид", "Любовь", "Макар", "Макарий", "Максим", "Марк",
	"Мартин (Мартын)", "Матвей", "Милан", "Мирослав", "Михаил", "Никодим", "Николай", "Нинель",
	"Олег", "Осип (Иосиф)", "Остап", "Павел", "Пантелеймон", "Пётр", "Платон", "Потап", "Прохор",
	"Радий", "Радик", "Радомир", "Радослав", "Ринат (Ренат)", "Родион", "Роман", "Ростислав", "Руслан",
	"Савелий", "Святослав", "Севастьян", "Семён", "Сергей", "Сидор", "Спартак", "Станислав", "Степан",
	"Тарас", "Тимофей", "Тихон", "Трофим", "Фёдор", "Федот", "Филипп", "Флор", "Харитон", "Юлиан", "Юрий",
}
