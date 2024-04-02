package main

import (
	"day21/Hero"
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	var heroes Hero.HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero.Hero{Age: rand.Intn(100), Name: fmt.Sprintf("英雄～%d", rand.Intn(100))}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println(heroes)
}
