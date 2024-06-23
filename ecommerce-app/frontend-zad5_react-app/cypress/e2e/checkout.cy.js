describe('CheckoutPage Component Tests', () => {
  beforeEach(() => {
    cy.intercept('GET', '/koszyk/1', { fixture: 'cartItems.json' }).as('getCartItems');
    document.cookie = 'koszykID=1';
    cy.visit('/checkout');
  });

  it('display cart and payment forms', () => {
    cy.contains('Your Cart').should('be.visible');
    cy.contains('Payment Details').should('be.visible');
  });

  it('fetch cart ID from cookie', () => {
    cy.contains('Your Cart').should('be.visible');
  });

  it('display cart contents', () => {
    cy.wait('@getCartItems');
    cy.contains('Product 1').should('be.visible');
    cy.contains('Product 2').should('be.visible');
  });

  it('display total price on checkout page', () => {
    cy.wait('@getCartItems');
    cy.contains('Total: $30.00').should('be.visible');
  });
});
