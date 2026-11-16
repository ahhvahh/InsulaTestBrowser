<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-133 — Persistir catálogo de artefato

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** persistencia.

## Porta pública

Inserção parametrizada idempotente por artifact_id; persistir metadados e storage_key, nunca bytes; conflito de sessão ou checksum falha.

### CONTR-133 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| artifact_id | texto | Sim |
| storage_key | texto | Sim |
| mime_type | texto | Sim |
| size_bytes | inteiro | Sim |
| sha256 | texto | Sim |
### CONTR-134 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| artifact_id | texto | Sim |
| saved | booleano | Sim |
### CONTR-135 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

Sessão existe; artifact_id único; tamanho não negativo; chave relativa validada; digest hexadecimal de 64 caracteres. Calcular created_at e purge_after com relógio e política de retenção injetados pelo serviço.

Aplicação explícita no [fluxo](../fluxos/FLUX-0502-salvar-artefato.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-133
  modulo: persistencia
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-133
    nome: EntradaPersistircatálogodeartefato
    campos:
    - id: CMP-1093
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
    - id: CMP-1094
      nome: artifact_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
    - id: CMP-1095
      nome: storage_key
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
    - id: CMP-1096
      nome: mime_type
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
    - id: CMP-1097
      nome: size_bytes
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
    - id: CMP-1098
      nome: sha256
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
  - id: CONTR-134
    nome: SaidaPersistircatálogodeartefato
    campos:
    - id: CMP-1099
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
    - id: CMP-1100
      nome: artifact_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
    - id: CMP-1101
      nome: saved
      tipo:
        id: TIP-0102
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
  - id: CONTR-135
    nome: FalhaPersistircatálogodeartefato
    campos:
    - id: CMP-1102
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0502
        documento: ../fluxos/FLUX-0502-salvar-artefato.md#sdd
  excecoes: []
  fluxos: []
```
