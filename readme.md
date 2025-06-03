# cp-tf-inspect

[![CI & Release](https://github.com/colpal/cp-tf-inspect/actions/workflows/main.yaml/badge.svg)](https://github.com/colpal/cp-tf-inspect/actions/workflows/main.yaml)

A fast and extensible command-line tool for inspecting Terraform workspaces.

---

## ✨ Features

- **List unique module sources** in a given Terraform workspace
- JSON output for easy pipe/sharing
- Extensible architecture—add new commands easily!
- Cross-platform (builds/releases available for Linux)
- Automated testing and release workflows via GitHub Actions

---

## 🚀 Quickstart

### 1. Download the latest release

Download the Linux binary from the [Releases page](https://github.com/colpal/cp-tf-inspect/releases/latest), or build from source as below.

### 2. Build from source

```bash
git clone https://github.com/colpal/cp-tf-inspect.git
cd cp-tf-inspect
go build -o cp-tf-inspect
```

### 3. Try it out

```bash
./cp-tf-inspect list-module-source --dir ./test/fixtures/basic
```

**Sample Output:**
```json
[
  "../modules/storage",
  "git::https://example.com/bar.git",
  "terraform-aws-modules/vpc/aws"
]
```

---

## ⚙️ Usage

```bash
cp-tf-inspect list-module-source --dir <path-to-terraform-workspace>
```

- `--dir` (required): Path to the root directory of your Terraform workspace.

---

## 🧪 Testing

**Unit tests:**
```bash
go test ./cmd/...
```

**Integration/CLI tests:**
- Make sure the compiled binary (`cp-tf-inspect`) is present in the project root:
```bash
go build -o cp-tf-inspect
go test ./test/...
```

---

## 🌱 Extending

To add a new subcommand:
1. Add a new `<command>.go` in the `cmd/` folder.
2. Register it in `cmd/root.go`.
3. Write unit tests in `cmd/<command>_test.go`.
4. [Optional] Add integration tests in `test/`.

---

### Branching Model

This project uses a **trunk-based workflow**:
- All production code lives on the `main` branch.
- Feature/fix branches → PR → review → merge to `main`.
- Releases are automated via GitHub Actions on merges to `main`.

---

## 📦 CI/CD

- [GitHub Actions](https://github.com/colpal/cp-tf-inspect/actions) runs all tests, builds the Linux binary, and attaches it to every release and PR for preview.
- The latest production binary is always available in the [Releases page](https://github.com/colpal/cp-tf-inspect/releases).

