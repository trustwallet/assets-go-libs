# Dependency Graph

<!-- sdd-knowledge-generated -->

## Feature Dependencies

```mermaid
flowchart LR
  file --> strings
  image --> http
  path --> file
  validation --> strings
  validation --> http
  validation --> file
```

## External Dependencies

| Package | Import Count |
|---------|--------------|
| `github.com` | 14 |
| `golang.org` | 1 |

