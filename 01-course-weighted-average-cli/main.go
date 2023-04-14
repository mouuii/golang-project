package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Course struct {
	name   string
	weight float64
	score  float64
}

func main() {
	var courses []Course
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入课程名、学分和成绩，用空格分隔（输入'q'退出）: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("输入格式错误，请重新输入。")
			continue
		}

		weight, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("学分输入错误，请重新输入。")
			continue
		}

		score, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("成绩输入错误，请重新输入。")
			continue
		}

		course := Course{
			name:   parts[0],
			weight: weight,
			score:  score,
		}

		courses = append(courses, course)
	}

	if len(courses) == 0 {
		fmt.Println("没有输入课程信息。")
		return
	}

	var totalWeight, totalScore float64
	for _, course := range courses {
		totalWeight += course.weight
		totalScore += course.weight * course.score
	}

	gpa := totalScore / totalWeight
	fmt.Printf("加权平均分：%.2f\n", gpa)
}
