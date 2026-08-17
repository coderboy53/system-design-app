---
id: ood-deck-of-cards
title: Design a Deck of Cards
module: exercises/object-oriented-design
order: 40
est_minutes: 20
timelines: [medium, long]
tags: [exercise, ood, inheritance, polymorphism]
prerequisites: [interview-approach]
related: []
source: https://github.com/donnemartin/system-design-primer/blob/master/solutions/object_oriented_design/deck_of_cards/deck_of_cards.ipynb
---

# Design a Deck of Cards

> **How to use this exercise:** clarify the constraints and assumptions first, sketch the
> classes yourself, then compare against the solution below. Object-oriented design questions
> follow the same opening move as [system design questions](../../study-guide/interview-approach.md#step-1-outline-use-cases-constraints-and-assumptions)
> — establish scope before writing anything.

## Constraints and assumptions

* Is this a generic deck of cards for games like poker and black jack?
    * Yes, design a generic deck then extend it to black jack
* Can we assume the deck has 52 cards (2-10, Jack, Queen, King, Ace) and 4 suits?
    * Yes
* Can we assume inputs are valid or do we have to validate them?
    * Assume they're valid

## Solution

```python
from abc import ABCMeta, abstractmethod
from enum import Enum
import sys


class Suit(Enum):

    HEART = 0
    DIAMOND = 1
    CLUBS = 2
    SPADE = 3


class Card(metaclass=ABCMeta):

    def __init__(self, value, suit):
        self.value = value
        self.suit = suit
        self.is_available = True

    @property
    @abstractmethod
    def value(self):
        pass

    @value.setter
    @abstractmethod
    def value(self, other):
        pass


class BlackJackCard(Card):

    def __init__(self, value, suit):
        super(BlackJackCard, self).__init__(value, suit)

    def is_ace(self):
        return self._value == 1

    def is_face_card(self):
        """Jack = 11, Queen = 12, King = 13"""
        return 10 < self._value <= 13

    @property
    def value(self):
        if self.is_ace() == 1:
            return 1
        elif self.is_face_card():
            return 10
        else:
            return self._value

    @value.setter
    def value(self, new_value):
        if 1 <= new_value <= 13:
            self._value = new_value
        else:
            raise ValueError('Invalid card value: {}'.format(new_value))


class Hand(object):

    def __init__(self, cards):
        self.cards = cards

    def add_card(self, card):
        self.cards.append(card)

    def score(self):
        total_value = 0
        for card in self.cards:
            total_value += card.value
        return total_value


class BlackJackHand(Hand):

    BLACKJACK = 21

    def __init__(self, cards):
        super(BlackJackHand, self).__init__(cards)

    def score(self):
        min_over = sys.MAXSIZE
        max_under = -sys.MAXSIZE
        for score in self.possible_scores():
            if self.BLACKJACK < score < min_over:
                min_over = score
            elif max_under < score <= self.BLACKJACK:
                max_under = score
        return max_under if max_under != -sys.MAXSIZE else min_over

    def possible_scores(self):
        """Return a list of possible scores, taking Aces into account."""
        # ...


class Deck(object):

    def __init__(self, cards):
        self.cards = cards
        self.deal_index = 0

    def remaining_cards(self):
        return len(self.cards) - deal_index

    def deal_card():
        try:
            card = self.cards[self.deal_index]
            card.is_available = False
            self.deal_index += 1
        except IndexError:
            return None
        return card

    def shuffle(self):  # ...
```

## Self-check

1. How is the generic deck extended to support blackjack?
2. Which methods are abstract on the base card class, and why?
3. What would you change to support a game where card values differ per suit?
