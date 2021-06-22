package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

/*
Gopher excercice workout problem. git for gourcing the project. @ https://github.com/gophercises/quiz
1. Open & print data from csv file
2. turn each entry into a question
3. Evaluate & store answers
*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	//filePath string = "problems.csv"

	total int
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "A csv file name with the format 'question,answer'.")
	timerDuration := flag.Int("timer", 3, "The duration for each question.")
	flag.Parse()

	//1 print file
	//f, err := ioutil.ReadFile(string(filePath))
	f, err := os.Open(*csvFilename)
	check(err)
	r, err := csv.NewReader(f).ReadAll()
	check(err)
	//2. turn each entry into a question
	timer := time.NewTimer(time.Duration(*timerDuration) * time.Second)
	for _, ques := range r {
		timer.Reset(time.Duration(*timerDuration) * time.Second)
		fmt.Printf("What is  %v = ?", ques[0])
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			{
				fmt.Printf("\nTime ran out. you Scored %v/%v.\n", total, len(r))
				return
			}
		case answr := <-answerCh:
			{
				if answr == ques[1] {
					total = total + 1
				}
			}
		}

		//corrAns, _ := strconv.Atoi(ques[1])
		// 3. Evaluate & store answers

	}
	fmt.Printf("Your Score is %v/%v\n", total, len(r))

}
