# cp-tf-inspect

[![Build & Release](https://github.com/colpal/cp-tf-inspect/actions/workflows/main.yaml/badge.svg)](https://github.com/colpal/cp-tf-inspect/actions/workflows/main.yaml)

A fast, extensible CLI tool for inspecting Terraform workspaces—supports unique module source listing, and more.

---

## 🚀 Quickstart

### 1. Download the latest Linux binary

Go to the [Releases page](https://github.com/colpal/cp-tf-inspect/releases/latest)  
Download the `cp-tf-inspect` binary for Linux.
```bash
curl -LO https://github.com/colpal/cp-tf-inspect/releases/download/vX.Y.Z/cp-tf-inspect
chmod +x cp-tf-inspect
./cp-tf-inspect --help
```
### 2. (Optional) Build from source

If you are on a platform other than Linux or want to build from source:
```bash
git clone https://github.com/colpal/cp-tf-inspect.git
cd cp-tf-inspect
go build -o cp-tf-inspect .
```

## 🛠 Usage

**Command:**  
`list-module-source`

- **Description:** Outputs a unique, sorted JSON array of all module sources found in the specified workspace or module directory.
- **Options:**
  - `--dir <directory>`: Path to a Terraform workspace.
- **Usage:**
  ```bash
  ./cp-tf-inspect list-module-source --dir <path-to-terraform-workspace>
  ```
- **Example:**
  ```bash
  ./cp-tf-inspect list-module-source --dir ./test/fixtures/basic
  ```
- **Output:**
  ```json
  [
    "../modules/storage",
    "git::https://example.com/bar.git",
    "terraform-aws-modules/vpc/aws"
  ]
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

We welcome contributions via **fork & pull request**!
  
### Contribution Steps

1. Fork this repository.
2. Create a feature/fix branch from `main`.
3. Make your changes, following the commit message guidelines below.
4. Push your branch and open a Pull Request (PR) to `main`.
5. Ensure all Go tests pass. Our GitHub Actions will automatically run checks and build PR preview releases.

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
