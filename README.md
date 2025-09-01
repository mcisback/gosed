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
history | gx -d 5 # remove line number 5 (start counting from 1)
```
Delete matching lines:
```bash
echo file.txt | gx 'Sublime/d' # remove all lines containing sublime
```

Use instead of tr:
```bash
echo $PATH | gx ':' '\n'
```