<p align="center">
  <a href="https://www.conventionalcommits.org/en/v1.0.0/" rel="noopener">
 <h3 align="center">Conventional Commits</h3></a>
</p>

## 📝 Table of Contents

- [About](#about)
- [Installing](#installing)
- [Usage](#usage)
- [Built Using](#built_using)

## 🧐 About <a name = "about"></a>

The Conventional Commits specification is a lightweight convention on top of commit messages. It provides an easy set of rules for creating an explicit commit history; which makes it easier to write automated tools on top of. This convention dovetails with SemVer, by describing the features, fixes, and breaking changes made in commit messages.

## Installing <a name = "installing"></a>

### Pre-commit

Install pre-commit

https://pre-commit.com/#install

If your usage `zsh`. Try adding this command to your .zshrc.
```[[ -e ~/.profile ]] && emulate sh -c 'source ~/.profile'```

```
pre-commit install --hook-type commit-msg
```

## 🎈 Usage <a name="usage"></a>
The commit message should be structured as follows:
```html
<type>(<scope>): <subject>

<body>

<footer>
```

### recommendations 

type     | description | 
:------: | :---------- |
build    |	Building a project or changing external dependencies
ci       |	Setting up CI and working with scripts
docs     |	Updating documentation
feat     |  Adding new functionality
fix      |	Error correction
perf     |	Changes aimed at improving performance
refactor |	Code edits without bug fixes or adding new features
revert   |	Rollback to previous commits
style    |	Code style edits (tabs, indents, dots, commas, etc.)
test     |	Adding Tests

### Example

```bash
docs: update readme

feat(lang): add Polish language

fix: prevent racing of requests
```