<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-103 — Dados da remoção de contêiner

**Estado:** em refinamento. **Versão aprovada:** nenhuma. **Módulo:** ambiente.

## Objetivo

Definir a solicitação, o resultado e a falha da remoção. Este arquivo agrupa três estruturas com IDs preservados: CONTR-103, CONTR-104 e CONTR-105. A execução pertence ao [FLUX-0201](../fluxos/FLUX-0201-remover-container.md).

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

Os identificadores são textos obrigatórios e não vazios. A resposta contém o identificador da sessão e uma confirmação booleana. A falha contém um código textual sanitizado.

As regras dos campos referenciam [FLUX-0201](../fluxos/FLUX-0201-remover-container.md) de forma declarativa. A ordem de validação, a verificação do vínculo com o contêiner e o tratamento de falhas ficam definidos nesse fluxo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).


## Formato piloto

Este documento usa **Spec Rules 0.2.0**, proposta local da MIL-0001. A versão identifica o novo formato, não uma release do plugin nem compatibilidade com o schema SDD 2.0. As referências abreviadas são resolvidas pelo [catálogo do piloto](../../../referencia/REF-003-spec-rules.md). Estado documental: **em refinamento**; implementação não comprovada.
<a id="sdd"></a>
<a id="spec-rules"></a>
## Spec Rules

<!-- A-BIOMA-SPEC-RULES:0.2.0 -->

```yaml
CONTR-103:
  formato: "a-bioma-spec-rules"
  versao: "0.2.0"
  tipo_documento: "contrato"
  modulo: "ambiente"
  estado: "em refinamento"
  implementacao: "nao implementado"
  objetivo: "Definir solicitação, resultado e falha da remoção de contêiner."
  acesso:
    tipo: "porta interna"
    contexto: "Injetado pelo serviço; IDs não concedem autorização."
    politica: "Chamadas públicas permitidas pela ARQ-005."
  papeis:
    entrada: "CONTR-103"
    saida: "CONTR-104"
    falha: "CONTR-105"
  contratos:
    CONTR-103:
      nome: "EntradaRemoverConteiner"
      campos:
        - id: "CMP-1007"
          nome: "session_id"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
        - id: "CMP-1008"
          nome: "container_id"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
    CONTR-104:
      nome: "SaidaRemoverConteiner"
      campos:
        - id: "CMP-1009"
          nome: "session_id"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
        - id: "CMP-1010"
          nome: "removed"
          tipo: "@TIP-0102"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
    CONTR-105:
      nome: "FalhaRemoverConteiner"
      campos:
        - id: "CMP-1011"
          nome: "code"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
  comportamento:
    doc: "@[ambiente](FLUX-0201)"
    relacao: "Define validação e execução; referência declarativa."
  restricoes:
    - "Identificadores obrigatórios e não vazios."
    - "Resposta preserva session_id e confirma removed=true."
    - "Falha contém somente código sanitizado."
  pendencias:
    - "Schema específico de contrato e migração geral pendentes."
    - "Sem implementação comprovada ou versão aprovada."
```
