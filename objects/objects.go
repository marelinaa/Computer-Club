package objects

import (
	"fmt"
	"time"
)

type Table struct {
	Number   int
	Revenue  int
	Duration int
}

type Event struct {
	time       time.Time
	identifier string
	body       string
	tableNum   int
}

func NewEvent(time time.Time, identifier, body string, tableNum int) (*Event, error) {
	return &Event{
		time:       time,
		identifier: identifier,
		body:       body,
		tableNum:   tableNum,
	}, nil
}

func (event *Event) EventToString() string {
	return fmt.Sprintf("%v %s %s %d", event.time.Format("15:04"), event.identifier, event.body, event.tableNum)
}

func (event *Event) Time() time.Time {
	return event.time
}

func (event *Event) Identifier() string {
	return event.identifier
}

func (event *Event) Body() string {
	return event.body
}

func (event *Event) TableNum() int {
	return event.tableNum
}

func (event *Event) PrintEvent() {
	fmt.Printf("%s %s %s\n", event.time.Format("15:04"), event.identifier, event.body)
}

type Club struct {
	tables          int
	openingTime     time.Time
	closingTime     time.Time
	hourlyRate      int
	currentVisitors map[string]bool
	currentTables   map[int]string
	waitList        []string
	//events      []Event
}

func NewClub(numTables int, openingTime, closingTime time.Time, hourlyRate int) (*Club, error) {
	// if firstName == "" || lastName == "" || birthDate == "" {
	// 	return nil, errors.New("entered data can not be empty")
	// }

	return &Club{
		tables:          numTables,
		openingTime:     openingTime,
		closingTime:     closingTime,
		hourlyRate:      hourlyRate,
		currentVisitors: make(map[string]bool),
		currentTables:   make(map[int]string),
		waitList:        []string{},
		//events:      events,
	}, nil
}

func (club *Club) Tables() int {
	return club.tables
}

func (club *Club) OpeningTime() time.Time {
	return club.openingTime
}

func (club *Club) ClosingTime() time.Time {
	return club.closingTime
}

func (club *Club) IsVisitorInClub(visitor string) bool {
	exists := club.currentVisitors[visitor]
	return exists
}

func (club *Club) AddVisitor(visitor string) {
	club.currentVisitors[visitor] = true
}

func (club *Club) RemoveVisitor(visitor string) {
	club.currentVisitors[visitor] = false
}

func (club *Club) AddToWaitList(visitor string) {
	club.waitList = append(club.waitList, visitor)
}

func (club *Club) RemoveFromWaitList(visitor string) {
	for i, v := range club.waitList {
		if v == visitor {
			club.waitList = append(club.waitList[:i], club.waitList[i+1:]...)
			break
		}
	}
}

func (club *Club) WhoUsesTable(numTable int) string {
	gamer, exists := club.currentTables[numTable]
	if exists {
		return gamer
	}

	return ""
}

func (club *Club) AddTable(numTable int, gamer string) {
	club.currentTables[numTable] = gamer
}

func (club *Club) RemoveTable(numTable int) {
	delete(club.currentTables, numTable)
}

func (club *Club) GetCurrTableCount() int {
	return len(club.currentTables)
}

func (club *Club) GetWaitListLength() int {
	return len(club.waitList)
}

func (club *Club) PrintClub() {
	fmt.Printf("Количество столов: %d\n", club.tables)
	fmt.Printf("Время начала работы: %s\n", club.openingTime.Format("15:04"))
	fmt.Printf("Время окончания работы: %s\n", club.closingTime.Format("15:04"))
	fmt.Printf("Стоимость часа: %d\n", club.hourlyRate)
}

type Gamer struct {
	Name  string
	Table int
}
