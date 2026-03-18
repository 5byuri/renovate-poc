'use strict';

const { test } = require('node:test');
const assert = require('node:assert/strict');
const { getCurrentTime, formatTime, elapsed } = require('./time');

test('getCurrentTime returns a valid ISO 8601 string', () => {
  const result = getCurrentTime();
  assert.match(result, /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z$/);
});

test('formatTime returns a non-empty string for a given Date', () => {
  const date = new Date(2024, 0, 1, 12, 0, 0);
  const result = formatTime(date);
  assert.ok(typeof result === 'string' && result.length > 0);
});

test('formatTime accepts a numeric timestamp', () => {
  const ts = new Date(2024, 5, 15, 8, 30, 0).getTime();
  const result = formatTime(ts);
  assert.ok(typeof result === 'string' && result.length > 0);
});

test('elapsed calculates the difference between two timestamps', () => {
  assert.equal(elapsed(1000, 3000), 2000);
  assert.equal(elapsed(500, 500), 0);
  assert.equal(elapsed(2000, 1000), -1000);
});
