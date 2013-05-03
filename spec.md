ini format specification
========================

Keys
----

At it's core the ini format is a collection of keys (`name=value` pairs). The
value can consist of any supported data type.

Examples:

```ini
host=`www.example.org`
ports={80, 8080}
https=false
```

Sections
--------

Keys are grouped into sections, either explicitly with a section name or
implicitly to the default section.

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

Comments
--------

Lines beginning with the character `;` are considered comments.

```ini
; This is a simple line comment.
```

Data types
----------

### Booleans

A boolean type represents the set of Boolean truth values denoted by the
predeclared constants `true` and `false`.

Examples:

```ini
debug=true
https=false
```

### Integer literals

An integer literal is a sequence of digits representing an integer constant. An
optional prefix sets a non-decimal base: `0` for octal, `0x` or `0X` for
hexadecimal. In hexadecimal literals, letters `a-f` and `A-F` represent values
10 through 15.

Examples:

```ini
; Decimal
port=8080

; Hexadecimal
; fLaC
magic=0x664C6143

; Octal
; -rw-r--r--
perm=0644
```

### Floating-point literals

A floating-point literal is a decimal representation of a floating-point
constant. It has an integer part, a decimal point, a fractional part, and an
exponent part. The integer and fractional part comprise decimal digits; the
exponent part is an `e` or `E` followed by an optionally signed decimal
exponent. One of the integer part or the fractional part may be elided; one of
the decimal point or the exponent may be elided.

Examples:

```ini
; 33%
threshold=0.33
; 4 GB.
size=4e9
```

### String literals

There are two forms of string literals: raw string literals and interpreted
string literals.

Raw string literals are character sequences between back quotes `` `` ``. Within
the quotes, any character is legal except back quote.

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
;    \303       (octal escape)
;    \xC3       (hexadecimal escape)
;    \n         (single-character escape)
;    \u00C4     (Unicode code point escape)
;    \U000000D6 (Unicode code point escape)
interpreted="\303\245\xC3\xA4\xc3\xb6\nÅ\u00C4\U000000D6"
```

[string literals section]: http://golang.org/ref/spec#String_literals

### Lists

TODO
