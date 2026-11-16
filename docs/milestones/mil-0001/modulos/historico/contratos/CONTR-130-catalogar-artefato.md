<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-130 — Catalogar artefato

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** historico.

## Porta pública

Validar metadados confirmados pelo produtor confiável e persistir localização do arquivo.

### CONTR-130 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| artifact_id | texto | Sim |
| storage_key | texto | Sim |
| mime_type | texto | Sim |
| size_bytes | inteiro | Sim |
| sha256 | texto | Sim |
### CONTR-131 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| artifact_id | texto | Sim |
| saved | booleano | Sim |
### CONTR-132 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

IDs opacos; storage_key relativa à raiz privada da sessão, sem traversal; tamanho não negativo; sha256 hexadecimal de 64 caracteres; MIME validado pelo produtor. Apenas artefato concluído pode ser catalogado como ready.

Aplicação explícita no [fluxo](../fluxos/FLUX-0401-catalogar-artefato.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-130
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
  - id: CONTR-130
    nome: EntradaCatalogarartefato
    campos:
    - id: CMP-1083
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
    - id: CMP-1084
      nome: artifact_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
    - id: CMP-1085
      nome: storage_key
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
    - id: CMP-1086
      nome: mime_type
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
    - id: CMP-1087
      nome: size_bytes
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
    - id: CMP-1088
      nome: sha256
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
  - id: CONTR-131
    nome: SaidaCatalogarartefato
    campos:
    - id: CMP-1089
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
    - id: CMP-1090
      nome: artifact_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
    - id: CMP-1091
      nome: saved
      tipo:
        id: TIP-0102
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
  - id: CONTR-132
    nome: FalhaCatalogarartefato
    campos:
    - id: CMP-1092
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0401
        documento: ../fluxos/FLUX-0401-catalogar-artefato.md#sdd
  excecoes: []
  fluxos: []
```
