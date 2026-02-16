package main

import (
	"bytes"
	"fmt"
	"os"
    "strings"
    "io"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Uso: viewer <arquivo> [termo] [--hex]")
		return
	}

	caminho := os.Args[1]

	var searchTerm string
	var forceHex bool

	for _, arg := range os.Args[2:] {
		if arg == "--hex" {
			forceHex = true
		} else {
			searchTerm = arg
		}
	}

	dados, err := os.ReadFile(caminho)
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}

	binario := ehBinario(dados)

	if forceHex || binario {

		if searchTerm != "" {
			fmt.Println("Buscando em modo binário...")
			buscarBytes(dados, searchTerm)
			fmt.Println()
		}

		mostrarHexStreaming(caminho)

	} else {

		mostrarTextoBytes(dados, searchTerm)

	}
}

func esperarComando() byte {
    fmt.Print("\nENTER=linha SPACE=página q=sair")

    buf := make([]byte, 1)
    os.Stdin.Read(buf)

    if buf[0] == '\r'{
        return '\n'
    }

    return buf[0]
}

func ehBinario(dados []byte) bool {

    limite := 512
    if len(dados) < limite {
        limite = len(dados)
    }

    naoImprimiveis := 0

    for _, b := range dados[:limite] {

        if b == 0 {
            return true
        }

        if (b < 32 || b > 126) &&
            b != '\n' &&
            b != '\r' &&
            b != '\t' {

            naoImprimiveis++
        }
    }

    proporcao := float64(naoImprimiveis) / float64(limite)

    return proporcao > 0.30
}

func buscarBytes(dados []byte, termo string) {
	alvo := []byte(termo)

	offset := 0
	for {
		i := bytes.Index(dados[offset:], alvo)
		if i == -1 {
			break
		}

		pos := offset + i
		fmt.Printf("Encontrado no offset 0x%X (%d)\n", pos, pos)

		offset = pos + 1
	}
}

func mostrarTextoBytes(dados []byte, termo string) {

    linhas := strings.Split(string(dados), "\n")

    const pageSize = 20
    pos := 0

    for {
        fmt.Print("\033[H\033[2J")

        linhasMostradas := 0

        for i := pos; i < len(linhas) && linhasMostradas < pageSize; i++ {

            texto := linhas[i]

            if termo == "" || termo == " " || strings.Contains(texto, termo) {
                fmt.Printf("%4d | %s\n", i+1, texto)
                linhasMostradas++
            }
        }

        if linhasMostradas == 0 {
            fmt.Println("(nenhuma linha corresponde ao filtro)")
        }

        fmt.Print("\nENTER=linha SPACE=página q=sair → ")

        cmd := lerComando()

        switch cmd {
        case 'q':
            return

        case ' ':
            pos += pageSize
            if pos >= len(linhas) {
                pos = len(linhas) - 1
            }

        case '\n':
            if pos < len(linhas)-1 {
                pos++
            }
        }

        if pos < 0 {
            pos = 0
        }

        if pos > len(linhas)-1 {
            pos = len(linhas) - 1
        }
    }
}

func mostrarHexStreaming(caminho string) {

    arquivo, err := os.Open(caminho)
    if err != nil {
        fmt.Println("Erro ao abrir arquivo:", err)
        return
    }
    defer arquivo.Close()

    const largura = 16
    buffer := make([]byte, largura)

    offset := 0
    linhas := 0
    const pageSize = 20

    for {
        n, err := arquivo.Read(buffer)
        if n == 0 {
            break
        }

        linha := buffer[:n]

        fmt.Printf("%08X  ", offset)

        for i := 0; i < largura; i++ {
            if i < len(linha) {
                fmt.Printf("%02X ", linha[i])
            } else {
                fmt.Print("   ")
            }
        }

        fmt.Print(" ")

        for _, b := range linha {
            if b >= 32 && b <= 126 {
                fmt.Printf("%c", b)
            } else {
                fmt.Print(".")
            }
        }

        fmt.Println()

        offset += n
        linhas++

        if linhas >= pageSize {
            cmd := esperarComando()

            if cmd == 'q' {
                return
            }
            linhas = 0
        }

        if err == io.EOF {
            break
        }
    }
}

func lerComando() byte {
    var b [1]byte
    os.Stdin.Read(b[:])

    if b[0] == '\r' {
        return '\n'
    }

    return b[0]
}

func lerLinhas(caminho string) ([]string, error){
    arquivo, err := os.ReadFile(caminho)
    if err != nil {
        return nil, err
    }

    conteudo := string(arquivo)

    linhas := strings.Split(conteudo, "\n")

    return linhas, nil
}
