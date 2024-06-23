it('check whether the response code is 400', () => {
  cy.intercept('POST', '/pay', { statusCode: 400 }).as('submitPayment');
  cy.get('[name="cardNumber"]').type('4111111111111111');
  cy.get('[name="expiryDate"]').type('12/24');
  cy.get('[name="cvv"]').type('123');
  cy.get('[name="amount"]').type('10.00');
  cy.contains('Submit Payment').click();
  cy.wait('@submitPayment');
  cy.get('[data-testid="error-message"]').should('have.text', '400');
});