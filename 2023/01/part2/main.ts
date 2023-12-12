import * as fs from 'fs'

const MATCH_POSITION = 1
const DIGITS_TO_NUMBERS: Record<string, string> = {
  'one': '1', '1': '1',
  'two': '2', '2': '2',
  'three': '3', '3': '3',
  'four': '4', '4': '4',
  'five': '5', '5': '5',
  'six': '6', '6': '6',
  'seven': '7', '7': '7',
  'eight': '8', '8': '8',
  'nine': '9', '9': '9'
}

const digitsRegex = Object.keys(DIGITS_TO_NUMBERS).flatMap((key) => key + '|').join('').slice(0, -1)
const globalRegex = new RegExp(`(?=(${digitsRegex}))`, 'g')

const file = fs.readFileSync('input.txt', 'utf-8')
const finelValue = file.toString().replace('\r\n', '\n').split('\n')
    .map((line) => {
      let result = 0;
      const matches = [...line.matchAll(globalRegex)]

      if (matches.length === 1) result = parseInt(DIGITS_TO_NUMBERS[matches[0][MATCH_POSITION]] + DIGITS_TO_NUMBERS[matches[0][MATCH_POSITION]])
      else if (matches.length !== 0) result = parseInt(DIGITS_TO_NUMBERS[matches[0][MATCH_POSITION]] + DIGITS_TO_NUMBERS[matches[matches.length - 1][MATCH_POSITION]])

      return result
    })
    .reduce((prev, current) => prev + current, 0);

console.log(`Result is: ${finelValue}`)