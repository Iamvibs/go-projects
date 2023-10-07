package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func problemPuller(filename string) ([]problem, error) {
	if fObj, err := os.Open(filename); err == nil {
		csvR := csv.NewReader(fObj)
		if clines, err := csvR.ReadAll(); err == nil {
			return parseProblem(clines), nil
		} else {
			return nil, fmt.Errorf("error in reading data of csv")
		}
	} else {
		return nil, fmt.Errorf("error in opening the file")
	}
}

func parseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

func main() {
	fName := flag.String("f", "quiz.csv", "path of csv file")

	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()

	problems, err := problemPuller(*fName)
	if err != nil {
		exit(fmt.Sprintf("something went wrong: %s", err.Error()))
	}

	correctAns := 0
	presentedProbs := 0

	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	done := make(chan bool)

problemLoop:

	for i, p := range problems {
		presentedProbs = i + 1
		fmt.Printf("Problem %d: %s=", i+1, p.q)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansC <- answer
			done <- true
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}

		<-done
	}

	fmt.Printf("Your result is %d out of %d\n", correctAns, presentedProbs)
	fmt.Printf("Press enter to exit")
	<-ansC
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
