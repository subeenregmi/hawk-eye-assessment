# hawk-eye-assessment

## depenedencies

- go 1.24.9
- make

## setup

to run:
```sh
make
```

to test:
```sh
make test
```

## how it works?

1. draw a random card from a shuffled deck
2. check the user input if its 'H' (higher) or 'L' (lower)
3. draw another next card and check if the user input is correct
4. repeat

## assumptions

- game ends in eight cards
- there are only two jokers
- everytime a card is drawn that card is then removed from the deck

## next setup

- introduce an endless Higher or Lower mode with a three-strike limit, where players earn points for each correct prediction
