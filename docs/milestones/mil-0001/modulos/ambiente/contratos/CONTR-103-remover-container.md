<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-103 — Remover contêiner

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** ambiente.

## Porta pública

Parar e remover somente o recurso rotulado com a sessão; ausência prévia é sucesso idempotente.

### CONTR-103 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| container_id | texto | Sim |
### CONTR-104 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| removed | booleano | Sim |
### CONTR-105 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

IDs não vazios; o container_id foi associado à session_id pelo serviço. Adaptador verifica rótulo real antes de remover.

Aplicação explícita no [fluxo](../fluxos/FLUX-0201-remover-container.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-103
  modulo: ambiente
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-103
    nome: EntradaRemovercontêiner
    campos:
    - id: CMP-1007
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0201
        documento: ../fluxos/FLUX-0201-remover-container.md#sdd
    - id: CMP-1008
      nome: container_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0201
        documento: ../fluxos/FLUX-0201-remover-container.md#sdd
  - id: CONTR-104
    nome: SaidaRemovercontêiner
    campos:
    - id: CMP-1009
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0201
        documento: ../fluxos/FLUX-0201-remover-container.md#sdd
    - id: CMP-1010
      nome: removed
      tipo:
        id: TIP-0102
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0201
        documento: ../fluxos/FLUX-0201-remover-container.md#sdd
  - id: CONTR-105
    nome: FalhaRemovercontêiner
    campos:
    - id: CMP-1011
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0201
        documento: ../fluxos/FLUX-0201-remover-container.md#sdd
  excecoes: []
  fluxos: []
```
