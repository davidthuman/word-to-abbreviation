# Word to Abbreviation

## Implementation

An HTTP Go server that returns an `index.html` file. This page contains some HTML, using Tailwind CSS through the CDN for styling, that allows a user to input a word and convert it to an abbreviation. The word is sent to the backend, which returns the associated abbreviation. Currently, no abbreviation mechanism is implemented, so the backend will just return `"Not Implemented!".

## Domains

- word2abbr.com
- word2abbreviation.com

## How to Abbreviate

I do not know of a formula to generate an abbreviation for a given word, but I will attempt to fond out.

Legal Information Institute (LII) from Cornell University Law School provides common abbreviations and omissions used in citations:

- [Months](https://www.law.cornell.edu/citation/4-600)
- [Frequently Cited Journals](https://www.law.cornell.edu/citation/4-700)
- [State Abbreviations](https://www.law.cornell.edu/citation/4-500)
- [Abbreviations for Words Used in Providing Case Histories](https://www.law.cornell.edu/citation/4-200)
- [Words Abbreviated in Case Names](https://www.law.cornell.edu/citation/4-100)

The [Abbreviation](https://en.wikipedia.org/wiki/Abbreviation) Wikipedia article is also a nice read.

Some abbreviations are created out of necessity. Text or instant message services previously has a character limit, leading to a necessity to shorten one's message as much as possible. [Conventional examples and vocabulary](https://en.wikipedia.org/wiki/SMS_language#Conventional_examples_and_vocabulary) page shows some examples of these conversions.


### Simple

Take the first 3 or 4 characters of a word. That is your abbreviation.

### First and Last

Take the first character and the last character. That is your abbreviation.

### Remove Vowels

Remove all of the vowels from the word. The remaining characters is your word.
