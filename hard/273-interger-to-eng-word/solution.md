first, gathering the possible word

| Digit | Possible Words |
|-------|------------------------------------------------------|
| 0     | zero |
| 1     | one, ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen |
| 2     | two, twenty |
| 3     | three, thirty  |
| 4     | four, forty  |
| 5     | five, fifty  |
| 6     | six, sixty  |
| 7     | seven, seventy  |
| 8     | eight, eighty  |
| 9     | nine, ninety  |

hundred, thousand
and more... Nonillion
Octillion
Septillion
Sextillion
Quintillion
Quadrillion
Trillion
Billion
Million
(ty chatgpt lol)
=================
example
the original input -> 1234567890123456789012345678901

1234567890123456789012345678901
slice3fromEnd()
1 234 567 890 123 456 789 012 345 678 901

1 234 567 890 123 456 789 012 345 678 901
addUnitToSlices()
1   Nonillion
234 Octillion
567 Septillion
890 Sextillion
123 Quintillion
456 Quadrillion
789 Trillion
012 Billion
345 Million
678 Thousand
901

1 Nonillion
234 Octillion
567 Septillion
890 Sextillion
123 Quintillion
456 Quadrillion
789 Trillion
012 Billion
345 Million
678 Thousand
901
addWordsInLessHundred()
the output here i think i must be 1 noni... two hundred thirty four .... and more

i think with this chain, i can solve many cases haha
slice3fromEnd().addUnitToSlices().addWordsInLessHundred()
btw could you suggest me better name for these func