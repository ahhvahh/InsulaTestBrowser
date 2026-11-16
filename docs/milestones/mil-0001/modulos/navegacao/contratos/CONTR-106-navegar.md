<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-106 — Navegar para uma URL

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** navegacao.

## Porta pública

Resolver a conexão privada da sessão e navegar via adaptador Chromium; controlar redirecionamentos e timeout.

### CONTR-106 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| tab_id | texto | Sim |
| url | texto | Sim |
| timeout_ms | inteiro | Sim |
### CONTR-107 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| tab_id | texto | Sim |
| url | texto | Sim |
| title | texto | Sim |
| document_revision | inteiro | Sim |
### CONTR-108 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

Sessão pertence ao contexto autorizado e está pronta; aba pertence à sessão; URL HTTP(S) é permitida pela política de rede; timeout_ms positivo e dentro da cota efetiva. Adaptador reaplica política de destino a redirecionamentos e resolução DNS.

Aplicação explícita no [fluxo](../fluxos/FLUX-0300-navegar.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-106
  modulo: navegacao
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-106
    nome: EntradaNavegarparaumaURL
    campos:
    - id: CMP-1012
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
    - id: CMP-1013
      nome: tab_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
    - id: CMP-1014
      nome: url
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
    - id: CMP-1015
      nome: timeout_ms
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
  - id: CONTR-107
    nome: SaidaNavegarparaumaURL
    campos:
    - id: CMP-1016
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
    - id: CMP-1017
      nome: tab_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
    - id: CMP-1018
      nome: url
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
    - id: CMP-1019
      nome: title
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
    - id: CMP-1020
      nome: document_revision
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
  - id: CONTR-108
    nome: FalhaNavegarparaumaURL
    campos:
    - id: CMP-1021
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0300
        documento: ../fluxos/FLUX-0300-navegar.md#sdd
  excecoes: []
  fluxos: []
```
