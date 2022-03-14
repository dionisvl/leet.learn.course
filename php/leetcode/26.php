<?php

class Solution
{

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function removeDuplicates(&$nums)
    {

        $underlining = '_';
        $currentKey = 0;

        while (
            isset($nums[$currentKey + 1])
            && $nums[$currentKey] !== $underlining
        ) {

            $currentValue = $nums[$currentKey];

            if ($currentValue === $nums[$currentKey + 1]) {
                unset($nums[$currentKey]);
                $nums[] = $underlining;
            }

            $currentKey++;
        }

        $nums = [...$nums];

        $k = count($nums);
        foreach ($nums as $key => $num){
            if ($num === $underlining) {
                $k = $key;
                break;
            }
        }
        return $k;
    }
}


$dataProvider = [
    1 => [
        'nums' => [1, 1, 2],
        'expectedNums' => [1, 2, '_'],
        'valid_k' => 2,
    ],
    2 => [
        'nums' => [0, 0, 1, 1, 1, 2, 2, 3, 3, 4],
        'expectedNums' => [0, 1, 2, 3, 4, '_', '_', '_', '_', '_'],
        'valid_k' => 5,
    ],
    3 => [
        'nums' => [1],
        'expectedNums' => [1],
        'valid_k' => 1,
    ],
    4 => [
        'nums' => [1,2],
        'expectedNums' => [1,2],
        'valid_k' => 2,
    ],

];

$s = new Solution();

foreach ($dataProvider as $exampleNum => $example) {
    $k = $s->removeDuplicates($example['nums']);

    if ($example['valid_k'] !== $k) {
        echo 'K = ' . $k . PHP_EOL;
        die('Error K in example #' . $exampleNum);
    }

    for ($i = 0; $i < $k; $i++) {
        if ($example['nums'][$i] !== $example['expectedNums'][$i]) {
            print_r($example['nums']);
            die('Error in example #' . $exampleNum);
        }
    }
}

echo 'all-right!';
