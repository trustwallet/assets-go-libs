# assets-go-libs

> The library for automation of the asset adding process at [assets repo](https://github.com/trustwallet/assets).

## What this repo is

- **Domain**: Go library for Trust Wallet assets repository automation — validation, path parsing, and file classification for blockchain token/coin metadata additions.
- **Route here**: `ValidateAsset`, `ValidateCoin`, `ValidateTokenList`, file path type classification (`TypeAssetFolder`, `TypeChainInfoFile`, …), EVM/TRON/Cosmos/Waves address validation, PNG logo dimension/size checks, assets-manager HTTP client, tokenlist cross-validation against `info.json`, validator-list entry validation.
- **Do not route here**: the `trustwallet/assets` repo itself (the JSON/PNG content); chain/coin definitions (`go-primitives`); HTTP client plumbing (`go-libs`); the assets-manager service implementation; any mobile, extension, or backend feature code.
- **Consumers**: CI scripts and automation tooling in `trustwallet/assets`; any Trust Wallet service that needs to validate asset or coin metadata before accepting a PR.
- **Ships**: Go module `github.com/trustwallet/assets-go-libs` — importable packages: `validation`, `validation/info`, `validation/tokenlist`, `validation/list`, `file`, `path`, `http`, `image`, `strings`, `client/assets-manager`. No binary.
- **Agent map**: validation rules → `knowledge/architecture/validate-asset-explain.md`; file type system → `knowledge/architecture/file-domain.md`; data models → `knowledge/architecture/data/models.md`; add/change a validator → `validation/info/fields_validators.go`; address rules → `validation/address.go`.

## Knowledge Map

For the structured knowledge base, see [knowledge/constitution.md](knowledge/constitution.md).

- [ci](knowledge/ci/index.md) — CI/CD pipelines, deployment automation, and release processes
- [code-conventions](knowledge/code-conventions/index.md) — Code conventions, style rules, and decision records
- [patterns](knowledge/patterns/index.md) — Coding patterns, recipes, and proven approaches

- [architecture](knowledge/architecture/index.md) — Architecture
- [build](knowledge/build/index.md) — Build
- [features](knowledge/features/index.md) — Features
- [security](knowledge/security/index.md) — Security

## Learnings

This repo may keep a living archive of incident-derived rules in ~~[`learnings/`](learnings/)~~ — each file a postmortem of a real bug or a non-obvious pattern that bit once and would bite again: root cause, the rule that prevents recurrence, and tags for matching. The folder is **optional and may be absent** — create it the first time you have a learning worth saving.

**Before** investigating any bug, regression, or "weird behavior", *if a `learnings/` directory exists*:

1. Search the frontmatter directly — it's the source of truth and always present:
   - `grep -ril "<keyword>" learnings/` — matches the frontmatter `tags:`/`summary:` + body.
   - `ls learnings/ | grep -i "<keyword>"` — matches the slug-style filename.
   - Skim each match's `summary:` line to decide whether to read the full body.
2. For a topic-organized ToC (grouped by surface + a tag index), open `learnings/index.md`. It is a **generated** artifact that garden **always regenerates** from frontmatter — never hand-edit it (any edit is discarded next run). Depending on the repo it's either gitignored (a derived artifact) or committed; either way it can be stale if a learning file changed without a regen, so prefer reading the learning files' frontmatter over trusting it blindly.
3. Found a match? **Read it before forming a hypothesis** — a 30-second read can turn a 2-hour investigation into a 5-minute fix.
4. Every file has frontmatter (`title`, `date`, `area`, `files`, `symptom`, `tags`, `summary`; `pr` when tied to a specific PR). `area` drives the index's surface grouping; `tags` drive its tag index.

**After** any fix, feature, or non-trivial change — if you learned something not already obvious from the code:

1. Add a new file `learnings/<slug>.md` with the frontmatter above, then a body covering: the symptom, the root cause (the actual mechanism, not just "the bug"), why prior fixes weren't enough if applicable, the rule going forward, and any regression guards. Create the `learnings/` folder if it doesn't exist yet.
2. If the learning extends an existing entry, edit that file instead of creating a duplicate.
3. Make the new file's `area`, `tags`, and `summary` accurate — those drive both `grep` and the generated index (`area` → its surface grouping, `tags` → its tag index, `summary` → its hook). **Never hand-edit `learnings/index.md`** — it's generated and always regenerated; edit the learning file's frontmatter instead.
4. Commit the learning **in the same PR as the fix** — never as a follow-up.

The bar: would a future agent save time by reading this before touching the same surface? If yes, write it; if it would just say "read the diff," skip it. Don't ask which learnings to capture — commit every candidate that clears the bar.

## Repository Knowledge Scope

This repo's `knowledge/` covers: **ci, code-conventions, patterns**

Topics NOT documented locally: architecture, build, conventions, core-libs, decisions, design, features, git-conventions, guides, observability, product, quality, security, tests, workflows, brand, business, legal, hr, prompts, api, specs, libs, components, references

## Constraints

- [TODO: Add project-specific constraints]

<!-- sdd-knowledge-generated -->
