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
	var opcaoUserMenu int
	var userConfirma string
	// sn := false
	margemLucro := 2.0
	itensHD := [8]string{"Pão", "Molho", "Salsicha", "Ervilha", "Milho", "Bacon", "Calabresa", "Purê"}
	//itensHDvalor := [8]float64{1.0, 0.5, 1.0, 0.25, 0.25, 2.0, 1.5, 1.0}
	itensHDconfirma := [8]bool{true, true, true, false, false, false, false, false}

	// Mostrar o menu para o usuários
	for {
		fmt.Println("Bem vindo ao sistem de pedidos HOTDOGS-ENGENHARIA-AEMS")
		fmt.Println("Selecione o número da opção desejada")
		fmt.Println("[1] - Pedir um Hot-Dog")
		fmt.Println("[2] - Definir a margem de lucro")
		fmt.Println("[3] - Sair")
		fmt.Scanln(&opcaoUserMenu)

		if opcaoUserMenu == 1 {
			fmt.Println("Vamos definir os adicionais de seu Hot-Dog")
			// Verifica quais adicionais o usuário deseja
			for pos, item := range itensHD {
				if pos >= 3 {
					fmt.Printf("Deseja adicionar %v?\n", item)
					fmt.Scanln(&userConfirma)
					if userConfirma == "s" || userConfirma == "S" {
						itensHDconfirma[pos] = true
					} else {
						itensHDconfirma[pos] = false
					}
				}
			}
		} else if opcaoUserMenu == 2 {
			fmt.Println("Digite a nova margem de lucro (apenas números)")
			fmt.Scanln(&margemLucro)
		} else if opcaoUserMenu == 3 {
			break
		} else {
			fmt.Println("Escolha uma opção válida")
		}
	}
}
