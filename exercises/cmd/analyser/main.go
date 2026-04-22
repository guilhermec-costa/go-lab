package main

import (
	"bufio"
	"fmt"
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
	fmt.Println(strings.Repeat("=", 30));
	studentName := ""
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Type student", studentIdx+1, "'s name > ")

		if scanner.Scan() {
			studentName = scanner.Text()
		}

		if len(studentName) == 0 {
			log.Println("Username must no be empty. Try again")
			continue
		}
		break
	}

	return studentName
}

func askForGrade(gradeIdx int) float32 {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Type grade ", gradeIdx + 1, " >> ");

		if scanner.Scan() {
			val, err := strconv.ParseFloat(scanner.Text(), 32)
			if err != nil {
				log.Print("Failed to parse grade \"", scanner.Text(), "\". Try again")
				continue
			}

			return float32(val)
		}
	}
}

type Args struct {
	StudentsCount   int
	ScorePerStudent int
}

func startProcessing(parsedArgs Args) {
	gradesResume := make(map[string][]float32)
	for s := range parsedArgs.StudentsCount {
		studentName := askStudentName(s)

		studentGrades := make([]float32, parsedArgs.ScorePerStudent)
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

	startProcessing(parsedArgs);
}
