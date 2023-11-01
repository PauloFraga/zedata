package compress

import (
	"compress/gzip"
	"io"
	"os"
)



type binario struct {

	arquivo_origem  string
	arquivo_destino  string
	tamanho_original int64
	tamanho_compressed int64
	blocos int64

}

func CompactarParaGzip(arquivoOrigem string) (binario, error) {
	var resp binario

	// Abra o arquivo de origem para leitura
	arquivoEntrada, err := os.Open(arquivoOrigem)
	if err != nil {
		return resp, err
	}
	defer arquivoEntrada.Close()

	arquivoDestino := arquivoOrigem + ".gz"

	// Crie o arquivo de destino para escrita
	arquivoSaida, err := os.Create(arquivoDestino)
	if err != nil {
		return resp, err
	}
	defer arquivoSaida.Close()

	// Crie um escritor gzip usando o arquivo de destino
	writer := gzip.NewWriter(arquivoSaida)

	writer.Header.Name = arquivoOrigem
	//writer.Header.Comment = infoEntrada.Name() + "|" + strconv.FormatInt(infoEntrada.Size(),10)

	//fmt.Println("Compactando: ... ", arquivoSaida.Name())
	_, err = io.Copy(writer, arquivoEntrada)
	if err != nil {
		return resp, err
	}

	writer.Close()
	infoEntrada, _ := arquivoEntrada.Stat()
	infoSaida, _ := arquivoSaida.Stat()

	resp.arquivo_destino = arquivoSaida.Name()
	resp.arquivo_origem = arquivoEntrada.Name()
	resp.tamanho_original = infoEntrada.Size()
	resp.tamanho_compressed = infoSaida.Size()

	println(resp.arquivo_origem)

	return resp, nil
}