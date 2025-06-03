# cp-tf-inspect

[![Build & Release](https://github.com/colpal/cp-tf-inspect/actions/workflows/main.yaml/badge.svg)](https://github.com/colpal/cp-tf-inspect/actions/workflows/main.yaml)

A fast, extensible CLI tool for inspecting Terraform workspaces—supports unique module source listing, and more.

---

## ✨ Features

- **Lists all unique module sources** in a Terraform workspace (JSON output)
- Build & release workflow with full automation
- PR preview builds and semantic versioned stable releases (Go, Linux)
- All releases and binaries available as GitHub Releases
- Contributor friendly, with required commit conventions and robust CI

---

## 🚀 Quickstart

**1. Download the latest stable binary**  
Visit [Releases](https://github.com/colpal/cp-tf-inspect/releases/latest) and download `cp-tf-inspect`.

**2. Build from source**
```bash
git clone https://github.com/colpal/cp-tf-inspect.git
cd cp-tf-inspect
go build -o cp-tf-inspect
```

**3. Usage**
```bash
./cp-tf-inspect list-module-source --dir <path-to-terraform-workspace>
```

Example:
```bash
./cp-tf-inspect list-module-source --dir ./test/fixtures/basic
```

---

## 🧪 Testing

**Unit tests:**
```bash
go test ./cmd/...
```

**Integration tests:**
```bash
go build -o cp-tf-inspect
go test ./test/...
```

---

## 🔄 Release & CI/CD Model

- **Every PR:**  
  - Unique preview release with a tag like `pr-<PR_NUMBER>-<RUN_NUMBER>-<RUN_ATTEMPT>`.
- **On every successful merge to main:**  
  - A new automatic versioned release (semver, e.g. `v1.0.0`, `v1.0.1`)
  - "Latest" release always updated.
- **Binaries:**  
  - All stable and PR preview builds have downloadable `cp-tf-inspect` binary.
  - Download URL pattern:
    ```
    https://github.com/colpal/cp-tf-inspect/releases/download/<tag>/cp-tf-inspect
    ```

---

## 👥 Contributing

**We do not support fork-based pull requests.**  
Please push feature branches directly to this repository and open PRs from them targeting `main`.

### Contribution Steps

1. **Create a feature branch in this repository from `main`.**
2. Make your changes, following the commit message guidelines below.
3. Push your branch and open a Pull Request (PR) to `main`.
4. Ensure all Go tests pass. Our GitHub Actions will automatically run checks and build PR preview releases.

### Commit Messages: Conventional Commits

We use [Conventional Commits](https://www.conventionalcommits.org/) for automated release management.

**Every commit meant to trigger a release must start with a valid type and a colon.**

#### Examples:
```
fix: correct terraform parsing bug
feat: add storage module support
chore: update dependencies
```

**Format:**  
`<type>[optional scope]: <short description>`

**Recommended types:**
- `feat`: New feature
- `fix`: Bug fix
- `chore`: Internal/tooling change
- `docs`: Documentation
- `refactor`: Code refactoring
- `test`: Testing only

*Any type is allowed* (with current config), as long as the message starts with `type:`.

---

## 🔒 Branching & Merge Policy

- **Main development happens on `main` only.**
- All features/fixes go through short-lived feature branches pushed directly to this repo.
- PRs must target `main` and pass CI before being merged.

---
