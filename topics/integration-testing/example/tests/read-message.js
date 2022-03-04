const puppeteer = require('puppeteer');
const assert = require('chai').assert;

(async () => {
    const browser = await puppeteer.launch({
        headless: false,
    });

    const page = await browser.newPage();
    await page.goto('http://localhost:8090/view?id=0');

    await page.$eval('#i-secret', el => el.value = 'password');
    await page.click('#b-submit');

    const messageId = await page.$eval('#id', el => el.textContent);
    assert.equal(messageId, '0');

    const message = await page.$eval('#message', el => el.textContent);
    assert.equal(message, 'decrypted: test message');

    await browser.close();
})();
