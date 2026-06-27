// SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>
//
// SPDX-License-Identifier: MIT

//! Minimal fixture library for testing the rust action.

/// Returns the sum of `a` and `b`.
pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn add_works() {
        assert_eq!(add(2, 3), 5);
    }
}
