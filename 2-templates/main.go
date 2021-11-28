package main

import (
	"log"
	"os"
	"text/template"
)

type meal struct {
	Name  string
	Price float64
}

type mealType struct {
	Name  string
	Meals []meal
}

type menu struct {
	Restaurant string
	MealTypes  []mealType
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index_template.gohtml"))
}

func main() {
	var m menu = menu{
		Restaurant: "Isabelle's Restaurant",
		MealTypes: []mealType{
			{
				Name: "Breakfast",
				Meals: []meal{
					{
						Name:  "Tocino w/ Rice",
						Price: 50,
					},
				},
			},
			{
				Name: "Lunch",
				Meals: []meal{
					{
						Name:  "Lucban Longganisa w/ Rice",
						Price: 80,
					},
					{
						Name:  "Fried Chicken w/ Rice",
						Price: 120,
					},
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
