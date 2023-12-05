# Advent of Code Day 1 
## Part One
You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

```
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
```

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

### Test Cases
- Case 1: There are two numbers in a line. The first number is the tens digit and the second number is the ones digit.
- Case 2: There is only one number in a line. That number is both the tens and ones digit.
- Case 3: There are multiple numbers and we only want the first and last.

Answer: 54697

## Part Two

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

```
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
```

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

### Solution
First, I tried building a buffer of characters, reading from left to right. I was already doing the same thing for numbers, so adding the number-words seemed like a simple first step. The problem with this implementation is the garbage data in each line. I would have had to use two pointers and loop through each line twice to eliminate the bad data.

Next, I tried using a regex with all the possible values "or'd". But this only teturned the first instance of each match and didn't account for overlapping words.

Finally, I used the Index and LastIndex functions from the strings package, looping through the possible tokens (len = 18) for each line (1000). I added separate counters to track the first and last index, updating the first/last numbers when it found a better match.

Answer: 54885
