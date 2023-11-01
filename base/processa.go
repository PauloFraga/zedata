package base

import (
	"fmt"
	"os"
		
)

// type binario struct {
 
// 	arquivo_origem  string
// 	arquivo_destino  string
// 	tamanho_original int64
// 	tamanho_compressed int64
// 	blocos int64

// }

//Compaction
// func CompactarParaGzip(arquivoOrigem string) (binario, error) {
// 	var resp binario
	
// 	// Abra o arquivo de origem para leitura
// 	arquivoEntrada, err := os.Open(arquivoOrigem)
// 	if err != nil {
// 		return  resp ,  err
// 	}
// 	defer arquivoEntrada.Close()

// 	arquivoDestino:= arquivoOrigem + ".gz"

// 	// Crie o arquivo de destino para escrita
// 	arquivoSaida, err := os.Create(arquivoDestino)
// 	if err != nil {
// 		return resp ,  err 
// 	}
// 	defer arquivoSaida.Close()
	
// 	// Crie um escritor gzip usando o arquivo de destino
// 	writer := gzip.NewWriter(arquivoSaida)
	
	
// 	writer.Header.Name = arquivoOrigem
// 	//writer.Header.Comment = infoEntrada.Name() + "|" + strconv.FormatInt(infoEntrada.Size(),10)
	
// 	//fmt.Println("Compactando: ... ", arquivoSaida.Name())
// 	_, err = io.Copy(writer, arquivoEntrada)
// 	if err != nil {
// 		return resp ,  err
// 	}
	
// 	writer.Close()
// 	infoEntrada, _ := arquivoEntrada.Stat()
// 	infoSaida, _ := arquivoSaida.Stat()

// 	resp.arquivo_destino = arquivoSaida.Name()
// 	resp.arquivo_origem = arquivoEntrada.Name()
// 	resp.tamanho_original = infoEntrada.Size()
// 	resp.tamanho_compressed = infoSaida.Size()
	
// 	println(resp.arquivo_origem)


	
// 	return resp, nil
// }


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
 


// func main() {

// 	pathFile := "."
	
//     arquivos, err := filepath.Glob(filepath.Join(pathFile, "*.pdf"))
//     if err != nil {
//         fmt.Println("Erro ao listar arquivos:", err)
//         return
//     } 


// 	var inicioBlocos int64 = 0

//     for _, arquivo := range arquivos {
//         info, err := os.Stat( arquivo )
//         if err != nil {
//             fmt.Println("Erro ao obter informações do arquivo:", err)
//             continue
//         }
//         if !info.IsDir() { 
//             // if  strings.Contains( info.Name() ,".gz") {
// 			// 	continue
				
// 			// 	} else {
					
// 			//COMPACTAR ARQUIVOS
// 			//===========================================================
// 				resp , err :=  CompactarParaGzip(arquivo)
// 				if err != nil {
// 					fmt.Println(info)
// 				}

// 				blocosArquivo, err :=  AppendToBinaryFile(arquivo + ".gz")
// 				if err != nil {
// 					fmt.Println(err) 
// 				  }

// 				//   println("apagando Temp:",resp.arquivo_destino )
// 				  os.Remove(resp.arquivo_destino)

// 				  if inicioBlocos == 0 {
// 					GravaIndex("Nome:" + resp.arquivo_destino + "|"+ "Origem:" + strconv.FormatInt(resp.tamanho_original,10) + "|Compress:" + strconv.FormatInt(resp.tamanho_compressed,10)  + "|Bloco:"  + strconv.FormatInt( blocosArquivo,10) + "|Inicio:0"  )
// 					inicioBlocos = blocosArquivo

// 				  } else {
// 					  GravaIndex("Nome:" + resp.arquivo_destino + "|"+ "Origem:" + strconv.FormatInt(resp.tamanho_original,10) + "|Compress:" + strconv.FormatInt(resp.tamanho_compressed,10)  + "|Bloco:"  + strconv.FormatInt( blocosArquivo,10) + "|Inicio:"  + strconv.FormatInt( inicioBlocos,10) )
// 					  inicioBlocos = inicioBlocos + blocosArquivo
// 				  }

// 					//   GravaIndex(resp.arquivo_origem + "|" + strconv.FormatInt(resp.tamanho_original,10)  + "|"  + strconv.FormatInt(resp.tamanho_compressed,10)  + "|" + strconv.FormatInt( qteBloco,10) )

// 				// }
				
//         } 
//     }

// } 