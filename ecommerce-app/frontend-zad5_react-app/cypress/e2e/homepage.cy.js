describe('HomePage Component Tests', () => {
  it('display welcome message', () => {
    cy.visit('/');
    cy.contains('Welcome to Our Store').should('be.visible');
  });

  it('navigate to checkout', () => {
    cy.visit('/');
    cy.contains('Go to Checkout').click();
    cy.url().should('include', '/checkout');
  });
});
