package utils

// 因为是小写，所以是私有数据，只能在本包使用
type student struct {
	Name  string
	Age   int
	score float64
}

// FactoryStudent 但是可以使用工厂模式来解决
func FactoryStudent(name string, age int, score float64) *student {
	return &student{
		Name:  name,
		Age:   age,
		score: score,
	}
}

// GetScore 解决私有化字段访问问题
func (s *student) GetScore() float64 {
	return s.score
}

// 继承，封装、多态

// 抽象
