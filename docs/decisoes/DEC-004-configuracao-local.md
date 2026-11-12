<!-- A-BIOMA-DOC-STANDARD -->
# DEC-004 — Arquivo de configuração ao lado do binário ou em pasta informada

**Decisão:** adotada como requisito documental.  
**Implementação:** não verificada.

## Contexto

O InsulaTestBrowser precisa administrar SQLite, acesso ao serviço, referências a credenciais, perfis de contêiner e preparação de imagens sem espalhar parâmetros operacionais em diversas opções de linha de comando.

## Decisão

Um único arquivo `config.yaml` contém a configuração do serviço. Por padrão, a aplicação o procura na mesma pasta do executável; com `--config-dir`, procura exclusivamente na pasta indicada. Caminhos relativos do conteúdo são interpretados em relação à pasta do arquivo.

A configuração deve incluir banco SQLite, acesso, referências seguras a credenciais, perfis de instância a partir de imagens existentes, instruções para preparar um contêiner novo e autorização explícita para gerar imagem OCI a partir de contêiner de **preparação** iniciado.

## Consequências

- A instalação é previsível e dispensa múltiplos argumentos operacionais.
- Erros de configuração devem impedir a inicialização, sem valores implícitos que reduzam segurança.
- Credenciais devem ficar protegidas fora do YAML e fora das imagens de navegador.
- Capturar contêiner de preparação permite clonar o ambiente configurado; é proibido salvar como imagem uma sessão que já navegou e pode conter dados privados.
- O formato `container-setup.yaml` permanece pendente de detalhamento e testes.

## Referências

- [Configuração de referência](../referencia/REF-001-configuracao.md)
- [Inicialização CLI](../interfaces/terminal/INTF-TERM-001-inicializacao.md)
- [Preparação de imagens](../fluxos/FLUX-002-preparacao-imagem.md)
