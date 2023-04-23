package main

//LVX
import "fmt"

type Person struct {
	Name  string
	Level int
	Exp   int
	Hp    int
	ATK   int
}
type hit interface {
	Attack(Enemy *Person)
}

func (p *Person) Attack(Enemy *Person) {
	Enemy.Hp -= p.ATK
}

func main() {

	Player1 := Person{"Dingzhen", 114514, 50, 5000, 500}
	Player2 := Person{"Alien", 25000, 50, 9000, 999}
	fmt.Printf("Player1: %s, Level: %d, Exp: %d, HP: %d, ATK: %d\n", Player1.Name, Player1.Level, Player1.Exp, Player1.Hp, Player1.ATK)
	fmt.Printf("Player2: %s, Level: %d, Exp: %d, HP: %d, ATK: %d\n", Player2.Name, Player2.Level, Player2.Exp, Player2.Hp, Player2.ATK)
	for {
		println("输入1或2使人物互相攻击")
		var n int
		fmt.Scanln(&n)
		if n == 1 {
			Player1.Attack(&Player2)
			if Player2.Hp <= 0 {
				println("Dingzhen win win win!")
				break
			}
			println("Dingzhen HP: ", Player1.Hp)
			println("Alien HP: ", Player2.Hp)
		} else if n == 2 {
			Player2.Attack(&Player1)
			if Player1.Hp <= 0 {
				println("Alien win win win!")
				break
			}
			println("Dingzhen HP: ", Player1.Hp)
			println("Alien HP: ", Player2.Hp)
		}
	}
}
