package example

import (
	"testing"
)

func TestStoreMonster(t *testing.T)  {
	m := &Monster{"乱世狂刀",30,"悲龙斩"}
	err := m.Store()
	if err != nil {
		t.Fatalf("保存数据失败, err:%v", err)
	}

	t.Logf("保存数据成功")

}

func TestRestoreMonster(t *testing.T) {

	m := &Monster{}
	err := m.ReStore()
	if err != nil {
		t.Fatalf("读取数据失败, err:%v", err)
	}
	if m.Name != "乱世狂刀" {
		t.Fatalf("读取姓名失败：名字不正确；读取的年龄为%s",m.Name)
	}
	if m.Age != 30 {
		t.Fatalf("读取年龄失败：年龄不正确；读取的年龄为%d",m.Age)
	}
	if m.Skill != "悲龙斩" {
		t.Fatalf("读取技能失败：技能不正确；读取的年龄为%s",m.Skill)
	}
	t.Logf("读取数据成功")
}
