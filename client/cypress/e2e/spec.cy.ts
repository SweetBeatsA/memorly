describe('My First Test', () => {
  it('Visits the Home Page and goes to log in', () => {
    cy.visit('http://memorly.kro.kr/')

    cy.contains('Log in').click()
  })
})