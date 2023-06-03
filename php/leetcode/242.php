<?php

declare(strict_types=1);

/**
 * 242. Valid Anagram
 * easy
 *
 * Given two strings s and t, return true if t is an anagram of s, and false otherwise.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
 * typically using all the original letters exactly once.
 */
class Solution
{

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    function isAnagram($s, $t)
    {
        if (
            $s === ''
            || strlen($s) !== strlen($t)
        ) {
            return false;
        }
        $letters1 = array_fill_keys(range('a', 'z'), 0);
        $letters2 = $letters1;
        foreach (str_split($s) as $char1) {
            $letters1[$char1]++;
        }
        foreach (str_split($t) as $char2) {
            $letters2[$char2]++;
        }
        return $letters1 === $letters2;
    }
}

$dataProvider = [
    7 => [
        's' => 'aacc',
        't' => "ccac",
        'result' => false,
    ],
    1 => [
        's' => 'anagram',
        't' => "nagaram",
        'result' => true,
    ],
    2 => [
        's' => 'sdf',
        't' => "nagaram",
        'result' => false,
    ],
    3 => [
        's' => '',
        't' => "nagaram",
        'result' => false,
    ],
    4 => [
        's' => 'sdfs',
        't' => "",
        'result' => false,
    ],
    5 => [
        's' => 'a',
        't' => "a",
        'result' => true,
    ],
    6 => [
        's' => 'rat',
        't' => "car",
        'result' => false,
    ],
];

$s = new Solution();

foreach ($dataProvider as $exampleNumber => $example) {
    $result = $s->isAnagram($example['s'], $example['t']);

    if ($result !== $example['result']) {
        echo 'EXAMPLE: ' . json_encode($example, JSON_THROW_ON_ERROR) . PHP_EOL;
        echo 'RESULT: ' . json_encode($result, JSON_THROW_ON_ERROR) . PHP_EOL;
        die('Error in example #' . $exampleNumber);
    }
}

echo 'all-right!';
