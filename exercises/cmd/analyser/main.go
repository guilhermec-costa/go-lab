package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type ArgParser[T any] func(string) (T, error)

func showHeader(title string, eqSideLen int) {
	fmt.Print(strings.Repeat("=", eqSideLen))
	fmt.Print(title)
	fmt.Print(strings.Repeat("=", eqSideLen), "\n")
}

func parseArgOrDefault[T any](
	args []string,
	argname string,
	_default T,
	parseFn ArgParser[T],
) T {
	for _, val := range args {
		parsed, found := strings.CutPrefix(val, argname+"=")
		if !found {
			continue
		}

		if len(parsed) == 0 {
			log.Println("Value after '=' in arg", argname, "missing")
			return _default
		}

		if val, err := parseFn(parsed); err != nil {
			log.Println("Failed to parse int argument. Returning default '", _default, "'")
			return _default
		} else {
			return val
		}
	}

	return _default
}

func askStudentName(studentIdx int) string {
	showHeader(fmt.Sprintf(" Student %v ", studentIdx+1), 10)

	studentName := ""

	scanner := bufio.NewScanner(os.Stdin)
OuterLoop:
	for {
		fmt.Print("Student's fullname > ")

		if scanner.Scan() {
			studentName = scanner.Text()
		}

		subnames := strings.Split(studentName, " ")

		if len(subnames) < 2 {
			log.Printf("Student must have at least 2 subnames. Try again")
			continue
		}

		for _, name := range subnames {
			if len(name) <= 2 {
				log.Printf("Student subname \"%v\" must have at least 3 characters. Try again", name)
				continue OuterLoop
			}
		}

		break
	}

	caser := cases.Title(language.Portuguese, cases.NoLower)
	return caser.String(studentName)
}

func askForGrade(gradeIdx int, maxGrade float32) Grade {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Type grade ", gradeIdx+1, " >> ")

		if scanner.Scan() {
			val, err := strconv.ParseFloat(scanner.Text(), 32 /* to 32 bits */)
			if err != nil {
				log.Print("Failed to parse grade \"", scanner.Text(), "\". Try again")
				continue
			}

			if val < 0.0 {
				log.Printf("Grade can not be less than 0.0. Try again")
				continue
			}

			if float32(val) > maxGrade {
				log.Printf("Grade can not be greater than %v. Try again", maxGrade)
				continue
			}

			return Grade(val)
		}
	}
}

type Args struct {
	StudentsCount   int
	ScorePerStudent int
	MaxScore        float32
}

type Grade float32
type GradeResume map[string][]Grade

func makeGradesResume(args Args) GradeResume {
	gradesResume := make(GradeResume)

	for s := range args.StudentsCount {
		studentName := askStudentName(s)
		gradesResume[studentName] = make([]Grade, args.ScorePerStudent)
		log.Printf("Student \"%v\" registered in the system.", studentName)

		for i := range args.ScorePerStudent {
			gradesResume[studentName][i] = askForGrade(i, args.MaxScore)
		}
		log.Printf("Student \"%v\" grade registered in the system.", studentName)
	}

	return gradesResume
}

func calcAvg(grades []Grade) float64 {
	var sum float32 = 0

	for _, n := range grades {
		sum += float32(n)
	}

	var prec float64 = 10
	return math.Trunc(
		float64(sum/float32(len(grades)))*prec) / prec
}

type StudentResume struct {
	name   string
	avg    Grade
	status string
}

func getStudentStatus(grade float64) string {
	switch {
	case grade >= 7.0:
		return "Aprovado"

	case grade >= 5:
		return "Recuperação"

	default:
		return "Reprovado"
	}
}

func processResumeByStudent(resume GradeResume) []StudentResume {
	var studentsResumes []StudentResume

	for k, v := range resume {
		avg := calcAvg(v)
		resume := StudentResume{
			name:   k,
			avg:    Grade(avg),
			status: getStudentStatus(avg),
		}
		studentsResumes = append(studentsResumes, resume)
	}

	return studentsResumes
}

func presentStudentsResumes(resumes []StudentResume) {
	showHeader("Average presentation", 10)
	for _, resume := range resumes {
		fmt.Printf("Student: %v | Average Grade: %v | Status: %v |\n", resume.name, resume.avg, resume.status)
	}
}

func main() {
	log.SetPrefix("analyser: ")
	log.SetFlags(0)

	cliArgs := os.Args[1:]

	parsedArgs := Args{
		StudentsCount: parseArgOrDefault(
			cliArgs,
			"-students",
			5,
			strconv.Atoi,
		),
		ScorePerStudent: parseArgOrDefault(
			cliArgs,
			"-scores-per-student",
			3,
			strconv.Atoi,
		),
		MaxScore: parseArgOrDefault(
			cliArgs,
			"-max-score",
			float32(10.0),
			func(s string) (float32, error) {
				val, err := strconv.ParseFloat(s, 32)
				return float32(val), err
			},
		),
	}

	gradesResume := makeGradesResume(parsedArgs)
	avgByStudent := processResumeByStudent(gradesResume)
	presentStudentsResumes(avgByStudent)
}
