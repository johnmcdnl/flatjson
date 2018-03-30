# FlatJSON

Transform a JSON string to a map of key value pairs

```json
{
  "glossary": {
    "title": "example glossary",
    "GlossDiv": {
      "title": "S",
      "GlossList": {
        "GlossEntry": {
          "ID": "SGML",
          "SortAs": "SGML",
          "GlossTerm": "Standard Generalized Markup Language",
          "Acronym": "SGML",
          "Abbrev": "ISO 8879:1986",
          "GlossDef": {
            "para":
              "A meta-markup language, used to create markup languages such as DocBook.",
            "GlossSeeAlso": ["GML", "XML"]
          },
          "GlossSee": "markup"
        }
      }
    }
  }
}
```

``` text
glossary.title                                                    example glossary
glossary.GlossDiv.title                                          S
glossary.GlossDiv.GlossList.GlossEntry.GlossDef.para             A meta-markup language, used to create markup languages such as DocBook.
glossary.GlossDiv.GlossList.GlossEntry.Acronym                   SGML
glossary.GlossDiv.GlossList.GlossEntry.Abbrev                    ISO 8879:1986
glossary.GlossDiv.GlossList.GlossEntry.GlossSee                  markup
glossary.GlossDiv.GlossList.GlossEntry.ID                        SGML
glossary.GlossDiv.GlossList.GlossEntry.SortAs                    SGML
glossary.GlossDiv.GlossList.GlossEntry.GlossTerm                 Standard Generalized Markup Language
glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso.#   2
glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso.0   GML
glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso.1   XML

```
