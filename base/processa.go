package base

import (
	"fmt"
	"os"
		
)


// appendEntrada
func AppendToBinaryFile(fileNameArquivo string) (int64, error) {

	filedata, err := os.OpenFile("zetadata.dat", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		 
		return 0,err
	}
	defer filedata.Close()

	// Tamanho do bloco desejado
	blockSize := 4096

	// Crie um buffer para armazenar o bloco
	buffer := make([]byte, blockSize)

	file, err := os.Open(fileNameArquivo)
	if err != nil {
		// fmt.Println("Erro ao abrir o arquivo:", err)
		fmt.Println("Erro ao abrir o arquivo:")
		return 0, err
	}
	defer file.Close()


	var qtdeBlocos int64 = 0
	for {
			// Leia um bloco do arquivo
			n, err := file.Read(buffer)
			if err != nil {
				fmt.Println("Erro ao ler o arquivo:", err)
				break
			}

			// Verifique se o bloco é o último
			if n < blockSize {
				// Preencha o restante do bloco com zeros
				for i := n; i < blockSize; i++ {
					buffer[i] = 0
					
				}
			}

			_, err = filedata.Write(buffer)
			if err != nil {
				return 0, err
				}

			
				qtdeBlocos++
		
				// Se lemos menos bytes do que o tamanho do bloco, isso significa que atingimos o final do arquivo
			if n < blockSize {
				//fmt.Println(n)
				break
			}
		}


		//gravaIndex(fileNameArquivo + "|" + strconv.Itoa( qtdeBlocos) + "\n")

	return qtdeBlocos, nil
}

// gravaIndex 
func GravaIndex( conteudo string)  {
	
	nomeArquivo := "Index.lst"

	// Abra o arquivo em modo de anexação (append)
	arquivo, err := os.OpenFile(nomeArquivo, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		//return
	}
	defer arquivo.Close()


	// Escreva o conteúdo no arquivo
	_, err = arquivo.WriteString(conteudo +"\n")	
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
}
 
