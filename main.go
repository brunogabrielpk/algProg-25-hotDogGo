package main 

import "fmt"

/* Itens do Hot-dog
 Salsicha, pão e molho [obrigatórios]
 Ervilha, Milho, bacon, calabresa, bacon, purê
 Valores: 
	Pão = R$ 1,00	
	Molho = R$ 0,50	
	Salcisha = R$ 1,00	
	Ervilha = R$ 0,25	
	Milho = R$ 0,25
	Bacon = R$ 2,00	
	Calabresa = R$ 1,50	
	Purê = R$ 1,00	
	****
	Margem de lucro inicial : 200%
 */

func main() {
var entradaUser int

	// Mostrar o menu para o usuários
	for {
		fmt.Println("Bem vindo ao sistem de pedidos HOTDOGS-ENGENHARIA-AEMS")
		fmt.Println("Selecione o número da opção desejada")
		fmt.Println("[1] - Pedir um Hot-Dog")
		fmt.Println("[2] - Definir a margem de lucro")
		fmt.Println("[3] - Sair")
		fmt.Scanln(&entradaUser)

		if entradaUser == 3 {
			break
		} else {
			fmt.Println("Não foi dessa vez")
		}
	}
}
