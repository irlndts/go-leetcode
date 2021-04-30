package problems

// https://leetcode.com/problems/rank-scores/

/*
SELECT
    Score as Score,
    DENSE_RANK() OVER(ORDER BY Score DESC) 'Rank'
FROM Scores;
*/
