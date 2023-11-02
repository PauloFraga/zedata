package main

import (
	"fmt"
	"os"
	"github.com/PauloFraga/zetadata/base/processa"
	"github.com/PauloFraga/zetadata/compress/gzip"
	"path/filepath"
	"strconv"
)

func main() {

	pathFile := "."

	arquivos, err := filepath.Glob(filepath.Join(pathFile, "*.pdf"))
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
		return
	}

	var inicioBlocos int64 = 0

	for _, arquivo := range arquivos {
		info, err := os.Stat(arquivo)
		if err != nil {
			fmt.Println("Erro ao obter informações do arquivo:", err)
			continue
		}
		if !info.IsDir() {

			//COMPACTAR ARQUIVOS
			//===========================================================
			resp, err := CompactarParaGzip(arquivo)
			if err != nil {
				fmt.Println(info)
			}

			blocosArquivo, err := AppendToBinaryFile(arquivo + ".gz")
			if err != nil {
				fmt.Println(err)
			}

			//   println("apagando Temp:",resp.arquivo_destino )
			os.Remove(resp.arquivo_destino)

			if inicioBlocos == 0 {
				GravaIndex("Nome:" + resp.arquivo_destino + "|" + "Origem:" + strconv.FormatInt(resp.tamanho_original, 10) + "|Compress:" + strconv.FormatInt(resp.tamanho_compressed, 10) + "|Bloco:" + strconv.FormatInt(blocosArquivo, 10) + "|Inicio:0")
				inicioBlocos = blocosArquivo

			} else {
				GravaIndex("Nome:" + resp.arquivo_destino + "|" + "Origem:" + strconv.FormatInt(resp.tamanho_original, 10) + "|Compress:" + strconv.FormatInt(resp.tamanho_compressed, 10) + "|Bloco:" + strconv.FormatInt(blocosArquivo, 10) + "|Inicio:" + strconv.FormatInt(inicioBlocos, 10))
				inicioBlocos = inicioBlocos + blocosArquivo
			}

		}
	}

}