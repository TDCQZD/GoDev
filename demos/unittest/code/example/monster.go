package example

import (
	"io/ioutil"
	"encoding/json"
)

type Monster struct{
	Name string
	Age int
	Skill string
}

func (m *Monster)Store() (err error) {
	data, err := json.Marshal(m)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("d:/monster.dat", data, 0755)
	return

	
}

func (m *Monster)ReStore() (err error) {
	data, err := ioutil.ReadFile("d:/monster.dat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, m)
	return

}


