const test = require("node:test");
const assert = require("node:assert");

const { add, sub } = require("./calc");

test("add(3, 4) equals 7", () => {
    assert.strictEqual(add(3, 4), 7);
});

test("sub(4, 3) equals 1", () => {
    assert.strictEqual(sub(4, 3), 1);
});

test("sub(3, 4) equals -1", () => {
    assert.strictEqual(sub(3, 4), -1);
});
