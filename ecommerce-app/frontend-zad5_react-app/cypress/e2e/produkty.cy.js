describe('Produkty Component Tests', () => {
  beforeEach(() => {
    cy.intercept('GET', '/produkty', { fixture: 'products.json' }).as('getProducts');
    cy.visit('/');
  });

  it('fetch products successfully', () => {
    cy.wait('@getProducts');
    cy.get('.product-item').should('have.length', 2);
    cy.contains('Product 1').should('be.visible');
    cy.contains('Product 2').should('be.visible');
  });

  it('handle error when fetching products', () => {
    cy.intercept('GET', '/produkty', { statusCode: 500, body: 'Error fetching products' }).as('getProductsError');
    cy.visit('/');
    cy.wait('@getProductsError');
    cy.contains('Error fetching products').should('be.visible');
  });

  it('add product to new cart', () => {
    cy.intercept('POST', '/newkoszyk', { fixture: 'newCart.json' }).as('newCart');
    cy.intercept('POST', '/koszyk/*/*', {}).as('addToCart');
    cy.contains('Add to Cart').click();
    cy.wait('@newCart');
    cy.wait('@addToCart');
    cy.contains('Product added to new cart').should('be.visible');
  });

  it('add product to existing cart', () => {
    document.cookie = 'koszykID=1';
    cy.intercept('POST', '/koszyk/1/*', {}).as('addToExistingCart');
    cy.contains('Add to Cart').click();
    cy.wait('@addToExistingCart');
    cy.contains('Product added to existing cart').should('be.visible');
  });

  it('handle error when adding product to cart', () => {
    document.cookie = 'koszykID=1';
    cy.intercept('POST', '/koszyk/1/*', { statusCode: 500, body: 'Error adding product to cart' }).as('addToCartError');
    cy.contains('Add to Cart').click();
    cy.wait('@addToCartError');
    cy.contains('Error adding product to cart').should('be.visible');
  });
});
