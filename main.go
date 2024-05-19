package main

import (
	"bufio"
	"errors"
	"example.com/computer-club/incomingEvents"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/computer-club/objects"
)

func parseTime(input string) (time.Time, error) {
	layout := "15:04"
	t, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ReadInputFile(filename string) (*objects.Club, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("can not open file")
	}

	scanner := bufio.NewScanner(file)

	// Чтение количества столов
	scanner.Scan()
	line := scanner.Text()
	numTables, err := strconv.Atoi(line)
	if err != nil {
		return nil, errors.New(line)
	}

	// Чтение времени начала и окончания работы
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	openingTime, err := parseTime(parts[0])
	if err != nil {
		return nil, errors.New(scanner.Text())
	}
	fmt.Printf("%s\n", openingTime.Format("15:04"))

	closingTime, err := parseTime(parts[1])
	if err != nil {
		return nil, errors.New(scanner.Text())
	}

	// Чтение стоимости часа
	scanner.Scan()
	lineRate := scanner.Text()
	hourlyRate, err := strconv.Atoi(lineRate)
	if err != nil {
		return nil, errors.New(lineRate)
	}

	var club *objects.Club
	club, err = objects.NewClub(numTables, openingTime, closingTime, hourlyRate)

	// Чтение событий
	var Events []objects.Event
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		eventParts := strings.Split(line, " ")
		//if len(eventParts) != 3 || len(eventParts) != 4 {
		//	continue
		//}

		eventTime, err := parseTime(eventParts[0])
		if err != nil {
			return nil, errors.New("can not parse time")
		}

		var event *objects.Event
		if len(eventParts) == 3 {
			event, err = objects.NewEvent(eventTime, eventParts[1], eventParts[2], 0)
		}

		//scanner.Scan()

		if len(eventParts) == 4 {
			tableID, err := strconv.Atoi(eventParts[3])
			if err != nil {
				return nil, errors.New(eventParts[3])
			}
			event, err = objects.NewEvent(eventTime, eventParts[1], eventParts[2], tableID)
			if err != nil {
				return nil, errors.New(line)
			}
		}

		//регулирование действий игроков
		switch event.Identifier() {
		case "1":
			incomingEvents.Id1(event, club)
		}

		Events = append(Events, *event)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("can not make club")
	}

	err = file.Close()
	if err != nil {
		return nil, errors.New("can not close file")
	}

	return club, nil

}

func main() {
	filename := os.Args[1]
	club, err := ReadInputFile(filename)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	// Дальнейшая обработка данных из файла
	club.PrintClub()
	// fmt.Printf("События:\n")
	// for _, event := range club.Events {
	// 	fmt.Printf("%s %s %s\n", event.Time.Format("15:04"), event.Identifier, event.Body)
	// }
}
