describe('Koszyk Component Tests', () => {
  beforeEach(() => {
    cy.intercept('GET', '/koszyk/1', { fixture: 'cartItems.json' }).as('getCartItems');
    cy.visit('/checkout');
  });

  it('fetch cart items successfully', () => {
    cy.wait('@getCartItems');
    cy.get('li').should('have.length', 2);
    cy.contains('Product 1').should('be.visible');
    cy.contains('Product 2').should('be.visible');
  });

  it('handle error when fetching cart items', () => {
    cy.intercept('GET', '/koszyk/1', { statusCode: 500, body: 'Error fetching cart items' }).as('getCartItemsError');
    cy.visit('/checkout');
    cy.wait('@getCartItemsError');
    cy.contains('Error fetching cart items').should('be.visible');
  });

  it('calculate total price', () => {
    cy.wait('@getCartItems');
    cy.contains('Total: $30.00').should('be.visible');
  });

  it('display empty cart message', () => {
    cy.intercept('GET', '/koszyk/1', { fixture: 'emptyCart.json' }).as('getEmptyCart');
    cy.visit('/checkout');
    cy.wait('@getEmptyCart');
    cy.contains('Your cart is empty.').should('be.visible');
  });
});
