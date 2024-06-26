package objects

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Table struct {
	number       int
	revenue      int
	startTime    time.Time
	endTime      time.Time
	duration     int
	occupiedTime time.Duration
}

func NewTable(number int, startTime, endTime time.Time) (*Table, error) {
	return &Table{
		number:       number,
		startTime:    startTime,
		duration:     int(endTime.Sub(startTime).Minutes()),
		occupiedTime: 0,
	}, nil
}

func (table *Table) CalculateDurationInMinutes() {
	table.duration = int(table.endTime.Sub(table.startTime).Minutes())
}

func (table *Table) Revenue() int {
	return table.revenue
}

func (table *Table) UpdateOccupiedTime() {
	table.occupiedTime += table.endTime.Sub(table.startTime)
}

func (table *Table) UpdateRevenue(rate int) {
	hours := math.Ceil(float64(table.duration) / 60)
	price := int(hours) * rate
	table.revenue += price
}

func (table *Table) UpdateDuration() {
	table.duration = int(table.endTime.Sub(table.startTime).Minutes())
}

func (table *Table) SetStartTime(startTime time.Time) {
	table.startTime = startTime
}

func (table *Table) SetDuration(duration int) {
	table.duration = duration
}

func (table *Table) SetEndTime(endTime time.Time) {
	table.endTime = endTime
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
	currentGamers   map[string]int
	currentTables   map[int]string
	waitList        []string
	tablesRevenue   map[int]Table
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
		currentGamers:   make(map[string]int),
		currentTables:   make(map[int]string),
		waitList:        []string{},
		tablesRevenue:   make(map[int]Table),
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

func (club *Club) HourlyRate() int {
	return club.hourlyRate
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

func (club *Club) AddGamer(gamer string, table int) {
	club.currentGamers[gamer] = table
}

func (club *Club) RemoveGamer(gamer string) {
	_, exists := club.currentGamers[gamer]
	if exists {
		delete(club.currentGamers, gamer)
	}
}

func (club *Club) GetGamerTable(gamer string) int {
	table, exists := club.currentGamers[gamer]
	if exists {
		return table
	}
	return 0
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

func (club *Club) GetWaitListLength() int {
	return len(club.waitList)
}

func (club *Club) GetClientFromWaitList(index int) string {
	if index >= 0 && index < len(club.waitList) {
		return club.waitList[index]
	}
	return ""
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
	_, exists := club.currentTables[numTable]
	if exists {
		delete(club.currentTables, numTable)
	}
}

func (club *Club) GetCurrTableCount() int {
	return len(club.currentTables)
}

func (club *Club) PrintClub() {
	fmt.Printf("Количество столов: %d\n", club.tables)
	fmt.Printf("Время начала работы: %s\n", club.openingTime.Format("15:04"))
	fmt.Printf("Время окончания работы: %s\n", club.closingTime.Format("15:04"))
	fmt.Printf("Стоимость часа: %d\n", club.hourlyRate)
}

func (club *Club) ClubCloses() *[]string {
	keys := make([]string, 0, len(club.currentVisitors))
	for key := range club.currentVisitors {
		if club.currentVisitors[key] {
			keys = append(keys, key)
			club.currentVisitors[key] = false
		}
	}

	sort.Strings(keys)
	return &keys
}

func (club *Club) GetTablesRevenue() map[int]Table {
	return club.tablesRevenue
}

func (club *Club) SetTablesRevenue(m map[int]Table) {
	club.tablesRevenue = m
}

type Gamer struct {
	Name  string
	Table int
}
