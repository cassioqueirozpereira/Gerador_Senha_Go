package main // Declara que este arquivo pertence ao pacote 'main'.
              // Em Go, 'main' é um pacote especial que indica que o programa é executável.
              // A função 'main' dentro deste pacote é o ponto de entrada do programa.

import ( // Inicia o bloco de importação de pacotes.
	// Pacote 'fmt' (format) é usado para formatação de E/S (entrada/saída),
	// como imprimir texto no console (Println, Printf) e ler entrada do usuário (Scanf).
	"fmt"
	// Pacote 'math/rand' é usado para gerar números pseudoaleatórios.
	// É fundamental para a aleatoriedade da senha.
	"math/rand"
	// Pacote 'time' é usado para funcionalidades relacionadas a tempo,
	// como obter a hora atual para 'semear' o gerador de números aleatórios
	// e para pausar a execução (Sleep).
	"time"
)

// main é a função principal que é executada quando o programa Go é iniciado.
func main() {
	// Declara uma variável inteira 'tamanhoSenha' e a inicializa com 0.
	// Esta variável armazenará o comprimento desejado da senha digitado pelo usuário.
	tamanhoSenha := 0

	// Inicia um loop infinito ('for {}').
	// Este loop continuará pedindo o tamanho da senha até que uma entrada válida seja fornecida.
	for {
		// fmt.Printf imprime a mensagem no console. O '\n' não é usado no final
		// para que a entrada do usuário seja na mesma linha que a pergunta.
		fmt.Printf("Digite com quantos caracteres vai querer a sua senha: ")

		// fmt.Scanf lê a entrada do usuário formatada.
		// "%d" indica que esperamos um número inteiro decimal.
		// "&tamanhoSenha" é crucial: o '&' pega o endereço de memória da variável 'tamanhoSenha',
		// permitindo que fmt.Scanf armazene o valor lido diretamente nessa localização de memória.
		fmt.Scanf("%d", &tamanhoSenha)
		fmt.Scanln()
		// Verifica se o tamanho da senha digitado pelo usuário é menor que 8.
		if tamanhoSenha < 8 {
			// Se for menor que 8, imprime uma mensagem de erro.
			fmt.Println("A senha precisa ter no mínimo 8 caracteres")
			// Pausa a execução do programa por 3 segundos para que o usuário possa ler a mensagem de erro.
			time.Sleep(3 * time.Second)
			// Chama a função para limpar a tela do terminal.
			clearScreenANSI()
			}else {
				// Se o tamanho for válido, 'break' sai do loop 'for' infinito.
				break
			}
		}
		
		// Chama a função 'generatePassword', passando o 'tamanhoSenha' validado pelo usuário.
		// O resultado (a senha gerada) é armazenado na variável 'password'.
		password := generatePassword(tamanhoSenha)
		
		// Chama a função para limpar a tela
		clearScreenANSI()

		// Imprime a senha gerada no console.
		fmt.Print("A sua senha é: ", password)
	}
	
	// generatePassword é uma função que cria uma senha aleatória com um comprimento especificado.
// Ela aceita um argumento 'length' (o comprimento da senha desejada) e retorna uma string.
func generatePassword(length int) string {
	// Define strings contendo diferentes conjuntos de caracteres.
	lowerCase := "abcdefghijklmnopqrstuvwxyz" // Letras minúsculas
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // Letras maiúsculas
	numbers := "0123456789"                   // Números
	special := "!@#$%^&*()+?><:{}[]"          // Caracteres especiais
	// Concatena todos os conjuntos de caracteres em uma única string para uso geral.
	allChars := lowerCase + upperCase + numbers + special

	// Cria um slice de bytes para armazenar os caracteres obrigatórios da senha.
	// Garante que a senha tenha pelo menos uma letra maiúscula, um número e um caractere especial.
	mandatory := []byte{
		// Seleciona aleatoriamente um caractere maiúsculo.
		upperCase[rand.Intn(len(upperCase))],
		// Seleciona aleatoriamente um número.
		numbers[rand.Intn(len(numbers))],
		// Seleciona aleatoriamente um caractere especial.
		special[rand.Intn(len(special))],
	}

	// Cria um slice de bytes 'password' com um comprimento inicial.
	// O comprimento é o total desejado menos o número de caracteres obrigatórios.
	// Isso porque os caracteres obrigatórios serão adicionados separadamente.
	// NOTA: Se 'length' for menor que 'len(mandatory)' (3), isso resultará em um pânico
	// ou em um slice de tamanho inválido. A validação do 'tamanhoSenha' no 'main'
	// (mínimo de 8) evita isso, mas uma verificação aqui seria mais robusta.
	password := make([]byte, length-len(mandatory))

	// Preenche o slice 'password' (excluindo os caracteres obrigatórios por enquanto)
	// com caracteres aleatórios de 'allChars'.
	for i := range password {
		// Seleciona um caractere aleatório de 'allChars' e o atribui à posição 'i'.
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	// Adiciona os caracteres obrigatórios ao final do slice 'password'.
	// 'append' retorna um novo slice que contém os elementos originais mais os novos.
	password = append(password, mandatory...)

	// Embaralha os caracteres da senha para garantir que os caracteres obrigatórios
	// não fiquem sempre no final e para aumentar a aleatoriedade geral.
	// 'rand.Shuffle' reorganiza um slice de forma aleatória.
	// 'len(password)' é o tamanho do slice a ser embaralhado.
	// A função anônima (closure) define como trocar dois elementos do slice.
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i] // Troca os elementos nas posições i e j.
	})

	// Converte o slice de bytes 'password' de volta para uma string e a retorna.
	return string(password)
}

// clearScreenANSI limpa a tela do terminal usando sequências de escape ANSI.
// Este método é amplamente suportado em terminais modernos (Linux, macOS, PowerShell).
// Pode não funcionar em terminais Windows legados (Prompt de Comando antigo) sem configuração.
func clearScreenANSI() {
	// "\033[2J" é a sequência ANSI para limpar a tela inteira.
	// "\033[H" é a sequência ANSI para mover o cursor para a posição inicial (topo esquerdo).
	// fmt.Print é usado para que nenhuma nova linha seja adicionada após o comando de limpeza.
	fmt.Print("\033[2J\033[H")
}