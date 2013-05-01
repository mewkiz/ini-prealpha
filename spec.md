ini format specification
========================

Keys
----

At it's core the ini format is a collection of keys (`name=value` pairs). The
value can consist of any supported data type.

Examples:

```ini
host=`www.example.org`
ports=[80, 8080]
https=false
timeout=30s
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

```ini
; 33%
threshold=0.33
; 4 GB.
size=4e9
```

### String literals

A string literal represents a string constant obtained from concatenating a
sequence of characters. There are two forms: raw string literals and interpreted
string literals.

Raw string literals are character sequences between back quotes ` `` `. Within
the quotes, any character is legal except back quote. The value of a raw string
literal is the string composed of the uninterpreted (implicitly UTF-8-encoded)
characters between the quotes; in particular, backslashes have no special meaning
and the string may contain newlines. Carriage returns inside raw string literals
are discarded from the raw string value.

Interpreted string literals are character sequences between double quotes `""`.
The text between the quotes, which may not contain newlines, forms the value of
the literal, with backslash escapes interpreted as described below. The
three-digit octal (`\nnn`) and two-digit hexadecimal (`\xnn`) escapes represent
individual bytes of the resulting string; all other escapes represent the
(possibly multi-byte) UTF-8 encoding of individual characters. Thus inside a
string literal `\377` and `\xFF` represent a single byte of value `0xFF`=255,
while `ÿ`, `\u00FF`, `\U000000FF` and `\xc3\xbf` represent the two bytes `0xc3`
`0xbf` of the UTF-8 encoding of character U+00FF.

Several backslash escapes allow arbitrary values to be encoded as ASCII text.
There are four ways to represent the integer value as a numeric constant: `\x`
followed by exactly two hexadecimal digits; `\u` followed by exactly four
hexadecimal digits; `\U` followed by exactly eight hexadecimal digits, and a
plain backslash `\` followed by exactly three octal digits. In each case the
value of the literal is the value represented by the digits in the corresponding
base.

Although these representations all result in an integer, they have different
valid ranges. Octal escapes must represent a value between 0 and 255 inclusive.
Hexadecimal escapes satisfy this condition by construction. The escapes `\u` and
`\U` represent Unicode code points so within them some values are illegal, in
particular those above `0x10FFFF` and surrogate halves.

After a backslash, certain single-character escapes represent special values:

	\a   U+0007 alert or bell
	\b   U+0008 backspace
	\f   U+000C form feed
	\n   U+000A line feed or newline
	\r   U+000D carriage return
	\t   U+0009 horizontal tab
	\v   U+000b vertical tab
	\\   U+005c backslash
	\"   U+0022 double quote

All other sequences starting with a backslash are illegal inside string
literals.

Examples:

```ini
[1984]
author=`Eric Arthur Blair (pseudonym "George Orwell")`

[Swedish]
raw=`åäöÅÄÖ`
interpreted="\xC3\xA5\xC3\xA4\xC3\xB6\u00C5\u00C4\u00D6"
```

### Duration

A duration is a possibly signed sequence of decimal numbers, each with optional
fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time
units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

Examples:

```ini
timeout=1m15s
```

### Lists

TODO
