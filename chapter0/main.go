package main

import "fmt"

func forCalc() int {
	count := 0
	var result int
	for count < 10 {
		result = count
		count++
	}
	return result
}

func ifCalc() string {
	ifNum := forCalc()
	if 5 > ifNum {
		return fmt.Sprintf("%d は5未満の数字です。", ifNum)
	} else if 8 > ifNum {
		return fmt.Sprintf("%d は 5 以上 8 未満の数字です。", ifNum)
	} else {
		return fmt.Sprintf("%d は8以上10以下の数字です。", ifNum)
	}
}

func switchCalc() string {
	switchNum := forCalc()
	switch switchNum {
	case 1:
		return fmt.Sprintf("%d です。", switchNum)
	case 2:
		return fmt.Sprintf("%d です。", switchNum)
	case 3:
		return fmt.Sprintf("%d です。", switchNum)
	case 4:
		return fmt.Sprintf("%d です。", switchNum)
	case 5:
		return fmt.Sprintf("%d です。", switchNum)
	case 6:
		return fmt.Sprintf("%d です。", switchNum)
	case 7:
		return fmt.Sprintf("%d です。", switchNum)
	case 8:
		return fmt.Sprintf("%d です。", switchNum)
	case 9:
		return fmt.Sprintf("%d です。", switchNum)
	case 10:
		return fmt.Sprintf("%d です。", switchNum)
	default:
		return fmt.Sprint("% です", switchNum)
	}

}

type Person struct {
	height int
	weight int
	name   string
}

func (p Person) addHeight() {
	p.height += 10
}

func (p *Person) addWeight() {
	p.weight += 10
}

func main() {
	sum := forCalc()
	fmt.Printf("%d\n", sum)

	ifSum := ifCalc()
	fmt.Println(ifSum)

	switchSum := switchCalc()
	fmt.Println(switchSum)

	q := Person{height: 160, weight: 50}
	fmt.Println(q)

	q.addHeight()
	q.addWeight()

	fmt.Println(q)
}
