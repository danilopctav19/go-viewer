# Go Viewer

Um visualizador de arquivos em **Go**, capaz de:

- Exibir arquivos de texto com paginação
- Detectar automaticamente arquivos binários
- Mostrar conteúdo em hexadecimal (estilo hexdump)
- Buscar termos em arquivos texto e binários

Projeto simples, rápido e sem dependências externas.

---

## Recursos

✔ Detecção automática de binário  
✔ Modo hexadecimal (`hex viewer`)  
✔ Busca por strings  
✔ Paginação no terminal  
✔ Funciona em Linux / Windows / macOS  

---

## 🚀 Compilação

É necessário ter o Go instalado.

```bash
go build viewer.go
```

## Uso

**(Para Windows, troque o ./viewer por viewer.exe)**

**Visualizar arquivo em texto**

Exemplo (Linux):

```bash
./viewer arquivo.txt
```

**Buscar termo em texto**

Exemplo:

```bash
./viewer arquivo.txt termo
```

**Forçar modo hexadecimal**

Exemplos:

```bash
./viewer arquivo.bin --hex
./viewer arquivo.exe --hex
```

**Buscar em binário**

```bash
./viewer arquivo.bin termo
```

## Controles

Durante a visualização:
- ENTER → próxima linha
- SPACE → próxima página
- q → sair

## Como funciona
O programa:
- Lê o arquivo informado
- Detecta se é texto ou binário
- Escolhe automaticamente o modo de exibição
- Permite busca por termos
- Pagina a saída no terminal
- Arquivos binários são exibidos em formato hexadecimal + ASCII.

## Objetivo do projeto
Este projeto foi criado para estudo de:
- Leitura de arquivos em Go
- Manipulação de bytes
- Detecção heurística de binários
- Paginação em terminal
- Estruturação de CLI tools

## Licença
Uso livre para estudo e modificação.
