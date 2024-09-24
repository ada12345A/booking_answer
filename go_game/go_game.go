package go_game

/**
 * Consider an N*N game board, with a black and white pieces that can be placed on it. You are given a board with placed pieces around it in a random spots.
 *
 * You need to implement a function that determines if a piece (black or white) is captured on a given coordination (x, y).
 *
 * A piece is defined as captured by the following rules:
 * 1. If all sides (up, down, left &amp; right) contains an opposite piece that surrounds/blocking it.
 * 2. If some of the sides are blocked (for example, right and down) and the other ones arae out of bound (OOB defined for coords: x: &lt;= 0, y: &lt;= 0) it's considered as blocked.
 * 3. If one of the sides is empty, it's free.
 * 4. If one of the sides contains the same piece type, and that piece is not captured (by the rules above), it's free.
 * 5. Note that pieces may be captured in a clustered way (related to #4).
 *
 * For example, consider the following coordinates:
 * coord(1,1) = B
 * coord(1,2) = W
 * coord(2,1) = W
 */

// https://github.com/zhaohuabing/booking-interview-practice/blob/master/src/com/zhaohuabing/GoGame.java
