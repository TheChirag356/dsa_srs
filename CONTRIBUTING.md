# Contributing to dsa\_srs

Thank you for your interest in contributing to **dsa\_srs**! We welcome contributions from everyone—whether it’s a bug report, feature request, or code improvements.

---

## 📋 Getting Started

1. **Fork the repository** on GitHub:
   [https://github.com/TheChirag356/dsa\_srs/fork](https://github.com/TheChirag356/dsa_srs/fork)

2. **Clone your fork** locally:

   ```bash
   git clone https://github.com/<your-username>/dsa_srs.git
   cd dsa_srs
   ```

3. **Create a new branch** for your feature or bugfix:

   ```bash
   git checkout -b feature/my-new-feature
   ```

4. **Install dependencies** (if any) and build:

   ```bash
   go mod tidy
   go build
   ```

5. **Make your changes**. Aim for clean, idiomatic Go and follow existing code patterns. Please include tests where appropriate.

6. **Run tests** to ensure nothing breaks:

   ```bash
   go test ./...
   ```

7. **Commit your changes** with a clear message:

   ```bash
   git add .
   git commit -m "Add feature X for reviewing cards by tag"
   ```

8. **Push your branch** to your fork:

   ```bash
   git push origin feature/my-new-feature
   ```

9. **Open a Pull Request** against the `main` branch of the upstream repository.

---

## 🛠️ Code Guidelines

* **Follow Go conventions** (use `gofmt`, `go vet`).
* **Write clear, descriptive commit messages**.
* **Add documentation** for new public functions and types.
* **Write tests** for bug fixes and new features when possible.
* **Keep pull requests small** and focused on a single change.

---

## 🐛 Reporting Bugs

If you encounter any bugs or unexpected behavior:

1. Search existing [issues](https://github.com/TheChirag356/dsa_srs/issues) to see if it’s already reported.
2. If not, **open a new issue** with:

   * A descriptive title and detailed steps to reproduce.
   * Go version and OS/architecture.
   * Relevant logs or screenshots.

---

## 💡 Suggesting Enhancements

Feature requests and ideas are welcome! Please open an issue with:

* A clear description of the feature.
* Motivation and use cases.
* Possible implementation approach (if you have one).

---

## 🎉 Contributors

We appreciate all contributions, big and small. All contributors will be acknowledged in the project’s README and contributors list.

Thank you for making **dsa\_srs** better! 🚀
