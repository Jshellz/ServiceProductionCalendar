package spc

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type holiday struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
}

var holidays = []holiday{
	{ID: strconv.Itoa(1), Name: "Новый год", Data: "01.01.2024"},
	{ID: strconv.Itoa(2), Name: "Новогодние каникулы", Data: "02.01.2024"},
	{ID: strconv.Itoa(3), Name: "Новогодние каникулы", Data: "03.01.2024"},
	{ID: strconv.Itoa(4), Name: "Новогодние каникулы", Data: "04.01.2024"},
	{ID: strconv.Itoa(5), Name: "Новогодние каникулы", Data: "05.01.2024"},
	{ID: strconv.Itoa(6), Name: "Новогодние каникулы", Data: "06.01.2024"},
	{ID: strconv.Itoa(7), Name: "Рождество Христово", Data: "07.01.2024"},
	{ID: strconv.Itoa(8), Name: "Новогодние каникулы", Data: "08.01.2024"},
	{ID: strconv.Itoa(9), Name: "День защитника Отечества", Data: "23.02.2024"},
	{ID: strconv.Itoa(10), Name: "Международный женский день", Data: "08.03.2024"},
	{ID: strconv.Itoa(11), Name: "Праздник Весны и Труда", Data: "01.05.2024"},
	{ID: strconv.Itoa(12), Name: "День Победы", Data: "09.05.2024"},
	{ID: strconv.Itoa(13), Name: "День России", Data: "12.06.2024"},
	{ID: strconv.Itoa(14), Name: "День народного единства", Data: "04.11.2024"},
}

func GetHoliday(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, holidays)
}

func CreateHoliday(c *gin.Context) {
	var newHoliday holiday
	if err := c.BindJSON(&newHoliday); err != nil {
		return
	}

	holidays = append(holidays, newHoliday)
	c.IndentedJSON(http.StatusCreated, newHoliday)
}

func HolidayById(c *gin.Context) {
	id := c.Param("id")
	holidays, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "holiday not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, holidays)
}

func CheckoutHoliday(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing is query parameter"})
		return
	}

	holidays, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "holiday not found"})
		return
	}

	if holidays.Data == holidays.Data {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "holiday not available"})
		return
	}

	holidays.Data = strconv.Itoa(0)
	c.IndentedJSON(http.StatusOK, holidays)
}

func DeleteHoliday(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing is query parameter"})
		return
	}

	holidays, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "holiday not found"})
		return
	}

	i := holidays
	if i == holidays {
		c.IndentedJSON(http.StatusOK, holidays)
	}
}

func getBookById(id string) (*holiday, error) {
	for i, h := range holidays {
		if h.ID == id {
			return &holidays[i], nil
		}
	}
	return nil, errors.New("holiday not found")
}
