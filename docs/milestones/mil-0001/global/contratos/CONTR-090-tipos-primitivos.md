<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../mil-0001.md).

# CONTR-090 — Tipos primitivos compartilhados

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** global.

Este catálogo define somente tipos usados nas portas. Não contém contrato vazio nem fluxo fictício. Strings de identificadores são opacas; as regras de cada campo pertencem ao módulo consumidor.

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-090
  modulo: global
  estado: proposto
  tipos:
  - id: TIP-0100
    nome: Texto
    base: texto
  - id: TIP-0101
    nome: Inteiro
    base: inteiro
  - id: TIP-0102
    nome: Booleano
    base: booleano
  contratos: []
  excecoes: []
  fluxos: []
```
