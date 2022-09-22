<p align="center">
  <a href="" rel="noopener">
 <h3 align="center">Conventional Branch</h3></a>
</p>

## 📝 Table of Contents

- [About](#about)
- [Usage](#usage)
- [Built Using](#built_using)


## 🧐 About <a name = "about"></a>
A consistent branch naming convention is part of code review best practices, and can make life much more easier for anyone who's collaborating and reviewing your code, in addition to using static analysis tools.


## 🎈 Usage <a name="usage"></a>
The branch name should be structured as follows:
```html
<type>/<short_description>

```

### recommendations 

type     | description | 
:------: | :---------- |
docs     |	Updating documentation
feat     |  Adding new functionality
fix      |	Error correction
bug      |	Bug correction
refactor |	Code edits without bug fixes or adding new features

### Example

```bash
docs/update-readme

feat/add-Polish-language

fix/prevent-racing-of-requests
```