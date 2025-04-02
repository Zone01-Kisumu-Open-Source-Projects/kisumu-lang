# Kisumu Lang Security Policy

## Supported Versions
| Version | Supported          |
|---------|--------------------|
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting Vulnerabilities
**Please do not report security issues publicly**
❗ **Use GitHub's private reporting**  [Security tab](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/security). Or use methods below. 
<!-- We should register domain and use this for sec: security@kisumu-lang.org -->
1. **Email:** kh3rld@duck.com (PGP Key [0xABCD1234])
2. **Response Time:** We acknowledge reports within 48 hours
3. **Disclosure:** Patched vulnerabilities will be disclosed via GitHub advisories

## Security Practices

### For Developers
- All code must pass:
  ```bash
  make lint # Includes gosec and static analysis
  ```
- Banned patterns:
  - Unsafe pointer arithmetic (`unsafe` package)
  - Shell command injection (e.g., unfiltered `exec.Command` input)

### For Contributors
1. **Dependencies:**
   - Audit third-party packages with:
     ```bash
     go mod audit
     ```
   - Report suspicious dependencies immediately

2. **Code Reviews:**
   - All PRs require security review from maintainers
   - Security-critical changes get 72-hour review period

### Threat Model
| Threat | Mitigation |
|--------|------------|
| Code injection | Input validation + AST sanitization |
| Memory corruption | Bounds checking + no pointer arithmetic |
| Dependency attacks | Pinned versions + automated audits |
