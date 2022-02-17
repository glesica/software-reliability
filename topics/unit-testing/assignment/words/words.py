from typing import List


def split_words(sentence: str) -> List[str]:
    """
    Split a string into a list of words. A word boundary is defined to be any
    number of space or tab characters.
    """
    words = []
    next_word = ""

    for character in sentence:
        if character == " ":
            if next_word != "":
                words.append(next_word)
                next_word = ""
        else:
            next_word = next_word + character

    if next_word != "":
        words.append(next_word)

    return words


def join_words(words: List[str]) -> str:
    """
    Take a list of words and return it as a sentence, with adjacent words
    separated by a single space character.
    """
    sentence = ""

    for word in words:
        if sentence != "":
            sentence = sentence + " "
        sentence = sentence + word

    return sentence


def count_words(sentence: str) -> int:
    """
    Count the number of words in the sentence. A word boundary is defined to be
    any number of space or tab characters.
    """
    return len(split_words(sentence))
