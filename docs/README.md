# GoSugar Documentation

Welcome! This documentation is an in-depth explanation of the GoSugar library. For a developer who has never seen this project before, you will find everything from start to finish.

## 📚 How to Use This Documentation?

### **1. Beginner Level (Quick Start)**
If you're seeing this project for the first time:
- Start with [`getting-started.md`](guides/getting-started.md)
- Learn what the project does and basic concepts
- Work through simple examples

### **2. Understanding the Architecture**
If you're curious about the internal structure, design decisions, and data flow:
- Read [`ARCHITECTURE.md`](architecture/ARCHITECTURE.md)
- Understand why it was designed this way
- See module relationships with diagrams

### **3. API Reference**
If you want to know how a specific function works:
- Find module-specific files in [`api/`](api/) folder
- All functions, parameters, return values for each module
- Examples and tested use cases

**Module References:**
- [`env.md`](api/env.md) - Environment variable management
- [`input.md`](api/input.md) - Reading user input
- [`validators.md`](api/validators.md) - Input validation
- [`random.md`](api/random.md) - Random data generation
- [`errors.md`](api/errors.md) - Error management
- [`file.md`](api/file.md) - File operations
- [`http.md`](api/http.md) - HTTP request operations

### **4. Guides & Learning**
To learn how to use GoSugar in specific scenarios:
- Read guides in [`guides/`](guides/) folder
- Writing CLI applications, error handling, best practices

**Available Guides:**
- [`getting-started.md`](guides/getting-started.md) - Getting started guide
- [`design-patterns.md`](guides/design-patterns.md) - Design patterns
- [`error-handling.md`](guides/error-handling.md) - Error handling
- [`testing-with-gosugar.md`](guides/testing-with-gosugar.md) - Writing tests

### **5. Design & Internal Structure**
If you want to develop the project or contribute:
- [`architecture/ARCHITECTURE.md`](architecture/ARCHITECTURE.md) - Full architecture
- [`architecture/design-decisions.md`](architecture/design-decisions.md) - Design decisions

## 🎯 Quick Navigation

| Your Question | Where to Look |
|---|---|
| "What is GoSugar?" | [`getting-started.md`](guides/getting-started.md) |
| "How does `EnvString()` work?" | [`api/env.md`](api/env.md) |
| "How do I handle errors?" | [`error-handling.md`](guides/error-handling.md) / [`api/errors.md`](api/errors.md) |
| "Can I write a CLI application?" | [`getting-started.md`](guides/getting-started.md) + [`api/input.md`](api/input.md) |
| "What's the architecture?" | [`ARCHITECTURE.md`](architecture/ARCHITECTURE.md) |
| "How do I write tests?" | [`testing-with-gosugar.md`](guides/testing-with-gosugar.md) |
| "Why use panic?" | [`design-decisions.md`](architecture/design-decisions.md) |
| "List of all functions?" | Select a module title in API reference |

## 🗂️ Documentation Structure

```
docs-en/
├── README.md (⬅️ You are here)
├── guides/
│   ├── getting-started.md          # First steps
│   ├── design-patterns.md          # Design patterns
│   ├── error-handling.md           # Error handling strategies
│   └── testing-with-gosugar.md    # Writing tests
├── api/
│   ├── env.md                      # Environment variables
│   ├── input.md                    # User input
│   ├── validators.md               # Validators
│   ├── random.md                   # Random data
│   ├── errors.md                   # Error handling
│   ├── file.md                     # File operations
│   └── http.md                     # HTTP requests
└── architecture/
    ├── ARCHITECTURE.md             # Full architecture explanation
    └── design-decisions.md         # Design decisions explained
```

## ⏱️ Estimated Reading Times

- **Quick start**: ~15 minutes (`getting-started.md`)
- **Learning one module deeply**: ~10-15 minutes (API reference)
- **Understanding architecture completely**: ~30-40 minutes (`ARCHITECTURE.md`)
- **Reading all documentation**: ~2 hours

## 💡 Important Assumptions

This documentation assumes:
- ✅ Basic knowledge of Go 1.18+ (packages, functions, generics)
- ✅ Familiarity with terminal/CLI applications
- ❌ **No prior knowledge of GoSugar is required**

## 🤝 Contributions & Questions

To improve this documentation:
1. You can point out missing topics
2. You can open issues for poorly explained sections
3. You can submit pull requests to improve documentation

## 🔗 Quick Links

- **Main Repository**: `github.com/coderianx/gosugar`
- **README.md**: Project overview (library installation, licensing)
- **ROADMAP.md**: Planned new features
- **info.md**: Old/internal documentation (for reference)

---

**Note:** The library has a modular structure. Each module works independently and you can use only what you need. Ready to start? → [`getting-started.md`](guides/getting-started.md) 🚀
