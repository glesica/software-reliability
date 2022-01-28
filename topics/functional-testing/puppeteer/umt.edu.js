const puppeteer = require('puppeteer');

// Verify the main menu options on the front page of the UM website.

async function clickGlobalMenuItem(page, itemIndex) {
    const link = await page.evaluateHandle((index) => {
        return document
            .querySelector('umt-global-nav')
            .shadowRoot
            .querySelector(`nav li:nth-child(${index}) a`);
    }, itemIndex);
    await link.click();
    await page.waitForNavigation();
}

async function getPageHeading(page) {
    await page.waitForSelector('h1');
    return await page.$eval('h1', (e) => e.innerHTML);
}

function assertEqual(actual, expected) {
    if (actual !== expected) {
        console.error(`expected: ${expected}`);
        console.error(`actual:   ${actual}`);
    }
}

async function verifyGlobalNavItem(browser, itemIndex, pageHeading) {
    const page = await browser.newPage();
    await page.goto('https://www.umt.edu');
    await clickGlobalMenuItem(page, itemIndex);
    assertEqual(await getPageHeading(page), pageHeading);
    await page.close();
}

(async () => {
    const browser = await puppeteer.launch({
        headless: true,
        defaultViewport: {
            width: 1366,
            height: 900,
        },
    });

    await verifyGlobalNavItem(browser, 1, 'About UM');
    await verifyGlobalNavItem(browser, 2, 'Academics at UM');
    await verifyGlobalNavItem(browser, 3, 'Admissions and Financial Aid');
    await verifyGlobalNavItem(browser, 4, 'Student Life at UM');
    await verifyGlobalNavItem(browser, 5, 'Research at UM');
    await verifyGlobalNavItem(browser, 6, 'Athletics');
    await verifyGlobalNavItem(browser, 8, 'Alumni');

    await browser.close();
})();
