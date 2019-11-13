package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const studentID = "0000000"

type webData struct {
	Title         string
	StudentID     string
	ActiveHome    string
	ActiveCalc    string
	CalcResult    string
	CalcProcess   string
	CalcRemainder string
	Time          string
	/*
		Add whatever you want
	*/
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tmpl/layout.html", "tmpl/index.html")
	wd := webData{
		Title:      "Home",
		StudentID:  studentID,
		ActiveHome: "active",
	}
	tmpl.Execute(w, &wd)
}
func calcHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tmpl/layout.html", "tmpl/calc.html")
	wd := webData{
		Title:      "Calculator",
		StudentID:  studentID,
		ActiveCalc: "active",
	}
	tmpl.Execute(w, &wd)
}
func resultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	operandOne, err := strconv.Atoi(strings.Join(r.Form["operand_1"], ""))
	operandTwo, err := strconv.Atoi(strings.Join(r.Form["operand_2"], ""))
	operator := strings.Join(r.Form["operator"], "")

	var calcResult, calcProcess, calcRemainder string

	if err != nil {
		calcResult = "Invalid operand!!"
	} else {
		switch operator {
		case "+":
			calcResult = strconv.Itoa(operandOne + operandTwo)
		case "-":
			calcResult = strconv.Itoa(operandOne - operandTwo)
		case "*":
			calcResult = strconv.Itoa(operandOne * operandTwo)
		case "/":
			calcResult = strconv.Itoa(operandOne / operandTwo)
			calcRemainder = strconv.Itoa(operandOne % operandTwo)
		}
		calcProcess = strconv.Itoa(operandOne) + " " + operator + " " + strconv.Itoa(operandTwo) + " = " + calcResult
	}

	tmpl, _ := template.ParseFiles("tmpl/layout.html", "tmpl/result.html")
	wd := webData{
		Title:         "Result",
		StudentID:     studentID,
		ActiveCalc:    "active",
		CalcResult:    calcResult,
		CalcProcess:   calcProcess,
		CalcRemainder: calcRemainder,
		Time:          time.Now().Format("2006-01-02 15:04:05"),
	}

	tmpl.Execute(w, &wd)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calc", calcHandler)
	http.HandleFunc("/result", resultHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
