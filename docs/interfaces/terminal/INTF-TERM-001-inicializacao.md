<!-- A-BIOMA-DOC-STANDARD -->
# INTF-TERM-001 — Inicializar o serviço com arquivo de configuração

**Situação:** Proposto; os comandos são exemplos **não executáveis ainda**.

## Objetivo

Iniciar o binário Go apontando opcionalmente para a pasta que contém `config.yaml`.

## Interface prevista

| Invocação conceitual | Fonte da configuração |
| --- | --- |
| `./insula-test-browser` | `config.yaml` ao lado do executável |
| `./insula-test-browser --config-dir /etc/insula-test-browser` | `/etc/insula-test-browser/config.yaml`, exclusivamente |

A grafia final do executável pode mudar durante a implementação. `--config-dir` recebe um **diretório**, não o caminho do arquivo. O diretório corrente não influencia a busca nem a resolução dos caminhos relativos da configuração.

## Entrada, saída e falhas

| Canal | Comportamento previsto |
| --- | --- |
| Argumentos | `--config-dir` opcional; sem argumento usa a pasta do binário |
| Arquivo | `config.yaml` obrigatório, validado antes de inicializar serviços |
| stdin/TTY | Não exigidos para a inicialização |
| stdout | Informações operacionais não sensíveis, formato a definir |
| stderr | Erros de parâmetros, arquivo, permissões ou serviços, sem conteúdo de segredos |
| Código de saída | Não definido; saída diferente de zero em falha é requisito |

Se a pasta informada não existir, não houver `config.yaml`, o YAML for inválido ou houver credencial indisponível, a inicialização **deve falhar** sem procurar outros caminhos.

A origem e o esquema de configuração estão documentados em [REF-001](../../referencia/REF-001-configuracao.md). O fluxo de sessões iniciado pelo serviço está em [FLUX-001](../../fluxos/FLUX-001-ciclo-de-vida.md), e o fluxo de preparação, em [FLUX-002](../../fluxos/FLUX-002-preparacao-imagem.md).

## Pendências

Ainda faltam assinatura e nome finais do executável, confirmação do protocolo UDS, mapeamento de mensagens stdout/stderr e códigos exatos de saída.
