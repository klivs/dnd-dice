package dnd_dice

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// DiceExpression accept a common dnd dice expression as string like "3d6", "2d20"
// It returns an array with the results and any error encountered.
func DiceExpression(diceExpression string) ([]int, error) {
	splitExpression := strings.Split(diceExpression, "d")

	rollTimes, err := strconv.Atoi(splitExpression[0])
	if err != nil {
		return nil, err
	}

	diceSize, err := strconv.Atoi(splitExpression[1])
	if err != nil {
		return nil, err
	}

	return Dice(rollTimes, diceSize)
}

// Dice accepts the roll times and the dice size as integers
// It returns an array with the results and any error encountered.
func Dice(rollTimes int, diceSize int) ([]int, error) {
	if diceSize <= 0 {
		return nil, errors.New("negative or zero Dice size")
	}
	if rollTimes <= 0 {
		return nil, errors.New("negative or zero roll times")
	}
	results := make([]int, 0)
	for i := 0; i < rollTimes; i++ {
		results = append(results, rand.Intn(diceSize)+1)
	}
	return results, nil
}

// D4 accepts the roll times of a d4 dice
// It returns an array with the results and any error encountered.
func D4(rollTimes int) ([]int, error) {
	return Dice(rollTimes, 4)
}

// D6 accepts the roll times of a d6 dice
// It returns an array with the results and any error encountered.

func D6(rollTimes int) ([]int, error) {
	return Dice(rollTimes, 6)
}

// D8 accepts the roll times of a d8 dice
// It returns an array with the results and any error encountered.

func D8(rollTimes int) ([]int, error) {
	return Dice(rollTimes, 8)
}

// D10 accepts the roll times of a d10 dice
// It returns an array with the results and any error encountered.
func D10(rollTimes int) ([]int, error) {
	return Dice(rollTimes, 10)
}

// D12 accepts the roll times of a d12 dice
// It returns an array with the results and any error encountered.
func D12(rollTimes int) ([]int, error) {
	return Dice(rollTimes, 12)
}

// D20 accepts the roll times of a d20 dice
// It returns an array with the results and any error encountered.
func D20(rollTimes int) ([]int, error) {
	return Dice(rollTimes, 20)
}

// D100 accepts the roll times of a d100 dice
// It returns an array with the results and any error encountered.
func D100(rollTimes int) ([]int, error) {
	return Dice(rollTimes, 100)
}
