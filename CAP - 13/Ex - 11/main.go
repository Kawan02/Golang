/*
- Uma das melhores maneiras de aprender é ensinando.
- Para este exercício escolha o seu código favorito dentre os que vimos estudando funções. Pode ser das aulas ou da seção de exercícios. Então:
    - Faça download e instale isso aqui: https://obsproject.com/
    - Grave um vídeo onde *você* ensina o tópico em questão
    - Faça upload do vídeo no YouTube
    - Compartilhe o vídeo no Twitter e me marque no tweet (@ellenkorbes)o que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/
package main

import "fmt"

func main() {
	def()
}

func def() { // defer => ele basicamente deixa tudo por ultimo e funciona para fechar abas/janelas que os usuarios não estão usando, mas esqueceram de fechar.
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

	/*
		Mais ou menos como funciona
		AbreArquivo()
		defer FecharArquivo()

		UsarOsDados{qualquercoisa}
	*/

}
