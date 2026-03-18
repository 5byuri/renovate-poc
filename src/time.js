'use strict';

/**
 * Returns the current UTC time as an ISO 8601 string.
 * @returns {string}
 */
function getCurrentTime() {
  return new Date().toISOString();
}

/**
 * Formats a Date object (or timestamp) as a human-readable time string
 * using the runtime's default locale.  The exact format varies by environment.
 * @param {Date|number} [date]
 * @returns {string}
 */
function formatTime(date) {
  const d = date instanceof Date ? date : new Date(date ?? Date.now());
  return d.toLocaleTimeString();
}

/**
 * Returns the elapsed milliseconds between two timestamps.
 * @param {number} start
 * @param {number} end
 * @returns {number}
 */
function elapsed(start, end) {
  return end - start;
}

module.exports = { getCurrentTime, formatTime, elapsed };
