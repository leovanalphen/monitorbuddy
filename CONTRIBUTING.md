# Contributing Guidelines

## Branch Naming

When starting work, create a branch from main using this format:

```bash
<type>/<short-description>
```

Where **\<type\>** matches the [Conventional Commits](https://www.conventionalcommits.org/) types:

- **feat** → new feature (e.g. **feat/ddcci-backend)**
- **fix** → bug fix (e.g. **fix/windows-timeout**)
- **docs** → documentation only (e.g. **docs/monitor-control-matrix**)
- **chore** → tooling, config, maintenance (e.g. **chore/ci-workflow**)
- **refactor** → code refactor without new features or fixes (e.g. **refactor/transport-layer**)
- **test** → tests only (e.g. **test/ddcci-parser**)
- **style** → formatting, whitespace, linter fixes (e.g. **style/golangci-lint-fixes**)

## Commit Messages

All commits must follow Conventional Commits:

```bash
<type>[optional scope]: <short description>
```

## Examples

- feat: add Linux DDC/CI backend
- fix(ddcci): correct checksum calculation
- docs: add vendor × feature matrix
- chore(ci): update GitHub Actions workflow
- This format enables automatic changelog generation and semantic versioning in the future.

## Merging Workflow

External contributors should open a PR against **main**.

Keep PRs focused (one feature or fix per PR).

## General Notes

Keep main always in a working, buildable state.

Add or update tests when fixing bugs or adding features.

Update documentation (**/docs**) if your change affects usage or capabilities.
