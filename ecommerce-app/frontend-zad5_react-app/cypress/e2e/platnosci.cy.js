describe('Platnosci Component Tests', () => {
  beforeEach(() => {
    cy.visit('/checkout');
  });

  it('submit payment successfully', () => {
    cy.intercept('POST', '/pay', {}).as('submitPayment');
    cy.intercept('DELETE', '/koszyk/1', {}).as('clearCart');
    cy.get('[name="cardNumber"]').type('4111111111111111');
    cy.get('[name="expiryDate"]').type('12/24');
    cy.get('[name="cvv"]').type('123');
    cy.get('[name="amount"]').type('10.00');
    cy.contains('Submit Payment').click();
    cy.wait('@submitPayment');
    cy.wait('@clearCart');
    cy.contains('Payment Successful!').should('be.visible');
  });

  it('handle error when submitting payment', () => {
    cy.intercept('POST', '/pay', { statusCode: 500, body: 'Error during payment' }).as('submitPaymentError');
    cy.get('[name="cardNumber"]').type('4111111111111111');
    cy.get('[name="expiryDate"]').type('12/24');
    cy.get('[name="cvv"]').type('123');
    cy.get('[name="amount"]').type('10.00');
    cy.contains('Submit Payment').click();
    cy.wait('@submitPaymentError');
    cy.contains('Payment or Cart clearing failed. Please try again later.').should('be.visible');
  });

  it('validate card number field', () => {
    cy.get('[name="expiryDate"]').type('12/24');
    cy.get('[name="cvv"]').type('123');
    cy.get('[name="amount"]').type('10.00');
    cy.contains('Submit Payment').click();
    cy.get('[data-testid="error-message"]').should('have.text', '400');
  });

  it('validate expiry date field', () => {
    cy.get('[name="cardNumber"]').type('4111111111111111');
    cy.get('[name="cvv"]').type('123');
    cy.get('[name="amount"]').type('10.00');
    cy.contains('Submit Payment').click();
    cy.get('[data-testid="error-message"]').should('have.text', '400');
  });

  it('validate CVV field', () => {
    cy.get('[name="cardNumber"]').type('4111111111111111');
    cy.get('[name="expiryDate"]').type('12/24');
    cy.get('[name="amount"]').type('10.00');
    cy.contains('Submit Payment').click();
    cy.get('[data-testid="error-message"]').should('have.text', '400');
  });

  it('validate amount field', () => {
    cy.get('[name="cardNumber"]').type('4111111111111111');
    cy.get('[name="expiryDate"]').type('12/24');
    cy.get('[name="cvv"]').type('123');
    cy.contains('Submit Payment').click();
    cy.get('[data-testid="error-message"]').should('have.text', '400');
  });
});
