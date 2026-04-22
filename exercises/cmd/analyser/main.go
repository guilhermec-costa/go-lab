package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"strconv"
	"strings"
)

func parsedIntArgOrDefault(args []string, argname string, _default int) int {
	for _, val := range args {
		parsed, found := strings.CutPrefix(val, argname+"=")
		if !found {
			return _default
		}

		if len(parsed) == 0 {
			log.Println("Value after '=' in arg", argname, "missing")
			return _default
		}

		if i, err := strconv.Atoi(parsed); err != nil {
			log.Println("Failed to parse int argument. Returning default '", _default, "'")
			return _default
		} else {
			return i
		}
	}
	return _default
}

func askStudentName(studentIdx int) string {
	fmt.Print(strings.Repeat("=", 10))
	fmt.Printf(" Student %v ", studentIdx + 1);
	fmt.Print(strings.Repeat("=", 10), "\n");
	studentName := ""

OuterLoop:
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Student's fullname > ")

		if scanner.Scan() {
			studentName = scanner.Text()
		}

		subnames := strings.Split(studentName, " ")

		if len(subnames) < 2 {
			log.Printf("Student must have at least 2 subnames. Try again");
			continue;
		}

		for _, name := range subnames {
			if len(name) <= 2 {
				log.Printf("Student subname \"%v\" must have at least 3 characters. Try again", name)
				continue OuterLoop;
			}
		}

		break
	}

	caser := cases.Title(language.Portuguese, cases.NoLower)
	return caser.String(studentName)
}

func askForGrade(gradeIdx int) Grade {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Type grade ", gradeIdx+1, " >> ")

		if scanner.Scan() {
			val, err := strconv.ParseFloat(scanner.Text(), 32 /* to 32 bits */)
			if err != nil {
				log.Print("Failed to parse grade \"", scanner.Text(), "\". Try again")
				continue
			}

			return Grade(val)
		}
	}
}

type Args struct {
	StudentsCount   int
	ScorePerStudent int
}

type Grade float32
type GradeResume map[string][]Grade

func startProcessing(parsedArgs Args) {
	gradesResume := make(GradeResume)

	for s := range parsedArgs.StudentsCount {
		studentName := askStudentName(s)
		gradesResume[studentName] = make([]Grade, parsedArgs.ScorePerStudent);
		log.Printf("Student \"%v\" registered in the system.", studentName);

		studentGrades := make([]Grade, parsedArgs.ScorePerStudent)
		for i := range parsedArgs.ScorePerStudent {
			studentGrades[i] = askForGrade(i)
		}
		gradesResume[studentName] = studentGrades
	}

	fmt.Println("Final state: ", gradesResume)
}

func main() {
	log.SetPrefix("analyser: ")
	log.SetFlags(0)

	cliArgs := os.Args[1:]

	parsedArgs := Args{
		StudentsCount:   parsedIntArgOrDefault(cliArgs, "-students", 5),
		ScorePerStudent: parsedIntArgOrDefault(cliArgs, "-score", 3),
	}

	startProcessing(parsedArgs)
}
