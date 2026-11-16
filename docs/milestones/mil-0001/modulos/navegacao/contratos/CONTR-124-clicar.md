<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-124 — Clicar em elemento

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** navegacao.

## Porta pública

Resolver referência opaca do elemento e acionar clique autorizado; devolver confirmação da ação e revisão atual.

### CONTR-124 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| tab_id | texto | Sim |
| element_ref | texto | Sim |
| document_revision | inteiro | Sim |
| timeout_ms | inteiro | Sim |
### CONTR-125 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| clicked | booleano | Sim |
| document_revision | inteiro | Sim |
### CONTR-126 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

Sessão e aba autorizadas; referência pertence à aba e revisão atual; elemento clicável; timeout positivo dentro da cota; rejeitar referência obsoleta. Não receber JavaScript nem valor de formulário.

Aplicação explícita no [fluxo](../fluxos/FLUX-0301-clicar.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-124
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
  - id: CONTR-124
    nome: EntradaClicaremelemento
    campos:
    - id: CMP-1063
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
    - id: CMP-1064
      nome: tab_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
    - id: CMP-1065
      nome: element_ref
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
    - id: CMP-1066
      nome: document_revision
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
    - id: CMP-1067
      nome: timeout_ms
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
  - id: CONTR-125
    nome: SaidaClicaremelemento
    campos:
    - id: CMP-1068
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
    - id: CMP-1069
      nome: clicked
      tipo:
        id: TIP-0102
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
    - id: CMP-1070
      nome: document_revision
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
  - id: CONTR-126
    nome: FalhaClicaremelemento
    campos:
    - id: CMP-1071
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0301
        documento: ../fluxos/FLUX-0301-clicar.md#sdd
  excecoes: []
  fluxos: []
```
