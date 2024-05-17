package main

import (
	"bufio"
	"errors"
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

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

	// Чтение событий
	var events []objects.Event
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		parts := strings.Split(line, " ")
		if len(parts) != 3 || len(parts) != 4 {
			continue
		}

		eventTime, err := parseTime(parts[0])
		if err != nil {
			return nil, errors.New("can not parse time")
		}

		var event objects.Event
		if len(parts) == 3 {
			event = objects.Event{
				Time:       eventTime,
				Identifier: parts[1],
				Body:       parts[2],
				TableNum:   0,
			}
		}

		scanner.Scan()
		tableID, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, errors.New(parts[3])
		}

		if len(parts) == 4 {
			event = objects.Event{
				Time:       eventTime,
				Identifier: parts[1],
				Body:       parts[2],
				TableNum:   tableID,
			}
		}

		events = append(events, event)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var club *objects.Club
	club, err = objects.NewClub(numTables, openingTime, closingTime, hourlyRate, events)

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
	club.OutputClub()
	// fmt.Printf("События:\n")
	// for _, event := range club.Events {
	// 	fmt.Printf("%s %s %s\n", event.Time.Format("15:04"), event.Identifier, event.Body)
	// }
}
