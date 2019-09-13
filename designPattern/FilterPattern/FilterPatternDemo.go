package main

import (
	"strings"
	"fmt"
)

type Person struct {
	name          string
	gender        string
	maritalStatus string
}

func (p *Person) getName() string {
	return p.name
}
func (p *Person) getGender() string {
	return p.gender
}
func (p *Person) getMaritalStatus() string {
	return p.maritalStatus
}

type Criteria interface {
	meetCriteria(persons []Person) []Person
}
type CriteriaMale struct{}

func (cm *CriteriaMale) meetCriteria(persons []Person) []Person {
	// malePersons := make([]Person, 0)
	var malePersons []Person
	for _, v := range persons {

		if strings.ToUpper(v.getGender()) == "MALE" {
		
			malePersons = append(malePersons, v)
		}
	}
	return malePersons
}

type CriteriaFemale struct{}

func (cf *CriteriaFemale) meetCriteria(persons []Person) []Person {
	femalePersons := make([]Person, 0)
	for _, v := range persons {
		p := v
		if strings.ToUpper(p.getGender()) == "FEMALE" {
			femalePersons = append(femalePersons, p)
		}
	}
	return femalePersons
}

type CriteriaSingle struct{}

func (cs *CriteriaSingle) meetCriteria(persons []Person) []Person {
	singlePersons := make([]Person, 0)
	for _, v := range persons {
		if strings.ToUpper(v.getMaritalStatus()) == "SINGLE" {
			singlePersons = append(singlePersons, v)
		}
	}
	return singlePersons
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

//使用需要的过滤组合
func (s *AndCriteria) AndCriteria(criteria Criteria, otherCriteria Criteria) {
	s.criteria = criteria
	s.otherCriteria = otherCriteria
}

func (ac *AndCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaPersons := ac.criteria.meetCriteria(persons)
	return ac.otherCriteria.meetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (oc *OrCriteria) OrCriteria(criteria Criteria, otherCriteria Criteria) {
	oc.criteria = criteria
	oc.otherCriteria = otherCriteria
}
func (oc *OrCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaItems := oc.criteria.meetCriteria(persons)
	otherCriteriaItems := oc.otherCriteria.meetCriteria(persons)
	for _, v1 := range otherCriteriaItems {
		for _, v2 := range firstCriteriaItems {
			if v1 == v2 {
				firstCriteriaItems = append(firstCriteriaItems, v1)
			}
		}
	}
	return firstCriteriaItems
}

func main() {
	persons := make([]Person, 0)
	persons = append(persons, Person{"Robert", "Male", "Single"})
	persons = append(persons, Person{"John", "Male", "Married"})
	persons = append(persons, Person{"Laura", "Female", "Married"})
	persons = append(persons, Person{"Diana", "Female", "Single"})
	persons = append(persons, Person{"Mike", "Male", "Single"})
	persons = append(persons, Person{"Bobby", "Male", "Single"})

	male := new(CriteriaMale)
	fmt.Println(male.meetCriteria(persons))
	female := new(CriteriaFemale)
	fmt.Println(female.meetCriteria(persons))
	single := new(CriteriaSingle)
	fmt.Println(single.meetCriteria(persons))

	singleMale := new(AndCriteria)
	singleMale.AndCriteria(single, male)
	fmt.Println(singleMale.meetCriteria(persons))

	singleOrFemale := new(OrCriteria)
	singleOrFemale.OrCriteria(single, male)
	fmt.Println(singleOrFemale.meetCriteria(persons))
}
