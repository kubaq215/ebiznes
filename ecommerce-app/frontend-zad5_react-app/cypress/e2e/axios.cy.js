describe('Axios Interceptors Tests', () => {
  beforeEach(() => {
    cy.window().then(win => {
      win.sessionStorage.setItem('authToken', 'testToken');
    });
    cy.visit('/');
  });

  // it('request interceptor - add token', () => {
  //   cy.intercept('GET', '/some-endpoint', req => {
  //     expect(req.headers).to.have.property('Authorization', 'Bearer testToken');
  //     req.reply({ statusCode: 200 });
  //   }).as('getSomeEndpoint');
  //   cy.request('/some-endpoint');
  //   cy.wait('@getSomeEndpoint');
  // });

  it('response interceptor - handle errors', () => {
    cy.intercept('GET', '/some-endpoint', { statusCode: 500, body: 'Test Error' }).as('getSomeEndpointError');
    cy.on('uncaught:exception', (err) => {
      expect(err.message).to.include('Something went wrong: Test Error');
      return false;
    });
    cy.request({ url: '/some-endpoint', failOnStatusCode: false });
    cy.wait('@getSomeEndpointError');
  });
});
