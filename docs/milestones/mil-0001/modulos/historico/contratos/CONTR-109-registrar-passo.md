<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-109 — Registrar passo de navegação

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** historico.

## Porta pública

Sanitizar metadados e solicitar persistência idempotente de um passo.

### CONTR-109 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| step_id | texto | Sim |
| operation | texto | Sim |
| target | texto | Sim |
| status | texto | Sim |
### CONTR-110 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| step_id | texto | Sim |
| sequence | inteiro | Sim |
### CONTR-111 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

IDs opacos não vazios; operation pertence ao catálogo de operações; status é completed, failed ou cancelled; target não contém credenciais, fragmentos, parâmetros secretos nem valores de formulário.

Aplicação explícita no [fluxo](../fluxos/FLUX-0400-registrar-passo.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-109
  modulo: historico
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-109
    nome: EntradaRegistrarpassodenavegação
    campos:
    - id: CMP-1022
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
    - id: CMP-1023
      nome: step_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
    - id: CMP-1024
      nome: operation
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
    - id: CMP-1025
      nome: target
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
    - id: CMP-1026
      nome: status
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
  - id: CONTR-110
    nome: SaidaRegistrarpassodenavegação
    campos:
    - id: CMP-1027
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
    - id: CMP-1028
      nome: step_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
    - id: CMP-1029
      nome: sequence
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
  - id: CONTR-111
    nome: FalhaRegistrarpassodenavegação
    campos:
    - id: CMP-1030
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0400
        documento: ../fluxos/FLUX-0400-registrar-passo.md#sdd
  excecoes: []
  fluxos: []
```
