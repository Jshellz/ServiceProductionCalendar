package spc

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Holiday struct {
	*gorm.Model
	ID   string    `gorm:"primary_key" json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"data"`
}

const (
	year       = 2024
	dayJONe    = 1
	dayJTwo    = 2
	dayJThree  = 3
	dayJFour   = 4
	dayJFive   = 5
	dayJSix    = 6
	dayJR      = 7
	dayJEight  = 8
	dayF       = 23
	dayM       = 8
	dayMayOne  = 1
	dayMayNine = 9
	dayJun     = 12
	dayNov     = 4
	hour       = 00
	min        = 00
	sec        = 00
	nsec       = 00
)

var Holidays = []Holiday{
	{ID: strconv.Itoa(1), Name: "Новый год", Date: time.Date(year, time.January, dayJONe, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(2), Name: "Новогодние каникулы", Date: time.Date(year, time.January, dayJTwo, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(3), Name: "Новогодние каникулы", Date: time.Date(year, time.January, dayJThree, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(4), Name: "Новогодние каникулы", Date: time.Date(year, time.January, dayJFour, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(5), Name: "Новогодние каникулы", Date: time.Date(year, time.January, dayJFive, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(6), Name: "Новогодние каникулы", Date: time.Date(year, time.January, dayJSix, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(7), Name: "Рождество Христово", Date: time.Date(year, time.January, dayJR, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(8), Name: "Новогодние каникулы", Date: time.Date(year, time.January, dayJEight, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(9), Name: "День защитника Отечества", Date: time.Date(year, time.February, dayF, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(10), Name: "Международный женский день", Date: time.Date(year, time.March, dayM, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(11), Name: "Праздник Весны и Труда", Date: time.Date(year, time.May, dayMayOne, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(12), Name: "День Победы", Date: time.Date(year, time.May, dayMayNine, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(13), Name: "День России", Date: time.Date(year, time.June, dayJun, hour, min, sec, nsec, time.Local)},
	{ID: strconv.Itoa(14), Name: "День народного единства", Date: time.Date(year, time.November, dayNov, hour, min, sec, nsec, time.Local)},
}

type CalendarService struct {
	Holiday
}

// GetHoliday получение всех праздников
func GetHoliday(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Holidays)
	h := connTmpl
	c.IndentedJSON(http.StatusOK, h)

}

type tmpl struct {
	Title    string
	holidays *Holiday
}

func connTmpl(w http.ResponseWriter, _ *http.Request) (h *Holiday, templates tmpl) {

	templates = tmpl{Title: "SPC", holidays: h}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return nil, templates
	}

	log.Fatal(tmpl.Execute(w, Holidays), templates)
	return h, templates
}

// CreateHoliday создание праздника
func CreateHoliday(c *gin.Context) {
	var newHoliday Holiday
	if err := c.BindJSON(&newHoliday); err != nil {
		return
	}

	Holidays = append(Holidays, newHoliday)
	c.IndentedJSON(http.StatusCreated, newHoliday)
}

// HolidayById получение праздника по id
func HolidayById(c *gin.Context) {
	id := c.Param("id")
	holidays, err := getHolidayById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "holiday not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, holidays)
}

// CheckoutHoliday изменение праздника
func CheckoutHoliday(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing is query parameter"})
		return
	}

	holidays, err := getHolidayById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "holiday not found"})
		return
	}

	if holidays.Date == holidays.Date {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "holiday not available"})
		return
	}

	var times time.Time

	holidays.Date = times
	c.IndentedJSON(http.StatusOK, holidays)
}

// DeleteHoliday удаление праздника
func DeleteHoliday(c *gin.Context) {
	id := c.Param("id")

	if err := deleteHoliday(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "holiday not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "holiday deleted successfully"})
}

func deleteHoliday(id string) error {
	deleteHolidays := delExec
	if deleteHolidays != nil {
		return deleteHoliday(id)
	}
	return errors.New("failed to delete")
}

// запрос на удаление из бд
func delExec(db gorm.DB) {
	db.Exec("DELETE FROM holidays WHERE id = $1", deleteHoliday("id"))
}

func getHolidayById(id string) (*Holiday, error) {
	for i, h := range Holidays {
		if h.ID == id {
			return &Holidays[i], nil
		}
	}
	return nil, errors.New("holiday not found")
}
