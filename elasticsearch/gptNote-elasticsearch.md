# Difference Between Keyword and Text in ES 

Yes, **Elasticsearch** supports **fuzzy search** even for fields defined as type `text`. Fuzzy search allows Elasticsearch to handle typographical errors or approximate matches in queries, which is particularly useful when users don't type the exact word correctly. 

### How Fuzzy Search Works in Elasticsearch:

- **Fuzzy search** in Elasticsearch is designed to match documents that are similar to the query string but may have spelling mistakes, typos, or other small variations.
- It operates by adjusting the **Levenshtein distance** (the number of single-character edits—insertions, deletions, or substitutions—needed to change one word into another).

### Applying Fuzzy Search in Elasticsearch Queries:

To use **fuzzy search** in Elasticsearch, you can include the `fuzziness` parameter in your search query. The `fuzziness` parameter controls how many character changes are allowed between the query term and the indexed terms.

#### Example: Fuzzy Search on a `text` Field

Let's say we have documents where the `name` field is of type `text`, and we want to allow fuzzy search for names, even if the query is slightly misspelled.

#### Example Document Mapping:
```json
{
  "mappings": {
    "properties": {
      "name": {
        "type": "text"
      }
    }
  }
}
```

#### Fuzzy Query Example:
If you want to perform a fuzzy search for the name "Jon", but the user misspells it as "Jno", you can use the following query:

```json
{
  "query": {
    "match": {
      "name": {
        "query": "Jno",
        "fuzziness": "AUTO"
      }
    }
  }
}
```

- **`fuzziness: "AUTO"`**: This automatically adjusts the fuzziness based on the length of the query term. It allows up to 1 or 2 character changes depending on the length of the search term.
  
- This query will match documents with names like "Jon", "John", "Jno", etc., even though the input was slightly misspelled.

### Fuzziness Levels:
- **AUTO**: Automatically selects the appropriate level of fuzziness.
  - For terms of 1-2 characters, no fuzziness is applied.
  - For terms of 3-5 characters, one edit is allowed.
  - For terms of 6 or more characters, two edits are allowed.
- **Explicit Fuzziness Levels**:
  - **0**: Exact match.
  - **1**: One character difference allowed.
  - **2**: Two character differences allowed.

### Full Example:
#### Indexing Data:
Let's index a few documents:
```json
POST /people/_doc/1
{
  "name": "Jon"
}

POST /people/_doc/2
{
  "name": "John"
}

POST /people/_doc/3
{
  "name": "Johnny"
}
```

#### Running a Fuzzy Search:
```json
GET /people/_search
{
  "query": {
    "match": {
      "name": {
        "query": "Jno",    // Misspelled query
        "fuzziness": "AUTO" // Fuzzy matching
      }
    }
  }
}
```

This query would match documents with names similar to "Jon", "John", or other close variations, despite the typo in the query.

### Limitations of Fuzzy Search:
- **Performance**: Fuzzy search can be slower compared to exact matches, especially for large datasets.
- **Character-based matching**: Fuzzy search works well for character-based typos but might not handle more complex linguistic issues (e.g., stemming or synonyms).

### Conclusion:
Yes, Elasticsearch supports fuzzy search on fields of type `text`. You can use the `fuzziness` parameter in queries to handle typos or approximate matches.

# FTS - Does full-text-search Support Fuzzy Search
**Full-Text Search (FTS)** refers to the ability to search for words or phrases within the entire body of text across documents. It is designed to allow users to perform more flexible and intuitive searches on large volumes of text data by supporting various advanced search techniques, including:

1. **Tokenization**: Breaking the text into smaller units (tokens), such as words or phrases.
2. **Stemming**: Reducing words to their base or root form (e.g., "running" → "run").
3. **Stopwords**: Ignoring common words like "the" and "is" in search queries.
4. **Ranking and Scoring**: Returning search results ranked by relevance rather than strict exact matches.

### Does Full-Text Search Support Text Search?

Yes, **full-text search (FTS)** is primarily used for **text search**. It is built specifically to handle **natural language queries** on text fields, where simple keyword or exact match queries are insufficient. With FTS, the system searches through documents to find the best match based on the context, synonyms, word relevance, and even proximity of terms.

### Does Full-Text Search Support Fuzzy Search?

Yes, **full-text search** typically supports **fuzzy search**. Fuzzy search allows for finding results that are not exact matches to the query, helping handle:

- **Misspellings** or **typos**.
- **Approximate matches** based on character changes (like transposed letters).
- **Phonetic matching** (in some advanced FTS systems).

In systems like **Elasticsearch** and **PostgreSQL** (which both support FTS), fuzzy search can be enabled or configured to allow a certain number of differences between the search query and the stored text, such as adding support for words with similar spelling or slight differences.

### Full-Text Search Features:

1. **Text Search**:
   - **Tokenization**: The text is broken into words (tokens) that are indexed for fast lookups.
   - **Stemming**: Words are reduced to their root form, so searching for "running" will match "run".
   - **Stop Words**: Common words like "and," "the," "is" may be ignored to focus on more meaningful terms.
   - **Ranking**: Results are scored based on how relevant they are to the query, often using algorithms such as TF-IDF (Term Frequency-Inverse Document Frequency).

2. **Fuzzy Search**:
   - Full-text search often supports **fuzzy matching**, allowing approximate matches. Fuzzy search can be helpful for handling user input that may contain:
     - Typos: Searching for "Jonh" would still return results for "John."
     - Similar spellings: The system can return close matches for terms that are spelled similarly.
     - Configurable **Levenshtein distance** (number of character edits): You can allow the system to match words with one or two character differences from the query.

### Example in Elasticsearch:

In Elasticsearch, full-text search (`match` query) can be combined with **fuzzy search** using the `fuzziness` parameter. Here's how it works:

```json
{
  "query": {
    "match": {
      "name": {
        "query": "Jonh",  // Misspelled "John"
        "fuzziness": "AUTO" // Fuzzy search to allow for minor spelling differences
      }
    }
  }
}
```

This query will search the `name` field, allowing fuzzy matches for words that are similar to "Jonh" (e.g., "John").

### Summary:
- **Full-Text Search (FTS)** is a powerful technique for searching through large text data, supporting features like tokenization, stemming, and ranking.
- **Text Search** is the core purpose of FTS, allowing searches that can handle variations in the text.
- **Fuzzy Search** is typically supported by FTS systems, enabling approximate matches even when the search query contains typos or near-matches to the text in the database.
