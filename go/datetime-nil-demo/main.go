package main

import (
	"flag"
	"fmt"
	"time"
)

var startDate *time.Time

func main() {

	var err error
	fmt.Println("Type a date")
	paramStartDate := flag.String("startdate", "", "--startdate 2021-10-20T10:30:00.000Z")
	flag.Parse()
	if paramStartDate != nil && *paramStartDate != "" {
		*startDate, err = time.Parse(time.RFC3339Nano, *paramStartDate)
		if err != nil {
			fmt.Println("Error parsing date")
		}
	}
	fmt.Printf("Date passed as param: %s\n", *paramStartDate)

	testDate(startDate)
}

func testDate(replayKafkaFrom *time.Time) {
	if replayKafkaFrom == nil {
		fmt.Println("Nil param received.")
	}
	replayFromTimestamp := replayKafkaFrom != nil && !replayKafkaFrom.IsZero()
	fmt.Printf("Replay from timestamp: %t\n", replayFromTimestamp)

	if replayKafkaFrom != nil {
		fmt.Printf("Date in time.Time format: %s\n", replayKafkaFrom)
		fmt.Printf("Date is zero: %t", replayKafkaFrom.IsZero())
	}

	t := NewKafkaConsumer("test", replayKafkaFrom)
	fmt.Println(t.consumerGroup)

}

type KafkaConsumer struct {
	consumerGroup string
	consumerDate  time.Time
}

func NewKafkaConsumer(cg string, cd *time.Time) KafkaConsumer {
	newkafka := KafkaConsumer{consumerGroup: cg}
	if cd != nil {
		newkafka.consumerDate = *cd
	}
	return newkafka
}
