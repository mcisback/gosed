# GOSED

## Fastest, easiest alternative to sed/grep ecc

Install
```bash
git clone https://github.com/mcisback/gosed.git
cd gosed
go build -o /usr/local/bin/gx . # System install
go build -o ~/.local/bin/gx . # Local user install or use your custom path
```

Example match:
```bash
history | gx '^[\d\s]+rclone'
```

Example replace:
```bash
history | gx '^[\d\s]+' 'replacement'
```

Remove nth line:
```bash
history | gx '5/d' # remove line number 5 (start counting from 1)
```

Remove nth:mth range lines:
```bash
history | gx '5:10/d' # Remove lines from number 5 to 10 (start counting from 1)
```

Delete matching lines:
```bash
echo file.txt | gx 'Sublime/d' # Remove all lines containing sublime
```

Show nth line:
```bash
history | gx '5/m' # Show line number 5 (start counting from 1)
```

Show nth:mth range lines:
```bash
history | gx '5:10/m' # Show lines from number 5 to 10 (start counting from 1)
```

Use instead of tr:
```bash
echo $PATH | gx ':' '\n'
```

Highlight every line containing bash:
```bash
echo README.md | gx 'bash/b'
```

Load from file instead of STDIN:
```bash
gx -f README.md 'bash/b'
```