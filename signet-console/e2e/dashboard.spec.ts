import { test, expect } from '@playwright/test';

test('dashboard renders correctly', async ({ page }) => {
  await page.goto('http://localhost:3000');
  
  // Check for the main heading
  await expect(page.getByText('Governance Dashboard')).toBeVisible();
  
  // Check for the Merkle Tree title
  await expect(page.getByText('Latest Seal Integrity Tree')).toBeVisible();
  
  // Check for the SIGNET brand
  await expect(page.getByText('SIGNET')).toBeVisible();
  
  // Check for the Stats
  await expect(page.getByText('Sealed Artifacts')).toBeVisible();
});

test('navigation works', async ({ page }) => {
  await page.goto('http://localhost:3000');
  
  // Click on the Artifacts nav item
  await page.click('text=Artifacts');
  
  // URL should update (even if page is blank/404 for now, the nav should trigger)
  await expect(page).toHaveURL(/.*artifacts/);
});
