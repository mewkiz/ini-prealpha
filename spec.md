ini format specification
========================

Tokens
------

Tokens form the vocabulary of the ini format. While breaking the input into
tokens, the next token is the longest sequence of characters that form a valid
token.

Keys
----

At it's core the ini format is a collection of keys (`name=value` pairs). The
value can consist of any supported data type.

A key name is a sequence of one or more letters and digits. The first character
in a key name must be a letter. See the [identifiers section][] of the Go
specification for a detailed definition of valid key names.

Examples:

```ini
host=`www.example.org`
ports={80, 8080}
https=false
ratio=0.25
```

[identifiers section]: http://golang.org/ref/spec#Identifiers


Sections
--------

Keys are grouped into sections, either explicitly with a section name or
implicitly to the default section.

A section name is a sequence of one or more characters prefixed with a left
bracket `[` and suffixed with a right bracket `]`. The character sequence may
not include new lines `\n` but may include brackets `[`, `]` and semicolons `;`.

The section name of the default section is always an empty string `""`.

Only white space and new line characters are allowed after a section name. Thus
comments are explicitly prohibited after a section name.

Examples:

```ini
; Default section:
contact=`info@example.org`

; Section "prod":
[prod]
url=`www.example.com`
debug=false
https=true

; Section "dev":
[dev]
url=`dev.example.internal`
debug=true
https=false
```

```ini
; Check for change:
[http://www.example.org/?q=test;ing]

[http://www.example.org/?id[]=1&id[]=2]
interval="1m30s"
```

Comments
--------

Line comments start with a semicolon `;` and stop at the end of the line. A line
comment acts like a newline.

A line comment may be used anywhere a new line could be. The only exception is
after section names, where no line comments are allowed.

Examples:

```ini
; This is a simple line comment.

[Zombie]
name="bgen" ; Behavior generator
args={
        2, ; Speed
        1, ; Intelligence
        9, ; Health
     }
```

Invalid examples:

```ini
[section] ; this is an invalid comment.

Data types
----------

### Booleans

Booleans are represented using the predeclared constants `true` and `false`.

Examples:

```ini
debug=true
https=false
```

### Integer literals

An integer literal is an optionally signed sequence of digits representing an
integer constant. An optional prefix sets a non-decimal base: `0` for octal,
`0x` or `0X` for hexadecimal. In hexadecimal literals, letters `a-f` and `A-F`
represent values 10 through 15.

Examples:

```ini
; Decimal representation.
port=8080

; Hexadecimal representation.
;    fLaC
magic=0x664C6143

; Octal representation.
;    -rw-r--r--
perm=0644

; Signed integer literal.
step=-3
```

### Floating-point literals

A floating-point literal is an optionally signed decimal representation of a
floating-point constant. It has an integer part, a decimal point, a fractional
part, and an exponent part. The integer and fractional part comprise decimal
digits; the exponent part is an `e` or `E` followed by an optionally signed
decimal exponent. One of the integer part or the fractional part may be elided;
one of the decimal point or the exponent may be elided.

Examples:

```ini
; 33%
threshold=0.33

; 4 GB.
size=4e9

; Signed floating-point literal.
delta=-0.03
```

### String literals

There are two forms of string literals: raw string literals and interpreted
string literals.

Raw string literals are character sequences between back quotes ``` `` ```.
Within the quotes, any character is legal except back quote.

Interpreted string literals are character sequences between double quotes `""`.
The text between the quotes, which may not contain newlines, forms the value of
the literal, with backslash escapes being interpreted.

See the [string literals section][] of the Go specification for a detailed
definition.

Examples:

```ini
[1984]
; raw string literals may contain double quotes.
author=`Eric Arthur Blair (pseudonym "George Orwell")`

[Swedish]
; raw string literals may contain new lines.
raw=`åäö
ÅÄÖ`

; interpreted string literals include many different escapes.
;    \n         (single-character escape)
;    \303       (octal escape)
;    \xC3       (hexadecimal escape)
;    \u00C4     (Unicode code point escape)
;    \U000000D6 (Unicode code point escape)
interpreted="\303\245\xC3\xA4\xc3\xb6\nÅ\u00C4\U000000D6"
```

[string literals section]: http://golang.org/ref/spec#String_literals

### Lists

TODO
