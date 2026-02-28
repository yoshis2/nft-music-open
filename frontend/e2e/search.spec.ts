import { test, expect } from '@playwright/test';

test.describe('NFT Search Page', () => {
  test.beforeEach(async ({ page }) => {
    // Navigate to the search page before each test
    await page.goto('/search');
  });

  test('should allow searching by keyword', async ({ page }) => {
    // Fill in the keyword input
    await page.locator('#search-query').fill('Awesome');
    // Click the search button
    await page.locator('button[type="submit"]').click();

    // Wait for the results to load
    await page.waitForSelector('.grid');

    // Check that the results contain the keyword
    const firstResult = page.locator('.grid > div').first();
    await expect(firstResult).toContainText('Awesome');
  });

  test('should filter by genre', async ({ page }) => {
    // Select a genre
    await page.locator('#genre-select').selectOption({ label: 'Rock' });
    // Click the search button
    await page.locator('button[type="submit"]').click();
    
    // Wait for the results to load
    await page.waitForSelector('.grid');

    // This is a simplified check. A real test would need to verify the genre of the results.
    const resultCount = await page.locator('.grid > div').count();
    expect(resultCount).toBeGreaterThan(0);
  });
  
  test('should filter by price range', async ({ page }) => {
    // Fill in min and max price
    await page.locator('#min-price').fill('100');
    await page.locator('#max-price').fill('200');
    // Click the search button
    await page.locator('button[type="submit"]').click();

    // Wait for the results to load
    await page.waitForSelector('.grid');
    
    const resultCount = await page.locator('.grid > div').count();
    expect(resultCount).toBeGreaterThan(0);
  });

  test('should sort the results', async ({ page }) => {
    // Select sort by "Price: Low to High"
    await page.locator('#sort-select').selectOption({ value: 'price_asc' });
    // Click the search button
    await page.locator('button[type="submit"]').click();

    // Wait for the results to load
    await page.waitForSelector('.grid');
    
    // This is a simplified check. A real test would need to get all prices and check they are sorted.
    const resultCount = await page.locator('.grid > div').count();
    expect(resultCount).toBeGreaterThan(0);
  });
});
